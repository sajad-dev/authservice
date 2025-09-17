package handler

import (
	"context"

	"github.com/sajad-dev/authservice/internal/domain/account"
	"github.com/sajad-dev/authservice/internal/domain/account/accountproto"
	"github.com/sajad-dev/authservice/internal/domain/account/dto/request"
)

type AccountHandler struct {
	accountproto.UnimplementedAccountServer
	Service            service.AccountCURDService
	ValidationInstance *validation.ValidationRequest
}

func (a *AccountHandler) CreateGRPC(ctx context.Context, req *accountproto.CreateRequest) (*accountproto.CreateReply, error) {

	reqValidation := request.ToRequestCreate(*req)

	res, err := a.Service.CreateService(*reqValidation)
	if err = apperrors.CreateServiceErr(err); err != nil {
		return nil, err
	}

	return &accountproto.CreateReply{Message: res.Message, Code: int32(res.Code)}, nil
}

func (a *AccountHandler) UpdateGRPC(ctx context.Context, req *accountproto.UpdateRequest) (*accountproto.UpdateReply, error) {
	reqValidation := updatereq.NewUpdateRequest(
		updatereq.WithID(int(req.ID)),
		updatereq.WithFirstName(req.FirstName),
		updatereq.WithLastName(req.LastName),
		updatereq.WithEmail(req.Email),
		updatereq.WithSMS(req.SMS),
		updatereq.WithUsername(req.Username),
		updatereq.WithTwoFactor(req.TwoFactor),
		updatereq.WithPassword(req.Password),
	)

	var err error
	reqValidation.Password, err = crypto.SumSHA256([]byte(req.Password))

	if err := globalerr.ValidationErr(a.ValidationInstance.ValidationStruct(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.UpdateService(*reqValidation)
	if err = apperrors.UpdateServiceErr(err); err != nil {
		return nil, err
	}

	return &accountproto.UpdateReply{Message: res.Message, Code: int32(res.Code)}, nil
}

func (a *AccountHandler) DeleteGRPC(ctx context.Context, req *accountproto.DeleteRequest) (*accountproto.DeleteReply, error) {
	reqValidation := deletereq.NewDeleteRequest(
		deletereq.WithID(int(req.ID)),
	)

	if err := globalerr.ValidationErr(a.ValidationInstance.ValidationStruct(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.DeleteService(*reqValidation)
	if err = apperrors.DeleteServiceErr(err); err != nil {
		return nil, err
	}

	return &accountproto.DeleteReply{Message: res.Message, Code: int32(res.Code)}, nil
}

func (a *AccountHandler) ReadGRPC(ctx context.Context, req *accountproto.ReadRequest) (*accountproto.ReadReply, error) {
	reqValidation := readreq.NewReadRequest(
		readreq.WithID(int(req.ID)),
	)

	if err := globalerr.ValidationErr(a.ValidationInstance.ValidationStruct(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.ReadService(*reqValidation)
	if err = apperrors.ReadServiceErr(err); err != nil {
		return nil, err
	}

	var reply accountproto.ReadReply
	if err := globalerr.ServerErr(adaptor.Adaptor(&reply, res)); err != nil {
		return nil, err
	}

	return &reply, nil
}

func NewAccountHandler(svc service.AccountCURDService, vld *validation.ValidationRequest) *AccountHandler {
	return &AccountHandler{Service: svc, ValidationInstance: vld}
}

var _ account.AccountCURDHandler = &AccountHandler{}
