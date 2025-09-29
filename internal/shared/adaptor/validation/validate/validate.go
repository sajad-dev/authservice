package validate

import "github.com/go-playground/validator"

type Validate struct {
	Validation *validator.Validate
}

func NewValidate(vld *validator.Validate) *Validate {
	return &Validate{
		Validation: vld,
	}
}

func (v Validate) Validate(validation any) error {
	return v.Validation.Struct(validation)
}
