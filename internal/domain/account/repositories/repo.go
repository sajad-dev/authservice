package repositories

import (
	"github.com/sajad-dev/authservice/internal/adaptor/sqldb"
	"github.com/sajad-dev/authservice/internal/domain/account"
	"github.com/sajad-dev/authservice/internal/domain/account/models"
)

type AccountRepo struct {
	DB sqldb.SqlDB[*models.Accounts]
}

func NewAccountRepo(db sqldb.SqlDB[*models.Accounts]) *AccountRepo {
	return &AccountRepo{DB: db}
}

var _ account.AccountCURDRepositories = &AccountRepo{}

