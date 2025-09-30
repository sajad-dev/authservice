package handler

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/account"
	"github.com/sajad-dev/authservice/authentication/internal/domain/account/accountproto"
	"github.com/sajad-dev/authservice/authentication/internal/domain/account/dto/gen/request"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/validation"
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs/globalerr"
)

type AccountHndlr struct {
	accountproto.UnimplementedAccountServer
	Service    account.AccountCURDService
	Validation validation.Validation
}

func NewAccountHandler(svc account.AccountCURDService, vld validation.Validation) *AccountHndlr {
	return &AccountHndlr{Service: svc, Validation: vld}
}

func (a *AccountHndlr) Create(req *accountproto.CreateRequest) (*accountproto.CreateResponse, error) {

	reqValidation := request.ToRequestCreate(req)

	if err := globalerr.ValidationErr(a.Validation.Validate(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Create(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

func (a *AccountHndlr) Update(req *accountproto.UpdateRequest) (*accountproto.UpdateResponse, error) {
	reqValidation := request.ToRequestUpdate(req)

	if err := globalerr.ValidationErr(a.Validation.Validate(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Update(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

func (a *AccountHndlr) Delete(req *accountproto.DeleteRequest) (*accountproto.DeleteResponse, error) {
	reqValidation := request.ToRequestDelete(req)

	if err := globalerr.ValidationErr(a.Validation.Validate(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Delete(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

func (a *AccountHndlr) Read(req *accountproto.ReadRequest) (*accountproto.ReadResponse, error) {
	reqValidation := request.ToRequestRead(req)

	if err := globalerr.ValidationErr(a.Validation.Validate(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Read(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

var _ account.AccountCURDHandler = &AccountHndlr{}
