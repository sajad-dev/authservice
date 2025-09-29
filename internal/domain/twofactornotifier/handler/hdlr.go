package handler

import (
	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier"
	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier/dto/gen/request"
	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier/twofactornotifierproto"

	"github.com/sajad-dev/authservice/internal/shared/adaptor/validation"
	"github.com/sajad-dev/authservice/internal/shared/errors/errs/globalerr"
)

type TwoFactorNotifierHndlr struct {
	twofactornotifierproto.UnimplementedTwofactorNotifierServer
	Service    twofactornotifier.TwoFactorNotifierService
	Validation validation.Validation
}

func NewTwoFactorNotifierHandler(svc twofactornotifier.TwoFactorNotifierService, vld validation.Validation) *TwoFactorNotifierHndlr {

	return &TwoFactorNotifierHndlr{
		Service:    svc,
		Validation: vld,
	}
}

func (a *TwoFactorNotifierHndlr) NotifierEmail(req *twofactornotifierproto.NotifierEmailRequest) (*twofactornotifierproto.TwoFactorNotifierResponse, error) {

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

var _ twofactornotifier.TwoFactorNotifierHandler = &TwoFactorNotifierHndlr{}
