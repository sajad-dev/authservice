package repositories

import (
	"github.com/sajad-dev/authservice/internal/domain/authentication"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/sqldb"
	"github.com/sajad-dev/authservice/internal/shared/models"
)

type AuthenticationRepo struct {
	DB sqldb.SqlDB[*models.Accounts]
}

func NewAuthenticationRepo(db sqldb.SqlDB[*models.Accounts]) *AuthenticationRepo {
	return &AuthenticationRepo{DB: db}
}

func (a AuthenticationRepo) Create(req *models.Accounts) error {
	return a.DB.Create(req)
}

func (s AuthenticationRepo) FindWithField(params *models.Accounts) ([]*models.Accounts, error) {
	return s.DB.FindWithField(s)
}

var _ authentication.AuthenticatorRepo = &AuthenticationRepo{}
