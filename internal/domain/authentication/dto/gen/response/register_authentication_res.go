package response

import (
	"github.com/sajad-dev/authservice/internal/shared/models"
	"github.com/sajad-dev/authservice/internal/domain/authentication/authenticationproto"
)

type RegisterResponse struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
	Data models.AccountFiltered `json:"data"`
	Token string `json:"token"`
} 

func (a *RegisterResponse) ToProto () *authenticationproto.RegisterResponse {
	return &authenticationproto.RegisterResponse{
		Code: a.Code,
		Msg: a.Msg,
		Data: AdaptorDataUserRegister(a.Data),
		Token: a.Token,
	}
}
func AdaptorDataUserRegister (data models.AccountFiltered) *authenticationproto.User{
	return &authenticationproto.User{
		FirstName: data.FirstName,
		LastName: data.LastName,
		Username: data.Username,
		Email: data.Email,
		TwoFactor: data.TwoFactor,
	}
}
