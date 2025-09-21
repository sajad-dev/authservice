package response

import (
	"github.com/sajad-dev/authservice/internal/shared/models"
	"github.com/sajad-dev/authservice/internal/domain/account/accountproto"
)

type ReadResponse struct {
	Code int32 `json:"code"`
	Data models.AccountFiltered `json:"data"`
	Msg string `json:"msg"`
} 

func (a *ReadResponse) ToProto () *accountproto.ReadResponse {
	return &accountproto.ReadResponse{
		Code: a.Code,
		Data: AdaptorDataUser(a.Data),
		Msg: a.Msg,
	}
}
func AdaptorDataUser (data models.AccountFiltered) *accountproto.User{
	return &accountproto.User{
		FirstName: data.FirstName,
		LastName: data.LastName,
		Username: data.Username,
		Email: data.Email,
		TwoFactor: data.TwoFactor,
	}
}
