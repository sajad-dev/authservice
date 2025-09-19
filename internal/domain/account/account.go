package account

import (
	"context"

	"github.com/sajad-dev/authservice/internal/domain/account/accountproto"
	"github.com/sajad-dev/authservice/internal/domain/account/dto/request"
	"github.com/sajad-dev/authservice/internal/domain/account/dto/response"
	"github.com/sajad-dev/authservice/internal/domain/account/models"
)

type AccountCURDHandler interface {
	CreateGRPC(ctx context.Context, req *accountproto.CreateRequest) (*accountproto.CreateResponse, error)
	UpdateGRPC(ctx context.Context, req *accountproto.UpdateRequest) (*accountproto.UpdateResponse, error)
	DeleteGRPC(ctx context.Context, req *accountproto.DeleteRequest) (*accountproto.DeleteResponse, error)
	ReadGRPC(ctx context.Context, req *accountproto.ReadRequest) (*accountproto.ReadResponse, error)
}

type AccountCURDService interface {
	CreateService(req request.CreateRequest) (response.CreateResponse, error)
	UpdateService(req request.UpdateRequest) (response.UpdateResponse, error)
	DeleteService(req request.DeleteRequest) (response.DeleteResponse, error)
	ReadService(req request.ReadRequest) (response.ReadResponse, error)
}

type AccountCURDRepositories interface {
	Create(req *models.Accounts) error
	Update(req *models.Accounts) error
	Delete(id int) error
	Read(id int) (*models.Accounts, error)
}
