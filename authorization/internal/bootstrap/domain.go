package bootstrap

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/groupproto"
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/handler"
	"google.golang.org/grpc"
)

func Handle(gc *grpc.Server) {
	groupproto.RegisterGroupServer(gc, &handler.GroupHdlr{})

}
