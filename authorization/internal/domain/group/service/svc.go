package service

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/group"
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/dto/gen/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/dto/gen/response"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs"
)

type GroupSvc struct {
	Repo group.GroupRepository
}

func NewGroupSvc(repo group.GroupRepository) *GroupSvc {
	return &GroupSvc{
		Repo: repo,
	}
}

func (a *GroupSvc) Create(req request.CreateRequest) (response.Response, error) {
	ok, err := a.Repo.Create(req.Subject, req.Group)
	if err != nil {
		return response.Response{}, errs.Err(err)
	}

	if !ok {
		return response.Response{
			Msg:  messages.ERR_ADD_POLICY_FAILED,
			Code: statuscode.VALIDATION_ERR,
		}, errs.Err(err)
	}

	return response.Response{
		Msg:  messages.SUCCESS_POLICY_ADDED,
		Code: statuscode.SUCCESSFUL,
	}, nil
}

func (a *GroupSvc) Delete(req request.DeleteRequest) (response.Response, error) {
	ok, err := a.Repo.Delete(req.Subject, req.Group)
	if err != nil {
		return response.Response{}, errs.Err(err)
	}

	if !ok {
		return response.Response{
			Msg:  messages.ERR_ADD_POLICY_FAILED,
			Code: statuscode.VALIDATION_ERR,
		}, errs.Err(err)
	}

	return response.Response{
		Msg:  messages.SUCCESS_POLICY_ADDED,
		Code: statuscode.SUCCESSFUL,
	}, nil
}

func (a *GroupSvc) GetAll() (response.GetAllResponse, error) {
	_, err := a.Repo.GetAll()
	if err != nil {
		return response.GetAllResponse{}, errs.Err(err)
	}

	return response.GetAllResponse{
		Msg:  messages.SUCCESS_POLICY_ADDED,
		Code: statuscode.SUCCESSFUL,
	}, nil
}

var _ group.GroupService = &GroupSvc{}

