package handler

import (
	"context"
	"log"
	"strings"

	authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type AuthorizeHdlr struct {
	Service    authorize.AuthorizeService
}

func NewAuthorizeHdlr(svc authorize.AuthorizeService) *AuthorizeHdlr {
	return &AuthorizeHdlr{Service: svc}
}

func (a *AuthorizeHdlr) Check(ctx context.Context, req *authz.CheckRequest) (*authz.CheckResponse, error) {
	authorization := req.Attributes.Request.Http.Headers["authorization"]
	log.Println(authorization)

	// role := req.Attributes.ContextExtensions["role"]

	extracted := strings.Fields(authorization)
	if len(extracted) == 2 && extracted[0] == "Bearer" {
		return &authz.CheckResponse{
			HttpResponse: &authz.CheckResponse_OkResponse{},
			Status: &status.Status{
				Code: int32(code.Code_OK),
			},
		}, nil
	}

	return &auth.CheckResponse{
		Status: &status.Status{
			Code: int32(code.Code_PERMISSION_DENIED),
		},
	}, nil
}

var _ authz.AuthorizationServer = &AuthorizeHdlr{}
