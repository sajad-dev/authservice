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
	Config config.Config
}

func NewGrpc(config config.Config) *Grpc {
	return &Grpc{
		Config: config,
	}
}

func (g *Grpc) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.Config.Server.Port))
	if err != nil {
		return errs.Err(err)
	}

	gc := grpc.NewServer()

	err = bootstrap.NewBootstrap(g.Config).Boot(gc)
	if err != nil {
		return err
	}

	log.Printf("Run server at port %d \n", g.Config.Server.Port)

	gc.Serve(lis)

	return nil
}

var _ http.Http = &Grpc{}
