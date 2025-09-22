package response

import (
	"github.com/sajad-dev/authservice/internal/domain/forgetpassword/forgetpasswordproto"
)

type ForgetResponse struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
	Token string `json:"token"`
} 

func (a *ForgetResponse) ToProto () *forgetpasswordproto.ForgetResponse {
	return &forgetpasswordproto.ForgetResponse{
		Code: a.Code,
		Msg: a.Msg,
		Token: a.Token,
	}
}
