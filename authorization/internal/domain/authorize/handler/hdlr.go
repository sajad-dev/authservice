package handler

import (
	"context"
	"errors"
	"log"

	authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/dto/request"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs/globalerr"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

type AuthorizeHdlr struct {
	Service authorize.AuthorizeService
}

func NewAuthorizeHdlr(svc authorize.AuthorizeService) *AuthorizeHdlr {
	return &AuthorizeHdlr{Service: svc}
}

func (a *AuthorizeHdlr) Check(ctx context.Context, req *authz.CheckRequest) (*authz.CheckResponse, error) {
	log.Println("I get that hahaha")
	httpReq := req.GetAttributes().GetRequest().GetHttp()
	if httpReq == nil {
		return nil, globalerr.ServerErr(errors.New("missing http request in envoy check request"))
	}

	headers := make(map[string]string)
	for k, v := range httpReq.GetHeaders() {
		headers[k] = v
	}

	serviceReq := request.AuthorizeRequest{
		Headers:  headers,
		Host:     httpReq.GetHost(),
		Body:     httpReq.GetBody(),
		Protocol: httpReq.GetProtocol(),
		Method:   httpReq.GetMethod(),
		Path:     httpReq.GetPath(),
		RawBody:  []byte(httpReq.GetBody()),
	}

	resp, err := a.Service.Check(serviceReq)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	if resp.Code == statuscode.SUCCESSFUL {
		return &authz.CheckResponse{
			HttpResponse: &authz.CheckResponse_OkResponse{},
			Status: &status.Status{
				Code: int32(code.Code_OK),
			},
		}, nil
	}

	return &authz.CheckResponse{
		Status: &status.Status{
			Code: int32(code.Code_PERMISSION_DENIED),
		},
	}, nil
}

var _ authz.AuthorizationServer = &AuthorizeHdlr{}
