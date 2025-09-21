package response

import (
	"github.com/sajad-dev/authservice/internal/domain/account/accountproto"
)

type UpdateResponse struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
} 

func (a *UpdateResponse) ToProto () *accountproto.UpdateResponse {
	return &accountproto.UpdateResponse{
		Code: a.Code,
		Msg: a.Msg,
	}
}
