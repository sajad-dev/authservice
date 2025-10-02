package repository

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize"
	authz "github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/authorize"
)

type AuthorizeRepo struct {
	Authz authz.Authorize
}

func NewAuthorizeRepo(authz authz.Authorize) *AuthorizeRepo {
	return &AuthorizeRepo{
		Authz: authz,
	}
}

func (a *AuthorizeRepo) Check(sub string, obj string, act string) (bool, error) {
	return a.Authz.Verify(sub, obj, act)
}

var _ authorize.AuthorizeRepository = &AuthorizeRepo{}
