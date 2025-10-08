package validation

import (
	"sync"

	"github.com/go-playground/validator"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb"
	"github.com/sajad-dev/authservice/authentication/internal/shared/validation/rule"
)

type Validation[T any] struct {
	Vld *validator.Validate
	DB  sqldb.SqlDB[T]
}

var (
	vld  *validator.Validate
	once *sync.Once
)

func (v *Validation[T]) _registerRule(vld *validator.Validate) {
	vld.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		return rule.Unique[T](fl, v.DB)
	})
}

func NewValidator[T any](db sqldb.SqlDB[T]) Validation[T] {
	vlda := Validation[T]{}
	vlda.DB = db

	once.Do(func() {
		vld = validator.New()
		vlda._registerRule(vld)
		vlda.Vld = vld
	})

	return vlda
}

func (v Validation[T]) Verify(req any) error {
	return v.Vld.Struct(req)
}
