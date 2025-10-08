package forgetpassword

import (
	"context"

	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/dto/gen/request"
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/dto/gen/response"
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/forgetpasswordproto"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
)

type ForgetPasswordHandler interface {
	Forget(ctx context.Context, req *forgetpasswordproto.ForgetRequest) (*forgetpasswordproto.ForgetResponse, error)
	Reset(ctx context.Context, req *forgetpasswordproto.ResetRequest) (*forgetpasswordproto.ResetResponse, error)
}

type ForgetPasswordService interface {
	Forget(req request.ForgetRequest) (response.ForgetResponse, error)
	Reset(req request.ResetRequest) (response.ResetResponse, error)
}

type ForgetPasswordRepository interface {
	Find(clm string, value string) (*models.Accounts, error)
	FindById(id int) (*models.Accounts, error)
	Update(row *models.Accounts) error
}
