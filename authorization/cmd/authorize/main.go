package main

import (
	"github.com/sajad-dev/authservice/authorization/internal/config"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/http/grpcserver"
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
