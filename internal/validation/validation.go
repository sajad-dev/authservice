package validation

import "github.com/go-playground/validator/v10"

type ValidationRequest struct {
	ValidationInstance *validator.Validate
}

func NewValidationRequest() *ValidationRequest {
	return &ValidationRequest{
		ValidationInstance: validator.New(),
	}
}

func (v ValidationRequest) ValidationStruct(validation any) error {
	return v.ValidationInstance.Struct(validation)
}
