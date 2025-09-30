package policy

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/dto/gen/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/dto/gen/response"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/policyproto"
)

type PolicyHandler interface {
	Create(req *policyproto.CreateRequest) (*policyproto.Response, error)
	Delete(req *policyproto.DeleteRequest) (*policyproto.Response, error)
	GetAll(req *policyproto.GetAllRequest) (*policyproto.GetAllResponse, error)
}

type PolicyService interface {
	Create(req request.CreateRequest) (response.Response, error)
	Delete(req request.DeleteRequest) (response.Response, error)
	GetAll() (response.GetAllResponse, error)
}

type PolicyRepository interface {
	Create(sub string, obj string, act string) (bool, error)
	Delete(sub string, obj string, act string) (bool, error)
	GetAll() ([][]string, error)
}
