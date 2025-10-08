package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb"
	"github.com/sajad-dev/authservice/authentication/internal/shared/validation/rule"
)

type Validation struct {
	Vld *validator.Validate
	DB  sqldb.SqlDBGlobal
}

// var (
// // Vld  *validator.Validate
// // once *sync.Once
// )

func (v *Validation) _registerRule(vld *validator.Validate) {
	vld.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		return rule.Unique(fl, v.DB)

	})
	vld.RegisterValidation("exists", func(fl validator.FieldLevel) bool {
		return rule.Exists(fl, v.DB)
	})

}

func NewValidator(db sqldb.SqlDBGlobal) Validation {
	vld := Validation{}
	vld.DB = db

	vld.Vld = validator.New()
	vld._registerRule(vld.Vld)

	return vld
}

func (v Validation) Verify(req any) error {
	return v.Vld.Struct(req)
}
