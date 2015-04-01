package apidoc

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
)

type (
	ApiParam struct {
		Field        string
		Description  string
		Group        string
		TypeName     string
		DefaultValue string
	}
	AnalysisOption struct {
		Index      int
		Field      reflect.StructField
		Tag        reflect.StructTag
		ColumnName string
	}
)

var ix int64 = 0

func (ad *ApiDefine) SetParams(v interface{}) {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)
	ad.Params = objectAnalysis("", rv, rt, nil)

	if v != nil {
		ad.SetParamExample(v)
	}
}

func (ad *ApiDefine) SetParamExample(v interface{}) {
	ad.ParamExample = newExample(v, exampleTypeParam)
	ad.ParamExample.Title = "Parameter"
}

func (ad *ApiDefine) AddParam(field, description string) {
	header := &ApiParam{
		Field:       field,
		Description: description,
	}
	ad.Params = append(ad.Params, header)
}

func (ad *ApiDefine) AddApiParamByOptional(header *ApiParam) {
	ad.Params = append(ad.Params, header)
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
	b.Write([]byte("@ApiParam"))
	if ap.Group != "" {
		b.Write([]byte(" (" + ap.Group + ")"))
	}
	if ap.TypeName != "" {
		b.Write([]byte(" {" + ap.TypeName + "}"))
	}
	if ap.DefaultValue != "" {
		b.Write([]byte(" field=" + ap.DefaultValue))
	}
	if ap.Field != "" {
		b.Write([]byte(" " + ap.Field))
	}
	if ap.Description != "" {
		b.Write([]byte(" " + ap.Description))
	}
	return b.Bytes()
}

func objectAnalysis(prefix string, rv reflect.Value, rt reflect.Type, option *AnalysisOption) []*ApiParam {

	var params []*ApiParam
	if ix > 100 {
		return params
	}
	ix++

	if rt.Kind() == reflect.Ptr ||
		rt.Kind() == reflect.Slice ||
		rt.Kind() == reflect.Map {

		elem := rt.Elem()
		obj := reflect.New(elem).Elem().Interface()
		rv = reflect.ValueOf(obj)
		rt = rv.Type()
		params = objectAnalysis(prefix+"p", rv, rt, option)

	} else if rt.Kind() == reflect.Struct {
		//var parentTag string
		var anonymous bool
		var parentTagOK bool
		if option != nil {
			anonymous = option.Field.Anonymous
			_, parentTagOK = getJSONTag(option.Tag)
		}
		if option == nil || parentTagOK || anonymous {
			for i := 0; i < rt.NumField(); i++ {
				//fmt.Println(i,rv.Field(i),rv.Field(i).Type.Kind())
				val := rv.Field(i)
				typ := val.Type()
				typField := rt.Field(i)
				//if tag, ok := getJSONTag(typField.Tag); ok {
				//fmt.Println("tag=ok===", typField.Name, tag)
				opt := &AnalysisOption{
					Index:      i,
					Field:      typField,
					Tag:        typField.Tag,
					ColumnName: typField.Name,
				}
				//fmt.Println("PkgPath", rt.Field(i).Name, typ.PkgPath())
				//doc := "[" + t.Get("doc") + "]"
				p := objectAnalysis(prefix+rt.Name()+strconv.Itoa(i), val, typ, opt)
				params = append(params, p...)
				//} else {
				//	fmt.Println("tag=ng===", typField.Name, typField.Tag.Get("json"))
				//}
			}
		}

	} else if option != nil {

		if tag, ok := getJSONTag(option.Tag); ok {
			p := &ApiParam{
				Field:       tag,
				Description: "It is " + tag + ".",
				//Group:        prefix,
				TypeName:     rt.Name(),
				DefaultValue: "",
			}
			if tag := getDocTag(option.Tag); tag != "" {
				p.Description = tag
			}
			//}
			params = []*ApiParam{p}
		}
	}
	return params
}

func getJSONTag(tag reflect.StructTag) (string, bool) {
	str := tag.Get("json")
	if str != "" {
		items := strings.Split(str, ",")
		if len(items) > 0 && items[0] != "-" && items[0] != "" {
			return items[0], true
		}
	}
	return "", false
}

func getDocTag(tag reflect.StructTag) string {
	if str := tag.Get("doc"); str != "" {
		return str
	}
	return ""
}
