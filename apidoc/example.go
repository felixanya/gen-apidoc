package apidoc

import (
	"bytes"
	"encoding/json"
)

const (
	exampleTypeHeader  ExampleType = "@apiHeaderExample"
	exampleTypeParam   ExampleType = "@apiParamExample"
	exampleTypesuccess ExampleType = "@apiSuccessExample"
)

type (
	Example struct {
		Data              []byte
		Title             string
		TypeName          string
		Typ               ExampleType //Example Type
		ProtocolAndStatus string
		Status int
	}
	ExampleType string
)

func newExample(v interface{}, et ExampleType, status int) *Example {
	b, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return &Example{
		Data:     b,
		Title:    "JSON Example",
		TypeName: "json",
		Typ:      et,
		Status:   status,
	}
}

func (e *Example) WriteIndentString(b *bytes.Buffer) {
	// @apiXXXXXExample [{type}] [title]
	b.Write(bPrefix)
	b.Write([]byte(e.Typ))
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
