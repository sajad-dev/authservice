package response

import (
	"github.com/sajad-dev/authservice/internal/shared/modeldto"
	"github.com/sajad-dev/authservice/internal/domain/account/accountproto"
)

type ReadResponse struct {
	Code int32 `json:"code"`
	Data modeldto.AccountFiltered `json:"data"`
	Message string `json:"message"`
} 

func (a *ReadResponse) ToProto () *accountproto.ReadResponse {
	return &accountproto.ReadResponse{
		Code: a.Code,
		Data: AdaptorDataUser(a.Data),
		Message: a.Message,
	}
}
func AdaptorDataUser (data modeldto.AccountFiltered) *accountproto.User{
	return &accountproto.User{
		FirstName: data.LastName,
	}
}
