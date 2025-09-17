package validation_test

import (
	"testing"

	"github.com/sajad-dev/hamsokhan/auth/internal/validation"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Name            string `validate:"required,min=3"`
	Email           string `validate:"required,email"`
	Password        string `validate:"required,min=6"`
	ConfirmPassword string `validate:"eqfield=Password"`
}

func TestValidationStruct(t *testing.T) {
	v := validation.NewValidationRequest()

	user := User{
		Name:            "Jo",             
		Email:           "invalid-email",   
		Password:        "123",             
		ConfirmPassword: "1234",           	
	}

	err := v.ValidationStruct(user)
	assert.Error(t,err)

}

