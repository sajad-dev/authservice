package service

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/dto/gen/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/dto/gen/response"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs"
)

type PolicySvc struct {
	Repo policy.PolicyRepository
}

func NewPolicySvc(repo policy.PolicyRepository) *PolicySvc {
	return &PolicySvc{
		Repo: repo,
	}
}

func (a *PolicySvc) Group(req *request.GroupRequest) (response.Response, error) {

	ok, err := a.Repo.AddGroup(req.Subject, req.Group)
	if err != nil {
		return response.Response{}, errs.Err(err)
	}

	if !ok {
		return response.Response{
			Msg:  messages.ERR_ADD_GROUP_FAILED,
			Code: statuscode.VALIDATION_ERR,
		}, nil
	}

	return response.Response{
		Msg:  messages.SUCCESS_GROUP_ADDED,
		Code: statuscode.SUCCESSFUL,
	}, nil

}
func (a *PolicySvc) Policy(req *request.PolicyRequest) (response.Response, error) {
	ok, err := a.Repo.AddPolicy(req.Subject, req.Group, req.Action)
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

var _ policy.PolicyService = &PolicySvc{}
