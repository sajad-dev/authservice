package twofactor

import (
	"github.com/sajad-dev/authservice/internal/domain/forgetpassword/dto/gen/request"
	"github.com/sajad-dev/authservice/internal/domain/forgetpassword/dto/gen/response"
	"github.com/sajad-dev/authservice/internal/domain/forgetpassword/forgetpasswordproto"
	"github.com/sajad-dev/authservice/internal/shared/models"
)

type TwoFactorHandler interface {
	Forget(req *forgetpasswordproto.ForgetRequest) (*forgetpasswordproto.ForgetResponse, error)
	Reset(req *forgetpasswordproto.ResetRequest) (*forgetpasswordproto.ResetResponse, error)
}

type TwoFactorService interface {
	Forget(req request.ForgetRequest) (response.ForgetResponse, error)
	Reset(req request.ResetRequest) (response.ResetResponse, error)
}

type TwoFactorRepo interface {
	FindByCode(code int, codeType string) ([]*models.TwoFactorCode, error)
	FindById(id int) (*models.Accounts, error)
}
