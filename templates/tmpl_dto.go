package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type Field struct {
	To          string
	Type        string
	Json        string
	Validations string
	From        string
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

type JsonModule struct {
	Name string
	Data MapperData
}

type JsonFormat struct {
	Package         string
	ProtoPackage    string
	ProtoImportPath string
	Data            []JsonModule
}

const (
	DOMAIN_PATH            = "../internal/domain"
	REQUEST_PATH           = "/mapper/request.json"
	RESPONSE_PATH          = "/mapper/response.json"
	REQUEST_SAVE_GEN_CODE  = "../internal/domain/%s/mapper/gen/%s_%s_req.go"
	RESPONSE_SAVE_GEN_CODE = "../internal/domain/%s/mapper/gen/%s_%s_res.go"
	TEMPL_FILE             = "../t.tmpl"
)

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

func requestGenCode(path string, tmpl *template.Template) {
	domainName := filepath.Base(path)
	confPath := filepath.Join(path, REQUEST_PATH)

	conf, err := os.ReadFile(confPath)
	if err != nil {
		panic(err)
	}

	var jsonFormat JsonFormat
	if err := json.Unmarshal(conf, &jsonFormat); err != nil {
		panic(err)
	}

	for _, data := range jsonFormat.Data {
		data.Data.Package = jsonFormat.Package
		data.Data.ProtoPackage = jsonFormat.ProtoPackage
		data.Data.ProtoImportPath = jsonFormat.ProtoImportPath
		data.Data.ProtoImportPath = jsonFormat.ProtoImportPath

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data.Data); err != nil {
			panic(err)
		}

		outPath := fmt.Sprintf(REQUEST_SAVE_GEN_CODE, domainName, data.Name, domainName)
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

	var jsonFormat JsonFormat
	if err := json.Unmarshal(conf, &jsonFormat); err != nil {
		panic(err)
	}

	for _, data := range jsonFormat.Data {
		data.Data.Package = jsonFormat.Package
		data.Data.ProtoPackage = jsonFormat.ProtoPackage
		data.Data.ProtoImportPath = jsonFormat.ProtoImportPath

		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data.Data); err != nil {
			panic(err)
		}

		outPath := fmt.Sprintf(RESPONSE_SAVE_GEN_CODE, domainName, data.Name, domainName)
		if err := os.WriteFile(outPath, buf.Bytes(), 0o644); err != nil {
			panic(err)
		}
	}
}

func HandleGenCode() {
	tmpl, err := template.ParseFiles(TEMPL_FILE)
	if err != nil {
		panic(err)
	}

	for _, dir := range readDomainDir() {
		requestGenCode(dir, tmpl)
		responseGenCode(dir, tmpl)
	}
}
