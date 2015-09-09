package apidoc

import (
	"fmt"
	"reflect"
	"strings"
)

type (
	Argument struct {
		Field        string
		Description  string
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

//--------------------------------------------
// PUBLIC METHODS
//--------------------------------------------

func NewArguments(v interface{}) []*Argument {
	var args []*Argument
	if v != nil {
		rv := reflect.ValueOf(v)
		rt := reflect.TypeOf(v)
		args = objectAnalysisDetail(rv, rt, nil)
	}
	return args
}

//--------------------------------------------
// PACKAGE PRIVATE METHODS
//--------------------------------------------

func objectAnalysisDetail(rv reflect.Value, rt reflect.Type, option *AnalysisOption) []*Argument {

	var params []*Argument

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

		params = objectAnalysisDetail(rv, rt, option)

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
			p := generateArguments(typeName, option)
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
				p := objectAnalysisDetail(valField, typ, option)
				params = append(params, p...)
			}
		}

	} else {
		params = generateArguments(rt.Name(), option)
	}
	return params
}

func generateArguments(typeName string, option *AnalysisOption) []*Argument {

	var params []*Argument

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

		p := &Argument{
			Field:        field,
			Description:  "It is " + field + ".",
			TypeName:     convertJsonTypeName(typeName),
			DefaultValue: "",
		}
		if tag := getDocTag(option.Tag); tag != "" {
			p.Description = tag
		}
		params = []*Argument{p}
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
