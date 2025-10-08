package main

import (
	"github.com/sajad-dev/authservice/authentication/internal/config"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/http/grpcserver"
)

func main() {
	err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	err = grpcserver.NewGrpc(config.Cfg).Run()
	if err != nil {
		panic(err)
	}
}
