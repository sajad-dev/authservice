package repositories

import (
	"github.com/sajad-dev/authservice/internal/domain/forgetpassword"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/sqldb"
	"github.com/sajad-dev/authservice/internal/shared/models"
)

type ForgetPasswordRepo struct {
	DB sqldb.SqlDB[*models.Accounts]
}

func NewForgetPasswordRepo(db sqldb.SqlDB[*models.Accounts]) *ForgetPasswordRepo {
	return &ForgetPasswordRepo{DB: db}
}

func (a ForgetPasswordRepo) Update(row *models.Accounts) error {
	return a.DB.Save(row)
}

func (s ForgetPasswordRepo) Find(clm string, value string) (*models.Accounts, error) {
	return s.DB.WhereField(clm, value)
}

func (s ForgetPasswordRepo) FindById(id int) (*models.Accounts, error) {
	return s.DB.GetByID(id)
}

var _ forgetpassword.ForgetPasswordRepo = &ForgetPasswordRepo{}
