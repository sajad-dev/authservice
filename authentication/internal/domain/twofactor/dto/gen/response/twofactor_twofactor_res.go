package response

import (
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/twofactorproto"
)

type TwoFactorResponse struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
	Token string `json:"token"`
	Data models.AccountFiltered `json:"data"`
} 

func (a *TwoFactorResponse) ToProto () *twofactorproto.TwoFactorResponse {
	return &twofactorproto.TwoFactorResponse{
		Code: a.Code,
		Msg: a.Msg,
		Token: a.Token,
		Data: AdaptorDataUser(a.Data),
	}
}
func AdaptorDataUser (data models.AccountFiltered) *twofactorproto.User{
	return &twofactorproto.User{
		FirstName: data.FirstName,
		LastName: data.LastName,
		Username: data.Username,
		Email: data.Email,
		TwoFactor: data.TwoFactor,
	}
}
