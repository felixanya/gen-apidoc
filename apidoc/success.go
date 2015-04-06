package apidoc

import (
	"bytes"
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
	ps := objectAnalysis(v)
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
	ad.SetSuccessExample(v)
}

func (ad *ApiDefine) AddSuccess(field, description string) {
	success := &ApiSuccess{
		Field:       field,
		Description: description,
	}
	ad.Success = append(ad.Success, success)
}

func (ad *ApiDefine) AddSuccessByOptional(group, typeName, field, description string) {
	success := &ApiSuccess{
		Group:       group,
		TypeName : typeName,
		Field:       field,
		Description: description,
	}
	ad.Success = append(ad.Success, success)
}

func (ad *ApiDefine) AddApiSuccessWithConfig(success *ApiSuccess) {
	ad.Success = append(ad.Success, success)
}

func (ad *ApiDefine) SetSuccessExample(v interface{}) {
	ad.SuccessExample = newExample(v, exampleTypesuccess)
	ad.SuccessExample.Title = "Response (success)"
	ad.SuccessExample.ProtocolAndStatus = "HTTP/1.1 200 OK"
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
	b.Write([]byte("@apiSuccess"))
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
