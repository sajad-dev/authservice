package request

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/authenticationproto"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required,exists=accounts"`
	Password string `json:"password" validate:"required"`
} 

func ToRequestLogin(pd *authenticationproto.LoginRequest) *LoginRequest {
	return &LoginRequest{
		Username: pd.Username,
		Password: pd.Password,
	}
}


