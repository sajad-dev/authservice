package postgres

import (
	"fmt"
	"time"

	"github.com/sajad-dev/authservice/internal/shared/adaptor/sqldb"
	"github.com/sajad-dev/authservice/internal/shared/errors/errs"
	"gorm.io/gorm"
)

type Postgres[M any] struct {
	DB *gorm.DB
}

func (p *Postgres[M]) Create(params M) error {
	return p.DB.Create(params).Error
}

func (p *Postgres[M]) Where(params M) ([]M, error) {
	var rows []M
	err := p.DB.Where(params).Find(rows).Error
	return rows, err
}

func (p *Postgres[M]) WhereField(column string, value string) (M, error) {
	var table M
	err := p.DB.Where(fmt.Sprintf("%s = ?", column), value).First(&table).Error
	return table, errs.Err(err)
}

func (p *Postgres[M]) Save(params M) error {
	return p.DB.Save(params).Error
}

func (p *Postgres[M]) GetByID(id int) (M, error) {
	var params M
	err := p.DB.First(params).Error
	return params, err
}

func (p *Postgres[M]) Delete(id int) error {
	var params M
	return p.DB.Delete(params, id).Error
}

func (p *Postgres[M]) RemoveExpierd(column string, id int) error {
	var params M
	return p.DB.Where(fmt.Sprintf("%s = ?", column), id).
		Or("expired_at < ?", time.Now()).
		Delete(params).Error
}

var _ sqldb.SqlDB[struct{}] = &Postgres[struct{}]{}
