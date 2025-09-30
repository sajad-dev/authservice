package repository

import (
	"github.com/sajad-dev/authservice/internal/domain/account"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/sqldb"
	"github.com/sajad-dev/authservice/internal/shared/models"
)

type AccountRepo struct {
	DB sqldb.SqlDB[*models.Accounts]
}

func NewAccountRepo(db sqldb.SqlDB[*models.Accounts]) *AccountRepo {
	return &AccountRepo{DB: db}
}

func (a AccountRepo) Create(req *models.Accounts) error {
	return a.DB.Create(req)
}

func (a AccountRepo) Update(req *models.Accounts, id int) error {
	req.ID = uint(id)
	return a.DB.Save(req)
}

func (r *AccountRepo) Read(id int) (*models.Accounts, error) {
	return r.DB.GetByID(id)

}

func (r *AccountRepo) Delete(id int) error {
	return r.DB.Delete(id)
}

var _ account.AccountCURDRepository = &AccountRepo{}
