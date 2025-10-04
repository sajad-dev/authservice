package handler

import (
	"context"

	"github.com/sajad-dev/authservice/authorization/internal/domain/policy"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/dto/gen/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/policyproto"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/validation"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs/globalerr"
)

type PolicyHdlr struct {
	policyproto.UnimplementedPolicyServer
	Service    policy.PolicyService
	Validation validation.Validation
}

func NewPolicyHdlr(svc policy.PolicyService, vld validation.Validation) *PolicyHdlr {
	return &PolicyHdlr{Service: svc, Validation: vld}
}

func (a *PolicyHdlr) Create(ctx context.Context,req *policyproto.CreateRequest) (*policyproto.Response, error) {

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

func (a *PolicyHdlr) Delete(ctx context.Context,req *policyproto.DeleteRequest) (*policyproto.Response, error) {

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

func (a *PolicyHdlr) GetAll(ctx context.Context,req *policyproto.GetAllRequest) (*policyproto.GetAllResponse, error) {
	res, err := a.Service.GetAll()
	if err = globalerr.ServerErr(err); err != nil {
		return nil, err
	}

	return res.ToProto(), nil
}

var _ policy.PolicyHandler = &PolicyHdlr{}
