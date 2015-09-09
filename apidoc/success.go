package apidoc

import (
	"bytes"
)

type (
	ApiSuccess struct {
		Group       string
		TypeName    string
		Field       string
		Description string
	}
)

func NewSuccessParam(data interface{}) []*ApiSuccess {
	args := NewArguments(data)
	var ss []*ApiSuccess
	for _, p := range args {
		ss = append(ss, &ApiSuccess{
			Group:       "success",
			Field:       p.Field,
			Description: p.Description,
			TypeName:    p.TypeName,
		})
	}
	return ss
}

func (ad *ApiDefine) AddSuccess(field string, params ...string) {
	success := &ApiSuccess{Field: field}
	if len(params) > 0 {
		success.Description = params[0]
	}
	if len(params) > 1 {
		success.Group = params[1]
	}
	if len(params) > 2 {
		success.Field = params[2]
	}
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
