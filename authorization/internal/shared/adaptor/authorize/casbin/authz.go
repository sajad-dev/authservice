package casbin

import (
	authz "github.com/casbin/casbin/v2"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/authorize"
)

type Casbin struct {
	enforcer *authz.Enforcer
}

func NewCasbin(enf *authz.Enforcer) *Casbin {
	return &Casbin{
		enforcer: enf,
	}
}

func (c *Casbin) AddPolicy(policy ...interface{}) (bool, error) {
	return c.enforcer.AddPolicy(policy...)
}

func (c *Casbin) RemovePolicy(policy ...interface{}) (bool, error) {
	return c.enforcer.RemovePolicy(policy...)
}

func (c *Casbin) GetAllPolicy() ([][]string, error) {
	return c.enforcer.GetPolicy()
}

func (c *Casbin) AddGroup(policy ...interface{}) (bool, error) {
	return c.enforcer.AddGroupingPolicy(policy...)
}

func (c *Casbin) RemoveGroup(policy ...interface{}) (bool, error) {
	return c.enforcer.RemoveGroupingPolicy(policy...)
}

func (c *Casbin) GetAllGroup() ([][]string, error) {
	return c.enforcer.GetGroupingPolicy()
}

func (c *Casbin) Verify(policy ...interface{}) (bool, error) {
	return c.enforcer.Enforce(policy...)
}

var _ authorize.Authorize = &Casbin{}
