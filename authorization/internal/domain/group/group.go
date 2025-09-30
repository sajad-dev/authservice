package policy

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/dto/gen/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/dto/gen/response"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/policyproto"
)

type GroupHandler interface {
	Create(req *policyproto.CreateRequest) (policyproto.Response, error)
	Delete(req *policyproto.DeleteRequest) (policyproto.Response, error)
	GetAll(req *policyproto.GetAllRequest) (policyproto.Response, error)
}

type GroupService interface {
	Create(req request.GroupRequest) (response.Response, error)
	Delete(req request.GroupRequest) (response.Response, error)
	GetAll(req request.GroupRequest) (response.Response, error)
}

type GroupRepository interface {
	Create(sub string, grp string) (bool, error)
	Delete(sub string, grp string) (bool, error)
	GetAll() ([][]string, error)
}
