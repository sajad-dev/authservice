package repository

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
)

type TwoFactorRepo struct {
	DBAccount sqldb.SqlDB[*models.Accounts]
	DBCode    sqldb.SqlDB[*models.TwoFactorCode]
}

func NewTwoFactorRepo(dbAcc sqldb.SqlDB[*models.Accounts], dbCode sqldb.SqlDB[*models.TwoFactorCode]) *TwoFactorRepo {
	return &TwoFactorRepo{DBAccount: dbAcc, DBCode: dbCode}
}

func (s TwoFactorRepo) FindByCode(code int, codeType string) ([]*models.TwoFactorCode, error) {
	return s.DBCode.Where(&models.TwoFactorCode{
		Code: code,
		Type: codeType,
	})
}

func (s TwoFactorRepo) FindById(id int) (*models.Accounts, error) {
	return s.DBAccount.GetByID(id)
}

var _ twofactor.TwoFactorRepository = &TwoFactorRepo{}
