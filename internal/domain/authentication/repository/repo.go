package repository

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

func (a AuthenticationRepo) Create(row *models.Accounts) error {
	return a.DB.Create(row)
}

func (s AuthenticationRepo) Find(clm string, value string) (*models.Accounts, error) {
	return s.DB.WhereField(clm, value)
}

var _ authentication.AuthenticatorRepository = &AuthenticationRepo{}
