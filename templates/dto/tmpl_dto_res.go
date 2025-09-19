package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type FieldRes struct {
	To           string
	Type         string
	Function     string
	HaveFunction bool
	Json         string
	From         string
}

type AdaptorFieldRes struct {
	FuncName     string
	Data         any
	ProtoPackage string
	ProtoStruct  string
	Fields       []FieldRes
}

type MapperDataRes struct {
	FuncName     string
	StructName   string
	ProtoPackage string
	ImportPath   []string
	ProtoStruct  string
	AdaptorField []AdaptorFieldRes
	Fields       []FieldRes
}

type JsonModuleRes struct {
	Name string
	Data MapperDataRes
}

type JsonFormatRes struct {
	ProtoPackage string
	ImportPath   []string
	Data         []JsonModuleRes
}

func responseGenCode(path string, tmpl *template.Template) {
	domainName := filepath.Base(path)
	confPath := filepath.Join(path, RESPONSE_PATH)

	conf, err := os.ReadFile(confPath)
	if err != nil {
		panic(err)
	}

	var jsonFormat JsonFormatRes
	if err := json.Unmarshal(conf, &jsonFormat); err != nil {
		panic(err)
	}

	for _, data := range jsonFormat.Data {
		data.Data.ProtoPackage = jsonFormat.ProtoPackage
		for _, val := range jsonFormat.ImportPath {
			data.Data.ImportPath = append(data.Data.ImportPath, val)
		}

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data.Data); err != nil {
			panic(err)
		}

		outPath := fmt.Sprintf(RESPONSE_SAVE_GEN_CODE, domainName, data.Name, domainName)

		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			panic(err)
		}

		if err := os.WriteFile(outPath, buf.Bytes(), 0o644); err != nil {
			panic(err)
		}
	}
}
