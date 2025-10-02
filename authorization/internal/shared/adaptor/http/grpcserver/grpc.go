package grpcserver

import (
	"fmt"
	"log"
	"net"

	"github.com/sajad-dev/authservice/authorization/internal/bootstrap"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/http"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs"
	"google.golang.org/grpc"
)

type Grpc struct {
	Port int
}

func NewGrpc(port int) *Grpc {
	return &Grpc{
		Port: port,
	}
}

func (g *Grpc) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.Port))
	if err != nil {
		panic(errs.Err(err))
	}

	gc := grpc.NewServer()

	bootstrap.Boot(gc)

	log.Printf("Run server at port %d \n",g.Port)

	gc.Serve(lis)
}

var _ http.Http = &Grpc{}
