package dto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type FieldReq struct {
	To          string
	Type        string
	Json        string
	Validations string
	From        string
}

type MapperDataReq struct {
	FuncName     string
	StructName   string
	ProtoPackage string
	ImportPath   []string
	ProtoStruct  string
	Fields       []FieldReq
}

type JsonModuleReq struct {
	Name string
	Data MapperDataReq
}

type JsonFormatReq struct {
	ProtoPackage string
	ImportPath   []string
	Data         []JsonModuleReq
}



func requestGenCode(path string, tmpl *template.Template) {
	domainName := filepath.Base(path)
	confPath := filepath.Join(path, REQUEST_PATH)

	conf, err := os.ReadFile(confPath)
	if err != nil {
		panic(err)
	}

	var jsonFormat JsonFormatReq
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

		outPath := fmt.Sprintf(REQUEST_SAVE_GEN_CODE, domainName, data.Name, domainName)

		if err := os.MkdirAll(filepath.Dir(outPath), 0o755); err != nil {
			panic(err)
		}

		if err := os.WriteFile(outPath, buf.Bytes(), 0o644); err != nil {
			panic(err)
		}
	}
}


