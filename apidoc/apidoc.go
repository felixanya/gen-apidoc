package apidoc

import (
	"bufio"
	"bytes"
	"os"
)

const (
	outputFileName      = ".apidoc_gen.go"
)

type (
	ApiDocument struct {
		OutputFileName           string
		PackageName          string
		ApiDefines         []ApiDefine
	}
)

func NewDocument(packageName string) ApiDocument {
	os.Remove(outputFileName)
	return ApiDocument{
		PackageName: packageName,
		OutputFileName: outputFileName,
	}
}

func (doc *ApiDocument) New(name string) ApiDefine {
	return ApiDefine{
		Name: name,
	}
}
func (doc *ApiDocument) Add(define ApiDefine) {
	doc.ApiDefines = append(doc.ApiDefines, define)
}

func (doc *ApiDocument) Write() error {

	var bt bytes.Buffer
	
	// Header
	writeApidocHeader(doc.PackageName, &bt)
	
	// Define items
	for _, define := range doc.ApiDefines {
		define.WriteBytes(&bt)
	}
	
	err := writeFile(doc.OutputFileName, bt.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// -------------------------------------
// PRIVATE METHOD
// -------------------------------------

func writeFile(fileName string, contents []byte) error {

	writeFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(writeFile)
	if _, err := writer.Write(contents); err != nil {
		return err
	}
	writer.Flush()
	return nil
}

func writeApidocHeader(packageName string, b *bytes.Buffer) {
	b.Write([]byte("package " + packageName))
	b.Write(bbreak)
	b.Write(bbreak)
}
