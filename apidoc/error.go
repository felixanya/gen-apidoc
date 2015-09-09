package apidoc

import (
	"bytes"
)

type (
	ApiError struct {
		Group       string
		TypeName    string
		Field       string
		Description string
	}
)

func NewErrorParam(data interface{}) []*ApiError {
	ps := NewArguments(data)
	var ss []*ApiError
	for _, p := range ps {
		ss = append(ss, &ApiError{
			Group:       "error",
			Field:       p.Field,
			Description: p.Description,
			TypeName:    p.TypeName,
		})
	}
	return ss
}

func (ad *ApiDefine) AddErrorParam(v interface{}) *ApiDefine {
	switch t := v.(type) {
	case *ApiError:
		ad.Errors = append(ad.Errors, v.(*ApiError))
	case []*ApiError:
		list := v.([]*ApiError)
		ad.Errors = append(ad.Errors, list...)
	default:
		println("failed to parameter type error", t)
		return nil
	}
	return ad
}

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
