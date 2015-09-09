package apidoc

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

type (
	ApiParam struct {
		Field        string
		Description  string
		Group        string
		TypeName     string
		Kind         reflect.Kind
		DefaultValue string
		Optional     bool
	}
	AnalysisOption struct {
		Index      int
		Field      reflect.StructField
		ParentTag  []reflect.StructTag
		Tag        reflect.StructTag
		ColumnName string
		Levels     []string
	}
)

func (ad *ApiDefine) SetParams(group string, v interface{}) {
	ad.Group = group
	ad.Params = objectAnalysis(group, v)

	if v != nil {
		ad.SetParamExample(v)
	}
}

func (ad *ApiDefine) SetParamExample(v interface{}) {
	ad.ParamExample = newExample("Parameter Example", v, exampleTypeParam, 0)
}

func (ad *ApiDefine) AddParam(field, description string) {
	param := &ApiParam{
		Field:       field,
		Description: description,
	}
	ad.Params = append(ad.Params, param)
}

func (ad *ApiDefine) AddParamByOptional(group, field, description, defaultValue string) {
	param := &ApiParam{
		Group:        group,
		Field:        field,
		Description:  description,
		DefaultValue: defaultValue,
		Optional:     true,
	}
	ad.Params = append(ad.Params, param)
}

func (ad *ApiDefine) AddApiParamWithConfig(param *ApiParam) {
	ad.Params = append(ad.Params, param)
}

// -----------------------
// @ApiParam
// -----------------------

func (ap *ApiParam) String() string {
	return string(ap.Byte())
}

func (ap *ApiParam) Byte() []byte {
	// @ApiParam [(group)] [{type}] [field=defaultValue] [description]
	var b bytes.Buffer
	b.Write([]byte("@apiParam"))
	if ap.Group != "" {
		b.Write([]byte(" (" + ap.Group + ")"))
	}
	if ap.TypeName != "" {
		b.Write([]byte(" {" + ap.TypeName + "}"))
	}
	if ap.Field != "" {
		str := ap.Field
		if ap.DefaultValue != "" {
			str += "=" + ap.DefaultValue
		}
		if ap.Optional {
			str = "[" + str + "]"
		}
		b.Write([]byte(" " + str))
	}
	if ap.Description != "" {
		b.Write([]byte(" " + ap.Description))
	}
	return b.Bytes()
}

func objectAnalysis(group string, v interface{}) []*ApiParam {
	var resp []*ApiParam
	if v != nil {
		rv := reflect.ValueOf(v)
		rt := reflect.TypeOf(v)
		resp = objectAnalysisDetail(group, rv, rt, nil)
	}
	return resp
}

func objectAnalysisDetail(group string, rv reflect.Value, rt reflect.Type, option *AnalysisOption) []*ApiParam {

	var params []*ApiParam

	if rt.Kind() == reflect.Ptr ||
		rt.Kind() == reflect.Slice ||
		rt.Kind() == reflect.Map {

		elem := rt.Elem()
		obj := reflect.New(elem).Elem().Interface()
		rv = reflect.ValueOf(obj)
		rt = rv.Type()

		if option != nil {
			option.Levels = append(option.Levels, "p")
		}

		params = objectAnalysisDetail(group, rv, rt, option)

	} else if rt.Kind() == reflect.Struct {

		//var parentTag string
		var exit bool
		var levels []string
		var parentTag []reflect.StructTag

		if option != nil {
			anonymous := option.Field.Anonymous
			_, jsonOK := getJSONTag(option.Tag)
			parentTag = append(option.ParentTag, option.Tag)
			levels = option.Levels
			exit = !anonymous && !jsonOK

			// struct用の記述を出力
			typeName := fmt.Sprintf("Object[%s]", rt.Name())
			p := generateApiParams(group, typeName, option)
			params = append(params, p...)
		}

		if option == nil || !exit {

			for i := 0; i < rt.NumField(); i++ {

				typField := rt.Field(i)
				valField := rv.Field(i)
				typ := valField.Type()

				option = &AnalysisOption{
					Index:      i,
					Field:      typField,
					Tag:        typField.Tag,
					ColumnName: typField.Name,
					ParentTag:  parentTag,
					Levels:     append(levels, "s"),
				}
				p := objectAnalysisDetail(group, valField, typ, option)
				params = append(params, p...)
			}
		}

	} else {
		params = generateApiParams(group, rt.Name(), option)
	}
	return params
}

func generateApiParams(group, typeName string, option *AnalysisOption) []*ApiParam {

	var params []*ApiParam

	if option == nil {
		return params
	}

	tags := append(option.ParentTag, option.Tag)

	field := ""
	for _, tag := range tags {
		if s, ok := getJSONTag(tag); ok {
			if field != "" {
				field += "."
			}
			field += s
		}
	}

	//println(">>>",option.Tag.Get("json"), field, "levels=", option.Levels, option.Index)

	if _, ok := getJSONTag(option.Tag); ok {

		p := &ApiParam{
			Field:        field,
			Description:  "It is " + field + ".",
			Group:        group,
			TypeName:     convertJsonTypeName(typeName),
			DefaultValue: "",
		}
		if tag := getDocTag(option.Tag); tag != "" {
			p.Description = tag
		}
		params = []*ApiParam{p}
	}
	return params
}

func getJSONTag(tag reflect.StructTag) (string, bool) {
	return getTagText(tag.Get("json"))
}

func getDocTag(tag reflect.StructTag) string {
	s, _ := getTagText(tag.Get("doc"))
	return s
}

func convertJsonTypeName(typeName string) string {
	switch typeName {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64":
		return "integer"
	case "Object[Time]":
		return "string[RFC3339]"
	case "ObjectId":
		return "string"
	}
	return typeName
}

func getTagText(str string) (string, bool) {

	if str == "" || str == "-" {
		return "", false
	}

	items := strings.Split(str, ",")

	value := ""
	for i, v := range items {
		if i == 0 {
			if v == "" {
				break
			}
			value = v
			//		} else if v == "omitempty" {
			//			value = ""
			//			break
		}
	}
	return value, value != ""
}
