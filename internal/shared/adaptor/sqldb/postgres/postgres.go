package postgres

import (
	"github.com/sajad-dev/authservice/internal/shared/adaptor/sqldb"
	"gorm.io/gorm"
)

type Postgres[M any] struct {
	DB *gorm.DB
}

func (p *Postgres[M]) Create(params M) error {
	return p.DB.Create(params).Error
}

func (p *Postgres[M]) FindWithField(params M) ([]M, error) {
	var rows []M
	err := p.DB.Where(params).Find(rows).Error
	return rows, err
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

var _ sqldb.SqlDB[struct{}] = &Postgres[struct{}]{}
