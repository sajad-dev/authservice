package handler

import (
	"context"

	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/dto/gen/request"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/twofactorproto"
	"github.com/sajad-dev/authservice/authentication/internal/shared/validation"
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs/globalerr"
)

type TwoFactorHdlr struct {
	twofactorproto.UnimplementedTwofactoryServer
	Service    twofactor.TwoFactorService
	Validation validation.Validation
}

func NewTwoFactorHdlr(svc twofactor.TwoFactorService, vld validation.Validation) *TwoFactorHdlr {

	return &TwoFactorHdlr{
		Service:    svc,
		Validation: vld,
	}
}

func (a *TwoFactorHdlr) Email(ctx context.Context, req *twofactorproto.EmailRequest) (*twofactorproto.TwoFactorResponse, error) {

	reqValidation := request.ToRequestEmail(req)

	if err := globalerr.ValidationErr(a.Validation.Verify(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Email(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

func (a *TwoFactorHdlr) Google(ctx context.Context, req *twofactorproto.GoogleRequest) (*twofactorproto.TwoFactorResponse, error) {

	reqValidation := request.ToRequestGoogle(req)

	if err := globalerr.ValidationErr(a.Validation.Verify(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Google(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

var _ twofactor.TwoFactorHandler = &TwoFactorHdlr{}
