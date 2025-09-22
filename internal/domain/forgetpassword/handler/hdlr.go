package handler

import (
	"github.com/sajad-dev/authservice/internal/domain/forgetpassword"
	"github.com/sajad-dev/authservice/internal/domain/forgetpassword/dto/gen/request"
	"github.com/sajad-dev/authservice/internal/domain/forgetpassword/forgetpasswordproto"

	"github.com/sajad-dev/authservice/internal/shared/adaptor/validation"
	"github.com/sajad-dev/authservice/internal/shared/errors/errs/globalerr"
)

type ForgetPasswordHndlr struct {
	forgetpasswordproto.UnimplementedForgetPasswordServer
	Service    forgetpassword.ForgetPasswordService
	Validation validation.Validation
}

func NewForgetPasswordHandler(svc forgetpassword.ForgetPasswordService, vld validation.Validation) *ForgetPasswordHndlr {

	return &ForgetPasswordHndlr{
		Service:    svc,
		Validation: vld,
	}
}

func (a *ForgetPasswordHndlr) Forget(req *forgetpasswordproto.ForgetRequest) (*forgetpasswordproto.ForgetResponse, error) {

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

func (a *ForgetPasswordHndlr) Reset(req *forgetpasswordproto.ResetRequest) (*forgetpasswordproto.ResetResponse, error) {

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

var _ forgetpassword.ForgetPasswordHandler = &ForgetPasswordHndlr{}
