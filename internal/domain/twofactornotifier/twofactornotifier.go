package twofactornotifier

import (
	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier/dto/gen/request"
	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier/dto/gen/response"
	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier/twofactornotifierproto"
	"github.com/sajad-dev/authservice/internal/shared/models"
)

type TwoFactorNotifierHandler interface {
	NotifierEmail(req *twofactornotifierproto.NotifierEmailRequest) (*twofactornotifierproto.TwoFactorNotifierResponse, error)
}

type TwoFactorNotifierService interface {
	NotifierEmail(req request.NotifierEmailRequest) (response.TwoFactorNotifierResponse, error)
}

type TwoFactorNotifierRepository interface {
	RemoveExpierd(id int) error
	FindById(id int) (*models.Accounts, error)
	CreateCode(row *models.TwoFactorCode) error
}
