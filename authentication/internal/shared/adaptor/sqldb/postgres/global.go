package postgres

import (
	"fmt"

	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
	"gorm.io/gorm"
)

type GlobalPostgres struct {
	DB *gorm.DB
}

func NewGlobalPostgres(db *gorm.DB) *GlobalPostgres {
	return &GlobalPostgres{
		DB: db,
	}
}

func (p *GlobalPostgres) Exists(tablename string, fieldname, value string) bool {
	var e []models.Accounts
	p.DB.Table(tablename).Where(fmt.Sprintf("%s = ?", fieldname), value).Find(&e)
	return (len(e) == 1)
}

var _ sqldb.SqlDBGlobal = &GlobalPostgres{}
