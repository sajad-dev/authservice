package account

import (
	"context"

	"github.com/sajad-dev/authservice/authentication/internal/domain/account/accountproto"
	"github.com/sajad-dev/authservice/authentication/internal/domain/account/dto/gen/request"
	"github.com/sajad-dev/authservice/authentication/internal/domain/account/dto/gen/response"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
)

type AccountCURDHandler interface {
	Create(ctx context.Context, req *accountproto.CreateRequest) (*accountproto.CreateResponse, error)
	Update(ctx context.Context, req *accountproto.UpdateRequest) (*accountproto.UpdateResponse, error)
	Delete(ctx context.Context, req *accountproto.DeleteRequest) (*accountproto.DeleteResponse, error)
	Read(ctx context.Context, req *accountproto.ReadRequest) (*accountproto.ReadResponse, error)
}

type AccountCURDService interface {
	Create(req request.CreateRequest) (response.CreateResponse, error)
	Update(req request.UpdateRequest) (response.UpdateResponse, error)
	Delete(req request.DeleteRequest) (response.DeleteResponse, error)
	Read(req request.ReadRequest) (response.ReadResponse, error)
}

type AccountCURDRepository interface {
	Create(req *models.Accounts) error
	Update(req *models.Accounts, id int) error
	Delete(id int) error
	Read(id int) (*models.Accounts, error)
}
