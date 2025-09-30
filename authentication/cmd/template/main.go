package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

const (
	DOMAIN_PATH            = "./internal/domain"
	REQUEST_PATH           = "/dto/gen/request.json"
	RESPONSE_PATH          = "/dto/gen/response.json"
	REQUEST_SAVE_GEN_CODE  = "./internal/domain/%s/dto/gen/request/%s_%s_req.go"
	RESPONSE_SAVE_GEN_CODE = "./internal/domain/%s/dto/gen/response/%s_%s_res.go"
	TEMPL_FILE_REQUEST     = "./templates/request.tmpl"
	TEMPL_FILE_RESPONSE    = "./templates/response.tmpl"
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

func readDomainDir() []string {
	dirEntries, err := os.ReadDir(DOMAIN_PATH)
	if err != nil {
		panic(err)
	}

	dirs := []string{}
	for _, entry := range dirEntries {
		if entry.IsDir() {
			dirs = append(dirs, filepath.Join(DOMAIN_PATH, entry.Name()))
		}
	}
	return dirs
}

func main() {
	tmplReq, err := template.ParseFiles(TEMPL_FILE_REQUEST)
	if err != nil {
		panic(err)
	}

	tmplRes, err := template.ParseFiles(TEMPL_FILE_RESPONSE)
	if err != nil {
		panic(err)
	}

	for _, dir := range readDomainDir() {
		requestGenCode(dir, tmplReq)
		responseGenCode(dir, tmplRes)
	}
}
