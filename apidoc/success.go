package apidoc

import (
	"bytes"
	"reflect"
)

type (
	ApiSuccess struct {
		Field       string
		Description string
		Group       string
		TypeName    string
	}
)

func (ad *ApiDefine) SetSuccess(v interface{}) {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)
	ps := objectAnalysis("", rv, rt, nil)
	var ss []*ApiSuccess
	for _, p := range ps {
		ss = append(ss, &ApiSuccess{
			Field:       p.Field,
			Description: p.Description,
			Group:       p.Group,
			TypeName:    p.TypeName,
		})
	}
	ad.Success = ss

	if v != nil {
		ad.SetSuccessExample(v)
	}
}

func (ad *ApiDefine) SetSuccessExample(v interface{}) {
	ad.SuccessExample = newExample(v, exampleTypesuccess)
}

func (ad *ApiDefine) AddSuccess(field, description string) {
	ad.Success = append(ad.Success, &ApiSuccess{
		Field:       field,
		Description: description,
	})
}

func (ad *ApiDefine) AddSuccessByOptional(success *ApiSuccess) {
	ad.Success = append(ad.Success, success)
}

// -----------------------
// @ApiSuccess
// -----------------------

func (as *ApiSuccess) String() string {
	return string(as.Byte())
}

func (as *ApiSuccess) Byte() []byte {
	// @ApiParam [(group)] [{type}] [field=defaultValue] [description]
	var b bytes.Buffer
	b.Write([]byte("@ApiSuccess"))
	if as.Group != "" {
		b.Write([]byte(" (" + as.Group + ")"))
	}
	if as.TypeName != "" {
		b.Write([]byte(" {" + as.TypeName + "}"))
	}
	if as.Field != "" {
		b.Write([]byte(" " + as.Field))
	}
	if as.Description != "" {
		b.Write([]byte(" " + as.Description))
	}
	return b.Bytes()
}
