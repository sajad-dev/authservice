package dto

import (
	"os"
	"path/filepath"
	"text/template"
)

const (
	DOMAIN_PATH            = "../../internal/domain"
	REQUEST_PATH           = "/dto/request.json"
	RESPONSE_PATH          = "/dto/response.json"
	REQUEST_SAVE_GEN_CODE  = "../../internal/domain/%s/dto/gen/request/%s_%s_req.go"
	RESPONSE_SAVE_GEN_CODE = "../../internal/domain/%s/dto/gen/response/%s_%s_res.go"
	TEMPL_FILE_REQUEST     = "../../req.tmpl"
	TEMPL_FILE_RESPONSE    = "../../res.tmpl"
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

func HandleGenCode() {
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
