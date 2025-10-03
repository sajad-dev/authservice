package grpcserver

import (
	"fmt"
	"log"
	"net"

	"github.com/sajad-dev/authservice/authorization/internal/bootstrap"
	"github.com/sajad-dev/authservice/authorization/internal/config"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/http"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs"
	"google.golang.org/grpc"
)

type Grpc struct {
	Config config.AppConfig
}

func NewGrpc(config config.AppConfig) *Grpc {
	return &Grpc{
		Config: config,
	}
}

func (g *Grpc) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", g.Config.GRPC_PORT))
	if err != nil {
		return errs.Err(err)
	}

	gc := grpc.NewServer()

	err = bootstrap.NewBootstrap(g.Config).Boot(gc)

	log.Printf("Run server at port %s \n", g.Config.GRPC_PORT)

	gc.Serve(lis)

	return nil
}

var _ http.Http = &Grpc{}
