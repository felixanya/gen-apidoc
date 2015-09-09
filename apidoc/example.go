package apidoc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	exampleTypeHeader  ExampleType = "@apiHeaderExample"
	exampleTypeParam   ExampleType = "@apiParamExample"
	exampleTypeSuccess ExampleType = "@apiSuccessExample"
	exampleTypeError   ExampleType = "@apiErrorExample"
)

type (
	Example struct {
		Data              []byte
		Title             string
		TypeName          string
		typ               ExampleType //Example Type
		ProtocolAndStatus string
		Status            int
	}
	ExampleType string
)

// NewParamExample is new param example
func NewParamExample(title string, v interface{}) *Example {
	return newExample(title, v, exampleTypeParam, 0)
}

// NewErrorExample is new error example
func NewErrorExample(title string, v interface{}, status int) *Example {
	return newExample(title, v, exampleTypeError, status)
}

// NewSuccessExample is new success example with status code
func NewSuccessExample(title string, v interface{}, status int) *Example {
	return newExample(title, v, exampleTypeSuccess, status)
}

// NewSuccessExample200 is new success example and status OK
func NewSuccessExample200(title string, v interface{}) *Example {
	return newExample(title, v, exampleTypeSuccess, http.StatusOK)
}

func newExample(title string, v interface{}, et ExampleType, code int) *Example {
	b, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	if title == "" {
		title = "JSON Example"
	}
	return &Example{
		Data:              b,
		Title:             title,
		TypeName:          "json",
		typ:               et,
		Status:            code,
		ProtocolAndStatus: fmt.Sprintf("HTTP/1.1 %d %s", code, http.StatusText(code)),
	}
}

func (e *Example) WriteIndentString(b *bytes.Buffer) {
	// @apiXXXXXExample [{type}] [title]
	b.Write(bPrefix)
	b.Write([]byte(e.typ))
	if e.TypeName != "" {
		b.Write([]byte(" {" + e.TypeName + "}"))
	}
	if e.Title != "" {
		b.Write([]byte(" " + e.Title))
	}
	b.Write(bbreak)

	if e.ProtocolAndStatus != "" {
		b.Write(bPrefix)
		b.Write([]byte(e.ProtocolAndStatus))
		b.Write(bbreak)
	}

	b.Write(bPrefix)
	json.Indent(b, e.Data, blockCommentPrefix, jsonIndentString)
	b.Write(bbreak)
}
