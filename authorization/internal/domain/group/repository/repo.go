package repository

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/group"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/authorize"
)

type GroupRepo struct {
	Authz authorize.Authorize
}

func NewGroupRepo(authz authorize.Authorize) *GroupRepo {
	return &GroupRepo{
		Authz: authz,
	}
}

func (a *GroupRepo) Create(sub string, grp string) (bool, error) {
	return a.Authz.AddGroup(sub, grp)
}

func (a *GroupRepo) Delete(sub string, grp string) (bool, error) {
	return a.Authz.RemoveGroup(sub, grp)
}

func (r *GroupRepo) GetAll() ([][]string, error) {
	return r.Authz.GetAllGroup()

}

var _ group.GroupRepository = &GroupRepo{}
