package validation

import (
	"github.com/go-playground/validator/v10"

)

type Validation struct {
	Vld *validator.Validate
}

func NewValidator() Validation {
	vld := Validation{}

	vld.Vld = validator.New()

	return vld
}

func (v Validation) Verify(req any) error {
	return v.Vld.Struct(req)
}
