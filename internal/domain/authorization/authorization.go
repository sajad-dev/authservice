package twofactor

import (
	"github.com/sajad-dev/authservice/internal/domain/twofactor/dto/gen/request"
	"github.com/sajad-dev/authservice/internal/domain/twofactor/dto/gen/response"
	"github.com/sajad-dev/authservice/internal/shared/models"
)

type TwoFactorHandler interface {
	AddGroupPolicy(req *twofactorproto.EmailRequest) (*twofactorproto.TwoFactorResponse, error)
	AddPolicy(req *twofactorproto.GoogleRequest) (*twofactorproto.TwoFactorResponse, error)
}

type TwoFactorService interface {
	AddGroupPolicy(req request.GoogleRequest) (response.TwoFactorResponse, error)
	AddPolicy(req request.EmailRequest) (response.TwoFactorResponse, error)
}

type TwoFactorRepo interface {
	FindByCode(code int, codeType string) ([]*models.TwoFactorCode, error)
	FindById(id int) (*models.Accounts, error)
}
