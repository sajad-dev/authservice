package policy

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/dto/gen/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/dto/gen/response"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/policyproto"
)

type PolicyHandler interface {
	Group(req *policyproto.GroupRequest) (policyproto.Response, error)
	Policy(req *policyproto.PolicyRequest) (policyproto.Response, error)
}

type PolicyService interface {
	Group(req request.GroupRequest) (response.Response, error)
	Policy(req request.PolicyRequest) (response.Response, error)
}

type PolicyRepository interface {
	AddGroup(sub string, grp string) (response.Response, error)
	AddPolicy(sub string, grp string, act string) (response.Response, error)
}
