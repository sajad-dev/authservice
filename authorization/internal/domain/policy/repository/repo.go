package repository

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/authorize"
)

type PolicyRepo struct {
	Authz authorize.Authorize
}

func NewPolicyRepo(authz authorize.Authorize) *PolicyRepo {
	return &PolicyRepo{
		Authz: authz,
	}
}

func (a *PolicyRepo) Create(sub string, obj string, act string) (bool,error) {
	return a.Authz.AddPolicy(sub,obj,act)
}

func (a *PolicyRepo) Delete(sub string, obj string, act string) (bool,error) {
	return a.Authz.RemovePolicy(sub,obj,act)
}

func (r *PolicyRepo) GetAll() ([][]string, error) {
	return r.Authz.GetAllGroup()

}

var _ policy.PolicyRepository = &PolicyRepo{}
