package handler

import (
	"github.com/sajad-dev/authservice/internal/domain/authentication"
	"github.com/sajad-dev/authservice/internal/domain/authentication/authenticationproto"
	"github.com/sajad-dev/authservice/internal/domain/authentication/dto/gen/request"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/validation"
	"github.com/sajad-dev/authservice/internal/shared/errors/errs/globalerr"
)

type AuthenticationHndlr struct {
	authenticationproto.UnimplementedAuthticationServer
	Service    authentication.AuthenticatorService
	Validation validation.Validation
}

func NewAuthenticationHandler(svc authentication.AuthenticatorService, vld validation.Validation) *AuthenticationHndlr {

	return &AuthenticationHndlr{
		Service:    svc,
		Validation: vld,
	}
}

func (a *AuthenticationHndlr) Login(req *authenticationproto.LoginRequest) (*authenticationproto.LoginResponse, error) {

	reqValidation := request.ToRequestLogin(req)

	if err := globalerr.ValidationErr(a.Validation.Validate(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Login(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

func (a *AuthenticationHndlr) Register(req *authenticationproto.RegisterRequest) (*authenticationproto.RegisterResponse, error) {

	reqValidation := request.ToRequestRegister(req)

	if err := globalerr.ValidationErr(a.Validation.Validate(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Register(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

var _ authentication.AuthenticatorHandler = &AuthenticationHndlr{}
