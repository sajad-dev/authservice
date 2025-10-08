package postgres

import (
	"fmt"

	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb"
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
	var exists bool
	p.DB.Table(tablename).Where(fmt.Sprintf("%s = ?", fieldname), value).Find(&exists)
	return exists
}

var _ sqldb.SqlDBGlobal = &GlobalPostgres{}
