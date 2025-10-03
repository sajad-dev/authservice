package main

import (
	"github.com/sajad-dev/authservice/authorization/internal/config"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/http/grpcserver"
)

func main() {
	conf := config.NewConfig()
	err := grpcserver.NewGrpc(*conf).Run()
	if err != nil {
		panic(err)
	}
}
