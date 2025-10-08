package authentication

import (
	"context"

	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/authenticationproto"
	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/dto/gen/request"
	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/dto/gen/response"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
)

type AuthenticatorHandler interface {
	Login(ctx *context.Context, req *authenticationproto.LoginRequest) (*authenticationproto.LoginResponse, error)
	Register(ctx *context.Context,req *authenticationproto.RegisterRequest) (*authenticationproto.RegisterResponse, error)
}
type AuthenticatorService interface {
	Login(req request.LoginRequest) (response.LoginResponse, error)
	Register(req request.RegisterRequest) (response.RegisterResponse, error)
}

type AuthenticatorRepository interface {
	Find(clm string, value string) (*models.Accounts, error)
	Create(row *models.Accounts) error
}
