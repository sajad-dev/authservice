package group

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/dto/gen/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/dto/gen/response"
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/groupproto"
)

type GroupHandler interface {
	Create(req *groupproto.CreateRequest) (*groupproto.Response, error)
	Delete(req *groupproto.DeleteRequest) (*groupproto.Response, error)
	GetAll(req *groupproto.GetAllRequest) (*groupproto.GetAllResponse, error)
}

type GroupService interface {
	Create(req request.CreateRequest) (response.Response, error)
	Delete(req request.DeleteRequest) (response.Response, error)
	GetAll() (response.GetAllResponse, error)
}

type GroupRepository interface {
	Create(sub string, grp string) (bool, error)
	Delete(sub string, grp string) (bool, error)
	GetAll() ([][]string, error)
}
