package repository

import (
	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/sqldb"
	"github.com/sajad-dev/authservice/internal/shared/models"
)

type TwoFactorNotifierRepo struct {
	DBAccount sqldb.SqlDB[*models.Accounts]
	DBCode    sqldb.SqlDB[*models.TwoFactorCode]
}

func NewTwoFactorNotifierRepo(dbAccount sqldb.SqlDB[*models.Accounts], dbCode sqldb.SqlDB[*models.TwoFactorCode]) *TwoFactorNotifierRepo {
	return &TwoFactorNotifierRepo{DBAccount: dbAccount, DBCode: dbCode}
}

func (s *TwoFactorNotifierRepo) RemoveExpierd(id int) error {
	return s.DBCode.RemoveExpierd("account_id", id)
}
func (s *TwoFactorNotifierRepo) FindById(id int) (*models.Accounts, error) {

	return s.DBAccount.GetByID(id)

}
func (s *TwoFactorNotifierRepo) CreateCode(row *models.TwoFactorCode) error {
	return s.DBCode.Create(row)
}

var _ twofactornotifier.TwoFactorNotifierRepository = &TwoFactorNotifierRepo{}
