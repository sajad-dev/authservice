package request

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/authenticationproto"
)

type RegisterRequest struct {
	Username string `json:"username" validate:""`
	Email string `json:"email" validate:""`
	FirstName string `json:"first_name" validate:""`
	PasswordConfirmation string `json:"password_confirmation" validate:""`
	Password string `json:"password" validate:""`
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


