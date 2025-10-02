package grpcserver

import (
	"net"

	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs"
	"google.golang.org/grpc"
)

type Grpc struct {
	Port int
}

func (g *Grpc) Run() {
	lis, err := net.Listen("tcp", "3000")
	if err != nil {
		panic(errs.Err(err))
	}

	gc := grpc.NewServer()
	
}
