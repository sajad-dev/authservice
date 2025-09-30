package handler

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/dto/gen/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/policyproto"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/validation"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs/globalerr"
)

type PolicyHndlr struct {
	policyproto.UnimplementedPolicyServer
	Service    policy.PolicyService
	Validation validation.Validation
}

func NewPolicyHandler(svc policy.PolicyService, vld validation.Validation) *PolicyHndlr {
	return &PolicyHndlr{Service: svc, Validation: vld}
}

func (a *PolicyHndlr) Create(req *policyproto.CreateRequest) (*policyproto.Response, error) {

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

func (a *PolicyHndlr) Delete(req *policyproto.DeleteRequest) (*policyproto.Response, error) {

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

func (a *PolicyHndlr) GetAll(req *policyproto.GetAllRequest) (*policyproto.GetAllResponse, error) {
	res, err := a.Service.GetAll()
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

var _ policy.PolicyHandler = &PolicyHndlr{}
