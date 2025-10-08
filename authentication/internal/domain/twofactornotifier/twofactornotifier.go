package twofactornotifier

import (
	"context"

	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/dto/gen/request"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/dto/gen/response"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/twofactornotifierproto"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
)

type TwoFactorNotifierHandler interface {
	NotifierEmail(ctx context.Context, req *twofactornotifierproto.NotifierEmailRequest) (*twofactornotifierproto.TwoFactorNotifierResponse, error)
}

type TwoFactorNotifierService interface {
	NotifierEmail(req request.NotifierEmailRequest) (response.TwoFactorNotifierResponse, error)
}

type TwoFactorNotifierRepository interface {
	RemoveExpierd(id int) error
	FindById(id int) (*models.Accounts, error)
	CreateCode(row *models.TwoFactorCode) error
}
