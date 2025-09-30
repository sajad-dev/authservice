package sqlite

import (
	"fmt"
	"time"

	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb"
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs"
	"gorm.io/gorm"
)

type Sqlite[M any] struct {
	DB *gorm.DB
}

func NewSqlite[M any](db *gorm.DB) *Sqlite[M] {
	return &Sqlite[M]{
		DB: db,
	}
}

func (p *Sqlite[M]) Create(params M) error {
	return p.DB.Create(params).Error
}

func (p *Sqlite[M]) Where(params M) ([]M, error) {
	var rows []M
	err := p.DB.Where(params).Find(&rows).Error
	return rows, err
}

func (p *Sqlite[M]) WhereField(column string, value string) (M, error) {
	var table M
	err := p.DB.Where(fmt.Sprintf("%s = ?", column), value).First(&table).Error
	return table, errs.Err(err)
}

func (p *Sqlite[M]) Save(params M) error {
	return p.DB.Save(params).Error
}

func (p *Sqlite[M]) GetByID(id int) (M, error) {
	var params M
	err := p.DB.First(&params,id).Error
	return params, err
}

func (p *Sqlite[M]) Delete(id int) error {
	var params M
	return p.DB.Delete(&params, id).Error
}

func (p *Sqlite[M]) RemoveExpierd(column string, id int) error {
	var params M
	return p.DB.Where(fmt.Sprintf("%s = ?", column), id).
		Or("expired_at < ?", time.Now()).
		Delete(&params).Error
}

var _ sqldb.SqlDB[struct{}] = &Sqlite[struct{}]{}
