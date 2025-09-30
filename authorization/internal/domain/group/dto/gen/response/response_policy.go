package response

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/policyproto"
)

type Response struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
} 

func (a *Response) ToProto () *policyproto.Response {
	return &policyproto.Response{
		Code: a.Code,
		Msg: a.Msg,
	}
}
