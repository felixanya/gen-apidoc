package apidoc

import (
	"bytes"
)

type (
	ApiParam struct {
		Group             string
		TypeName          string
		TypeSizeMin       int
		TypeSizeMax       int
		TypeSizeRange     [2]int
		TypeAllowedValues []interface{}
		Field             string
		DefaultValue      string
		Description       string
		Optional          bool
	}
)

func (ad *ApiDefine) SetParams(group string, v interface{}) {
	args := NewArguments(v)
	var ap []*ApiParam
	for _, p := range args {
		ap = append(ap, &ApiParam{
			Field:        p.Field,
			Description:  p.Description,
			Group:        group,
			TypeName:     p.TypeName,
			Optional:     p.Optional,
			DefaultValue: p.DefaultValue,
		})
	}
	ad.Params = ap
	ad.Group = group
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
		var str string
		if ap.Optional {
			str = "[" + ap.Field + "]"
		} else {
			str = ap.Field
		}
		if ap.DefaultValue != "" {
			str += "=" + ap.DefaultValue
		}
		b.Write([]byte(" " + str))
	}
	if ap.Description != "" {
		b.Write([]byte(" " + ap.Description))
	}
	return b.Bytes()
}