package response

import (
	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier/twofactornotifierproto"
)

type TwoFactorNotifierResponse struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
} 

func (a *TwoFactorNotifierResponse) ToProto () *twofactornotifierproto.TwoFactorNotifierResponse {
	return &twofactornotifierproto.TwoFactorNotifierResponse{
		Code: a.Code,
		Msg: a.Msg,
	}
}
