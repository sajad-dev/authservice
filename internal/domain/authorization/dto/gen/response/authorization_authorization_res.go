package response

import (
	"github.com/sajad-dev/authservice/internal/domain/authorization/authorizationproto"
)

type AuthorizationResponse struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
} 

func (a *AuthorizationResponse) ToProto () *authorizationproto.AuthorizationResponse {
	return &authorizationproto.AuthorizationResponse{
		Code: a.Code,
		Msg: a.Msg,
	}
}
