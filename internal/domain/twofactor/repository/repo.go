package repositories

import (
	"github.com/sajad-dev/authservice/internal/domain/forgetpassword"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/sqldb"
	"github.com/sajad-dev/authservice/internal/shared/models"
)

type TwoFactorRepo struct {
	DBAccount sqldb.SqlDB[*models.Accounts]
	DBCode    sqldb.SqlDB[*models.TwoFactorCode]
}

func NewTwoFactorRepo(db sqldb.SqlDB[*models.Accounts]) *TwoFactorRepo {
	return &TwoFactorRepo{DB: db}
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

var _ forgetpassword.TwoFactorRepo = &TwoFactorRepo{}
