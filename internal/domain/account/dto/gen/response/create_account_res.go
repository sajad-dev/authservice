package response

import (
	"github.com/sajad-dev/authservice/internal/domain/account/accountproto"
)

type CreateResponse struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
} 

func (a *CreateResponse) ToProto () *accountproto.CreateResponse {
	return &accountproto.CreateResponse{
		Code: a.Code,
		Msg: a.Msg,
	}
}
