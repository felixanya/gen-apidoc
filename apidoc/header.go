package apidoc

import (
	"bytes"
	"net/http"
	"strings"
)

type (
	ApiHeader struct {
		Group        string
		TypeName     string
		Field        string
		Description  string
		DefaultValue string
		Optional     bool
	}
)

func (ad *ApiDefine) SetHeaders(headers http.Header) {
	m := map[string]string{}
	for key, v := range headers {
		ad.AddHeader(key, "")
		m[key] = strings.Join(v, "")
	}

	ad.SetHeaderExample(m)
}

func (ad *ApiDefine) SetHeaderExample(v interface{}) {
	if v != nil {
		ad.HeaderExample = newExample("Headers Example", v, exampleTypeHeader, 0)
	}
}

func (ad *ApiDefine) AddHeader(key, description string) {
	header := &ApiHeader{
		Field:       key,
		Description: description,
	}
	ad.Headers = append(ad.Headers, header)
}

func (ad *ApiDefine) AddHeaderByOptional(header *ApiHeader) {
	ad.Headers = append(ad.Headers, header)
}

// -----------------------
// @ApiHeader
// -----------------------
func (ah *ApiHeader) String() string {
	b := ah.Byte()
	return string(b)
}

func (ah *ApiHeader) Byte() []byte {
	// @ApiHeader [(group)] [{type}] [field=defaultValue] [description]
	var b bytes.Buffer
	b.Write([]byte("@apiHeader"))
	if ah.Group != "" {
		b.Write([]byte(" (" + ah.Group + ")"))
	}
	if ah.TypeName != "" {
		b.Write([]byte(" {" + ah.TypeName + "}"))
	}
	if ah.Field != "" {
		var str string
		if ah.Optional {
			str = "[" + ah.Field + "]"
		} else {
			str = ah.Field
		}
		if ah.DefaultValue != "" {
			str += "=" + ah.DefaultValue
		}
		b.Write([]byte(" " + str))
	}
	if ah.Description != "" {
		b.Write([]byte(" " + ah.Description))
	}
	return b.Bytes()
}
