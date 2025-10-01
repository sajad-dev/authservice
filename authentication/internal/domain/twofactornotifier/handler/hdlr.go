package handler

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/dto/gen/request"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/twofactornotifierproto"

	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/validation"
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs/globalerr"
)

type TwoFactorNotifierHdlr struct {
	twofactornotifierproto.UnimplementedTwofactorNotifierServer
	Service    twofactornotifier.TwoFactorNotifierService
	Validation validation.Validation
}

func NewTwoFactorNotifierHdlr(svc twofactornotifier.TwoFactorNotifierService, vld validation.Validation) *TwoFactorNotifierHdlr {

	return &TwoFactorNotifierHdlr{
		Service:    svc,
		Validation: vld,
	}
}

func (a *TwoFactorNotifierHdlr) NotifierEmail(req *twofactornotifierproto.NotifierEmailRequest) (*twofactornotifierproto.TwoFactorNotifierResponse, error) {

	reqValidation := request.ToRequestNotifierEmail(req)

	if err := globalerr.ValidationErr(a.Validation.Validate(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.NotifierEmail(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

var _ twofactornotifier.TwoFactorNotifierHandler = &TwoFactorNotifierHdlr{}
