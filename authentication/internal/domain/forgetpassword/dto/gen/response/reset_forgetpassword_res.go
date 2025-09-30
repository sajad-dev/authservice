package response

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/forgetpasswordproto"
)

type ResetResponse struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
} 

func (a *ResetResponse) ToProto () *forgetpasswordproto.ResetResponse {
	return &forgetpasswordproto.ResetResponse{
		Code: a.Code,
		Msg: a.Msg,
	}
}
