package config

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type AppConfig struct {
	ROOT_PROJECT      string
	APP_NAME          string
	DESCRIPTION       string
	GRPC_PORT         string
	AUTHOR            string
	DEBUG             string
	JWT               string
	DATABASE_NAME     string
	DATABASE_USER     string
	DATABASE_PASSWORD string
	DATABASE_PORT     string
	DATABASE_HOST     string
	MODEL_CONF        string
}

func NewAppConfig() *AppConfig {
	_, b, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(b), "../..")

	return &AppConfig{
		ROOT_PROJECT:      root,
		APP_NAME:          "GOLANG_APP",
		DESCRIPTION:       "",
		AUTHOR:            "Sajad pourajam",
		DEBUG:             "true",
		JWT:               "",
		DATABASE_NAME:     "GOLANG_APP",
		DATABASE_USER:     "root",
		DATABASE_PASSWORD: "root",
		MODEL_CONF:        fmt.Sprintf("%s/casbin/model.conf", root),
		DATABASE_PORT:     "3306",
		DATABASE_HOST:     "127.0.0.1",
		GRPC_PORT:         "3000",
	}
}
