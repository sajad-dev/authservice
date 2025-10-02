package bootstrap

import (
	authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/handler"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/service"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/crypto/hs256"

	"google.golang.org/grpc"
)

func Boot(gc *grpc.Server) {
	// authzInstans := casbin.NewEnforcer()

	authz.RegisterAuthorizationServer(gc, handler.NewAuthorizeHdlr(
		service.NewAuthorizeSvc(
			nil,
			hs256.NewJWT([]byte("hi")),
		),
	))
}
