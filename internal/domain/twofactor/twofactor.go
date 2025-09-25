package twofactor

import (
	"github.com/sajad-dev/authservice/internal/domain/twofactor/dto/gen/request"
	"github.com/sajad-dev/authservice/internal/domain/twofactor/dto/gen/response"
	"github.com/sajad-dev/authservice/internal/domain/twofactor/twofactorproto"
	"github.com/sajad-dev/authservice/internal/shared/models"
)

type TwoFactorHandler interface {
	Email(req *twofactorproto.EmailRequest) (*twofactorproto.TwoFactorResponse, error)
	Google(req *twofactorproto.GoogleRequest) (*twofactorproto.TwoFactorResponse, error)
}

type TwoFactorService interface {
	Google(req request.GoogleRequest) (response.TwoFactorResponse, error)
	Email(req request.EmailRequest) (response.TwoFactorResponse, error)
}

type TwoFactorRepo interface {
	FindByCode(code int, codeType string) ([]*models.TwoFactorCode, error)
	FindById(id int) (*models.Accounts, error)
}
