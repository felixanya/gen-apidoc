package apidoc

import (
	"bytes"
)

type (
	ApiError struct {
		Field       string
		Description string
		Group       string
		TypeName    string
	}
)

func (ad *ApiDefine) SetError(code int, v interface{}) {
	ps := objectAnalysis(v)
	var ss []*ApiError
	for _, p := range ps {
		ss = append(ss, &ApiError{
			Field:       p.Field,
			Description: p.Description,
			Group:       p.Group,
			TypeName:    p.TypeName,
		})
	}
	ad.Errors = ss

	if v != nil {
		ad.SetErrorExample(v)
	}
}

func (ad *ApiDefine) SetErrorExample(v interface{}) {
	ad.ErrorExample = newExample(v, exampleTypesuccess)
	ad.ErrorExample.Title = "Response (error)"
	ad.ErrorExample.ProtocolAndStatus = "HTTP/1.1 200 OK"
}

func (ad *ApiDefine) AddError(field, description string) {
	ad.Errors = append(ad.Errors, &ApiError{
		Field:       field,
		Description: description,
	})
}

func (ad *ApiDefine) AddErrorByOptional(success *ApiError) {
	ad.Errors = append(ad.Errors, success)
}

// -----------------------
// @ApiError
// -----------------------

func (as *ApiError) String() string {
	return string(as.Byte())
}

func (as *ApiError) Byte() []byte {
	// @ApiParam [(group)] [{type}] [field=defaultValue] [description]
	var b bytes.Buffer
	b.Write([]byte("@ApiError"))
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
