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

func (a *PolicyHndlr) Policy(req *policyproto.PolicyRequest) (*policyproto.Response, error) {

	reqValidation := request.ToRequestPolicy(req)

	if err := globalerr.ValidationErr(a.Validation.Validate(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Policy(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

func (a *PolicyHndlr) Group(req *policyproto.GroupRequest) (*policyproto.Response, error) {

	reqValidation := request.ToRequestGroup(req)

	if err := globalerr.ValidationErr(a.Validation.Validate(reqValidation)); err != nil {
		return nil, err
	}

	res, err := a.Service.Group(*reqValidation)
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

var _ policy.PolicyHandler = &PolicyHndlr{}
