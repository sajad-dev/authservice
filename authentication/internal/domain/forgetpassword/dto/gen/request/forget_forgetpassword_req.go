package request

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/forgetpasswordproto"
)

type ForgetRequest struct {
	Email string `json:"email" validate:""`
} 

func ToRequestForget(pd *forgetpasswordproto.ForgetRequest) *ForgetRequest {
	return &ForgetRequest{
		Email: pd.Email,
	}
}


