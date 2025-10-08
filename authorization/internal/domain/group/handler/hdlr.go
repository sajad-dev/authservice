package handler

import (
	"context"

	"github.com/sajad-dev/authservice/authorization/internal/domain/group"
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/dto/gen/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/groupproto"
	"github.com/sajad-dev/authservice/authorization/internal/shared/validation"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs/globalerr"
)

type GroupHdlr struct {
	groupproto.UnimplementedGroupServer
	Service    group.GroupService
	Validation validation.Validation
}

func NewGroupHdlr(svc group.GroupService, vld validation.Validation) *GroupHdlr {
	return &GroupHdlr{Service: svc, Validation: vld}
}

func (a *GroupHdlr) Create(ctx context.Context, req *groupproto.CreateRequest) (*groupproto.Response, error) {

	reqValidation := request.ToRequestCreate(req)

	if err := globalerr.ValidationErr(a.Validation.Verify(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Create(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

func (a *GroupHdlr) Delete(ctx context.Context, req *groupproto.DeleteRequest) (*groupproto.Response, error) {

	reqValidation := request.ToRequestDelete(req)

	if err := globalerr.ValidationErr(a.Validation.Verify(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Delete(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

func (a *GroupHdlr) GetAll(ctx context.Context, req *groupproto.GetAllRequest) (*groupproto.GetAllResponse, error) {
	res, err := a.Service.GetAll()
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

var _ group.GroupHandler = &GroupHdlr{}
