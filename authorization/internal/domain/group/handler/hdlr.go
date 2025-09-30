package handler

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/group"
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/dto/gen/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/groupproto"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/validation"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs/globalerr"
)

type GroupHndlr struct {
	groupproto.UnimplementedGroupServer
	Service    group.GroupService
	Validation validation.Validation
}

func NewGroupHandler(svc group.GroupService, vld validation.Validation) *GroupHndlr {
	return &GroupHndlr{Service: svc, Validation: vld}
}

func (a *GroupHndlr) Create(req *groupproto.CreateRequest) (*groupproto.Response, error) {

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

func (a *GroupHndlr) Delete(req *groupproto.DeleteRequest) (*groupproto.Response, error) {

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

func (a *GroupHndlr) GetAll(req *groupproto.GetAllRequest) (*groupproto.GetAllResponse, error) {
	res, err := a.Service.GetAll()
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

var _ group.GroupHandler = &GroupHndlr{}

