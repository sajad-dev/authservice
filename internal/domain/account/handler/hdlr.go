package handler

import (
	"context"

	"github.com/sajad-dev/authservice/internal/domain/account"
	"github.com/sajad-dev/authservice/internal/domain/account/accountproto"
	"github.com/sajad-dev/authservice/internal/domain/account/dto/request"
	"github.com/sajad-dev/authservice/internal/pkg/errs/grpc/globalerr"
	"github.com/sajad-dev/authservice/internal/validation"
)

type AccountHandler struct {
	accountproto.UnimplementedAccountServer
	Service            account.AccountCURDService
	ValidationInstance *validation.ValidationRequest
}

func NewAccountHandler(svc account.AccountCURDService, vld *validation.ValidationRequest) *AccountHandler {
	return &AccountHandler{Service: svc, ValidationInstance: vld}
}

func (a *AccountHandler) CreateGRPC(ctx context.Context, req *accountproto.CreateRequest) (*accountproto.CreateResponse, error) {

	reqValidation := request.ToRequestCreate(*req)

	res, err := a.Service.CreateService(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

func (a *AccountHandler) UpdateGRPC(ctx context.Context, req *accountproto.UpdateRequest) (*accountproto.UpdateResponse, error) {
	reqValidation := request.ToRequestUpdate(*req)

	res, err := a.Service.UpdateService(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

func (a *AccountHandler) DeleteGRPC(ctx context.Context, req *accountproto.DeleteRequest) (*accountproto.DeleteResponse, error) {
	reqValidation := request.ToRequestDelete(*req)

	res, err := a.Service.DeleteService(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return &accountproto.DeleteResponse{Message: res.Message, Code: int32(res.Code)}, nil
}

func (a *AccountHandler) ReadGRPC(ctx context.Context, req *accountproto.ReadRequest) (*accountproto.ReadResponse, error) {
	reqValidation := request.ToRequestRead(*req)

	res, err := a.Service.ReadService(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}



var _ account.AccountCURDHandler = &AccountHandler{}
