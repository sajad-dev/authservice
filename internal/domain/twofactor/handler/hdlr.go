package handler

import (
	"github.com/sajad-dev/authservice/internal/domain/forgetpassword"
	"github.com/sajad-dev/authservice/internal/domain/forgetpassword/dto/gen/request"
	"github.com/sajad-dev/authservice/internal/domain/forgetpassword/forgetpasswordproto"

	"github.com/sajad-dev/authservice/internal/shared/adaptor/validation"
	"github.com/sajad-dev/authservice/internal/shared/errors/errs/globalerr"
)

type TwoFactorHndlr struct {
	forgetpasswordproto.UnimplementedTwoFactorServer
	Service    forgetpassword.TwoFactorService
	Validation validation.Validation
}

func NewTwoFactorHandler(svc forgetpassword.TwoFactorService, vld validation.Validation) *TwoFactorHndlr {

	return &TwoFactorHndlr{
		Service:    svc,
		Validation: vld,
	}
}

func (a *TwoFactorHndlr) Forget(req *forgetpasswordproto.ForgetRequest) (*forgetpasswordproto.ForgetResponse, error) {

	reqValidation := request.ToRequestForget(req)

	if err := globalerr.ValidationErr(a.Validation.Validate(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Forget(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

func (a *TwoFactorHndlr) Reset(req *forgetpasswordproto.ResetRequest) (*forgetpasswordproto.ResetResponse, error) {

	reqValidation := request.ToRequestReset(req)

	if err := globalerr.ValidationErr(a.Validation.Validate(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Reset(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

var _ forgetpassword.TwoFactorHandler = &TwoFactorHndlr{}
