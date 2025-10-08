package handler

import (
	"context"

	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword"
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/dto/gen/request"
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/forgetpasswordproto"

	"github.com/sajad-dev/authservice/authentication/internal/shared/validation"
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs/globalerr"
)

type ForgetPasswordHdlr struct {
	forgetpasswordproto.UnimplementedForgetPasswordServer
	Service    forgetpassword.ForgetPasswordService
	Validation validation.Validation
}

func NewForgetPasswordHdlr(svc forgetpassword.ForgetPasswordService, vld validation.Validation) *ForgetPasswordHdlr {

	return &ForgetPasswordHdlr{
		Service:    svc,
		Validation: vld,
	}
}

func (a *ForgetPasswordHdlr) Forget(ctx *context.Context, req *forgetpasswordproto.ForgetRequest) (*forgetpasswordproto.ForgetResponse, error) {

	reqValidation := request.ToRequestForget(req)

	if err := globalerr.ValidationErr(a.Validation.Verify(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Forget(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

func (a *ForgetPasswordHdlr) Reset(ctx *context.Context, req *forgetpasswordproto.ResetRequest) (*forgetpasswordproto.ResetResponse, error) {

	reqValidation := request.ToRequestReset(req)

	if err := globalerr.ValidationErr(a.Validation.Verify(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Reset(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

var _ forgetpassword.ForgetPasswordHandler = &ForgetPasswordHdlr{}
