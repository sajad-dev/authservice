package validation

import (
	"sync"

	"github.com/go-playground/validator"
)

type Validation struct {
	Vld *validator.Validate
}

var (
	vld  *validator.Validate
	once *sync.Once
)

func _registerRule(vld *validator.Validate) {}

func NewValidator() Validation {
	once.Do(func() {
		vld = validator.New()
		_registerRule(vld)
	})

	return Validation{
		Vld: vld,
	}
}

func (v Validation) Verify(req any) error {
	return v.Vld.Struct(req)
}
