package request

import (
	"github.com/sajad-dev/authservice/internal/domain/authentication/authenticationproto"
)

type LoginRequest struct {
	Username string `json:"username" validate:""`
	Password string `json:"password" validate:""`
} 

func ToRequestLogin(pd *authenticationproto.LoginRequest) *LoginRequest {
	return &LoginRequest{
		Username: pd.Username,
		Password: pd.Password,
	}
}


