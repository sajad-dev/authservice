package handler

import (
	"github.com/sajad-dev/authservice/internal/domain/twofactor"
	"github.com/sajad-dev/authservice/internal/domain/twofactor/dto/gen/request"
	"github.com/sajad-dev/authservice/internal/domain/twofactor/twofactorproto"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/validation"
	"github.com/sajad-dev/authservice/internal/shared/errors/errs/globalerr"
)

type TwoFactorHndlr struct {
	twofactorproto.UnimplementedTwofactoryServer
	Service    twofactor.TwoFactorService
	Validation validation.Validation
}

func NewTwoFactorHandler(svc twofactor.TwoFactorService, vld validation.Validation) *TwoFactorHndlr {

	return &TwoFactorHndlr{
		Service:    svc,
		Validation: vld,
	}
}

func (a *TwoFactorHndlr) Email(req *twofactorproto.EmailRequest) (*twofactorproto.TwoFactorResponse, error) {

	reqValidation := request.ToRequestEmail(req)

	if err := globalerr.ValidationErr(a.Validation.Validate(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Email(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

func (a *TwoFactorHndlr) Google(req *twofactorproto.GoogleRequest) (*twofactorproto.TwoFactorResponse, error) {

	reqValidation := request.ToRequestGoogle(req)

	if err := globalerr.ValidationErr(a.Validation.Validate(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Google(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

var _ twofactor.TwoFactorHandler = &TwoFactorHndlr{}
