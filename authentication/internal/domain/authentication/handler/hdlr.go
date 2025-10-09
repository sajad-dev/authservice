package handler

import (
	"context"
	"log"

	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication"
	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/authenticationproto"
	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/dto/gen/request"
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs/globalerr"
	"github.com/sajad-dev/authservice/authentication/internal/shared/validation"
)

type AuthenticationHdlr struct {
	authenticationproto.UnimplementedAuthenticationServer
	Service    authentication.AuthenticatorService
	Validation validation.Validation
}

func NewAuthenticationHdlr(svc authentication.AuthenticatorService, vld validation.Validation) *AuthenticationHdlr {

	return &AuthenticationHdlr{
		Service:    svc,
		Validation: vld,
	}
}

func (a *AuthenticationHdlr) Login(ctx context.Context, req *authenticationproto.LoginRequest) (*authenticationproto.LoginResponse, error) {
	log.Println("req")

	reqValidation := request.ToRequestLogin(req)

	if err := globalerr.ValidationErr(a.Validation.Verify(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Login(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

func (a *AuthenticationHdlr) Register(ctx context.Context, req *authenticationproto.RegisterRequest) (*authenticationproto.RegisterResponse, error) {

	reqValidation := request.ToRequestRegister(req)

	if err := globalerr.ValidationErr(a.Validation.Verify(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Register(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

var _ authentication.AuthenticatorHandler = &AuthenticationHdlr{}
