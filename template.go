package main

import (
	"bytes"
	"encoding/json"
	"os"
	"text/template"
)

type Field struct {
	To   string
	From string
}

type MapperData struct {
	Package         string
	FuncName        string
	StructName      string
	ProtoPackage    string
	ProtoImportPath string
	ProtoStruct     string
	Fields          []Field
}

func main() {
	// load config
	conf, err := os.ReadFile("t.json")
	if err != nil {
		panic(err)
	}
	var data MapperData
	json.Unmarshal(conf, &data)

	// parse template
	tmpl, err := template.ParseFiles("t.tmpl")
	if err != nil {
		panic(err)
	}

	// render
	var buf bytes.Buffer
	tmpl.Execute(&buf, data)

	// save output
	outPath := "./create_req_mapper.go"
	err = os.WriteFile(outPath, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

