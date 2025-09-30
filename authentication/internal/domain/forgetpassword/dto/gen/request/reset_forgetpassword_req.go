package request

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/forgetpasswordproto"
)

type ResetRequest struct {
	Token string `json:"token" validate:""`
	PasswordConfirmation string `json:"password_confirmation" validate:""`
	Password string `json:"password" validate:""`
} 

func ToRequestReset(pd *forgetpasswordproto.ResetRequest) *ResetRequest {
	return &ResetRequest{
		Token: pd.Token,
		PasswordConfirmation: pd.PasswordConfirmation,
		Password: pd.Password,
	}
}


