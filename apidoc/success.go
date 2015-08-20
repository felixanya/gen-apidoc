package apidoc

import (
	"bytes"
	"net/http"
)

type (
	ApiSuccess struct {
		Field       string
		Description string
		Group       string
		TypeName    string
	}
)

func NewSuccessParam(data interface{}) []*ApiSuccess {
	ps := objectAnalysis("success", data)
	var ss []*ApiSuccess
	for _, p := range ps {
		ss = append(ss, &ApiSuccess{
			Field:       p.Field,
			Description: p.Description,
			//Group:       group,
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

func (ad *ApiDefine) SetSuccessWithExample(title string, v interface{}) {
	ad.Success = []*ApiSuccess{}
	ad.SuccessExample = []*Example{}
	ad.AddSuccessWithExampleAndStatus(title, v, http.StatusOK)
}

func (ad *ApiDefine) SetSuccessWithExampleAndStatus(title string, v interface{}, status int) {
	ad.Success = []*ApiSuccess{}
	ad.SuccessExample = []*Example{}
	ad.AddSuccessWithExampleAndStatus(title, v, status)
}

func (ad *ApiDefine) AddSuccessWithExample(title string, v interface{}) {
	ad.AddSuccessWithExampleAndStatus(title, v, http.StatusOK)
}

func (ad *ApiDefine) AddSuccessWithExampleAndStatus(title string, v interface{}, status int) {
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

func (ad *ApiDefine) SetSuccessExample(title string, v interface{}) {
	ad.SuccessExample = []*Example{}
	ad.AddSuccessExampleWithStatus(title, v, http.StatusOK)
}

func (ad *ApiDefine) SetSuccessExampleWithStatus(title string, v interface{}, status int) {
	ad.SuccessExample = []*Example{}
	ad.AddSuccessExampleWithStatus(title, v, status)
}

func (ad *ApiDefine) AddSuccessExample(title string, v interface{}) {
	ad.AddSuccessExampleWithStatus(title, v, http.StatusOK)
}

func (ad *ApiDefine) AddSuccessExampleWithStatus(title string, v interface{}, status int) {
	if title == "" {
		title = "Response (success)"
	}
	example := newExample(title, v, exampleTypeSuccess, status)
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
