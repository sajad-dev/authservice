package request

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/forgetpasswordproto"
)

type ResetRequest struct {
	Token string `json:"token" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	Password string `json:"password" validate:"required,min=8"`
} 

func ToRequestReset(pd *forgetpasswordproto.ResetRequest) *ResetRequest {
	return &ResetRequest{
		Token: pd.Token,
		PasswordConfirmation: pd.PasswordConfirmation,
		Password: pd.Password,
	}
}


