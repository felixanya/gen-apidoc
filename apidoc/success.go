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

func (ad *ApiDefine) SetSuccessWithExample(title string, v interface{}) {
	ad.Success = []*ApiSuccess{}
	ad.SuccessExample = []*Example{}
	ad.AddSuccessWithExample(title, v)
}

func (ad *ApiDefine) AddSuccessWithExample(title string, v interface{}) {
	ps := objectAnalysis("success", v)
	var ss []*ApiSuccess
	for _, p := range ps {
		ss = append(ss, &ApiSuccess{
			Field:       p.Field,
			Description: p.Description,
			Group:       title,
			TypeName:    p.TypeName,
		})
	}
	ad.Success = append(ad.Success, ss...)
	ad.AddSuccessExample(title, v)
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

func (ad *ApiDefine) SetSuccessExample(title string, v interface{}) {
	ad.SuccessExample = []*Example{}
	ad.AddSuccessExample(title, v)
}

func (ad *ApiDefine) AddSuccessExample(title string, v interface{}) {
	example := newExample(v, exampleTypesuccess)
	example.ProtocolAndStatus = "HTTP/1.1 200 OK"
	if title != "" {
		example.Title = title
	} else {
		example.Title = "Response (success)"
	}
	ad.SuccessExample = append(ad.SuccessExample, example)
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
