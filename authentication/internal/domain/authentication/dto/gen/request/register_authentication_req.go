package request

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/authenticationproto"
)

type RegisterRequest struct {
	Username string `json:"username" validate:"required,max=256,unique=accounts"`
	Email string `json:"email" validate:"required,max=256,unique=accounts"`
	FirstName string `json:"first_name" validate:"required,max=64"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	Password string `json:"password" validate:"required,min=8"`
} 

func ToRequestRegister(pd *authenticationproto.RegisterRequest) *RegisterRequest {
	return &RegisterRequest{
		Username: pd.Username,
		Email: pd.Email,
		FirstName: pd.FirstName,
		PasswordConfirmation: pd.PasswordConfirmation,
		Password: pd.Password,
	}
}


