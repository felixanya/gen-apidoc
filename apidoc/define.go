package apidoc

import (
	"bytes"
	"net/http"
)

const (
	blockCommentStart  = "/**"
	blockCommentPrefix = " * "
	blockCommentEnd    = " */"
	lineBreak          = "\n"
	jsonIndentString   = "\t"
)

type (
	ApiDefine struct {
		Name           string
		Group          string
		Request        *http.Request
		Headers        []*ApiHeader
		HeaderExample  *Example
		Params         []*ApiParam
		ParamExample   *Example
		Success        []*ApiSuccess
		SuccessExample *Example
		Errors         []*ApiError
		ErrorExample   *Example
		//Errors      []*ApiError
	}
)

var (
	bStartLine = []byte(blockCommentStart + lineBreak)
	bPrefix    = []byte(blockCommentPrefix)
	bEndLine   = []byte(blockCommentEnd + lineBreak)
	bSpaceLine = []byte(blockCommentPrefix + lineBreak)
	bbreak     = []byte(lineBreak)
)

//func New2(req *http.Request, name string) ApiDefine {
//	items := strings.Split(req.URL.Path, "/")
//	var packageName string
//	for _, u := range items {
//		if u != "" {
//			packageName = u
//			break
//		}
//	}
//	return ApiDefine{
//		Request: req,
//		Group:   strings.Title(packageName),
//		Name:    name,
//	}
//}

func New(name string) ApiDefine {
	return ApiDefine{
		Name: name,
	}
}

func (ad *ApiDefine) WriteBytes(b *bytes.Buffer) {

	// Define
	b.Write(bStartLine)
	b.Write(bPrefix)
	b.Write([]byte("@ApiDefine " + ad.Name))
	b.Write(bbreak)

	b.Write(bSpaceLine)

	// Header
	if len(ad.Headers) > 0 {
		for _, header := range ad.Headers {
			b.Write(bPrefix)
			b.Write(header.Byte())
			b.Write(bbreak)
		}
		b.Write(bSpaceLine)
	}
	if ad.HeaderExample != nil {
		ad.HeaderExample.WriteIndentString(b)
	}
	b.Write(bSpaceLine)

	// Param
	if len(ad.Params) > 0 {
		for _, param := range ad.Params {
			b.Write(bPrefix)
			b.Write(param.Byte())
			b.Write(bbreak)
		}
		b.Write(bSpaceLine)
	}
	if ad.ParamExample != nil {
		ad.ParamExample.WriteIndentString(b)
	}
	b.Write(bSpaceLine)

	// Success
	if len(ad.Success) > 0 {
		for _, param := range ad.Success {
			b.Write(bPrefix)
			b.Write(param.Byte())
			b.Write(bbreak)
		}
		b.Write(bSpaceLine)
	}
	if ad.SuccessExample != nil {
		ad.SuccessExample.WriteIndentString(b)
	}

	b.Write(bEndLine)
	b.Write(bbreak)
	b.Write(bbreak)
}

func WriteRow(bb *bytes.Buffer, text string) {
	WriteRowByte(bb, []byte(text))
}

func WriteRowByte(bb *bytes.Buffer, b []byte) {
	bb.Write(bPrefix)
	bb.Write(b)
	bb.Write(bbreak)
}
