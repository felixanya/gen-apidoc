package apidoc

import (
	"bytes"
	"fmt"
)

type (
ApiError struct {
	Field       string
	Description string
	Group       string
	TypeName    string
}
)

func (ad *ApiDefine) AddError(field string, params ...string) {
	success := &ApiError{Field: field}
	if len(params) > 0 {
		success.Description = params[0]
	}
	if len(params) > 1 {
		success.Group = params[1]
	}
	if len(params) > 2 {
		success.Field = params[2]
	}
	ad.Errors = append(ad.Errors, success)
}

func (ad *ApiDefine) SetErrorWithExample(title string, v interface{}, status int) {
	ad.Errors = []*ApiError{}
	ad.ErrorExample = []*Example{}
	ad.AddErrorWithExample(title, v, status)
}

func (ad *ApiDefine) AddErrorWithExample(title string, v interface{}, status int) {
	ps := objectAnalysis("error", v)
	var ss []*ApiError
	for _, p := range ps {
		ss = append(ss, &ApiError{
			Field:       p.Field,
			Description: p.Description,
			Group:       title,
			TypeName:    p.TypeName,
		})
	}
	ad.Errors = append(ad.Errors, ss...)
	ad.AddErrorExample(title, v, status)
}

func (ad *ApiDefine) SetErrorExample(title string, v interface{}, status int) {
	ad.ErrorExample = []*Example{}
	ad.AddErrorExample(title, v, status)
}

func (ad *ApiDefine) AddErrorExample(title string, v interface{}, status int) {
	example := newExample(v, exampleTypeError, status)
	example.ProtocolAndStatus = fmt.Sprintf("HTTP/1.1 %d ERROR", status)
	if title != "" {
		example.Title = title
	} else {
		example.Title = "Response (error)"
	}
	ad.ErrorExample = append(ad.ErrorExample, example)
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
	b.Write([]byte("@apiError"))
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
