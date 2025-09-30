package response

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/policyproto"
)

type GetAllResponse struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
} 

func (a *GetAllResponse) ToProto () *policyproto.GetAllResponse {
	return &policyproto.GetAllResponse{
		Code: a.Code,
		Msg: a.Msg,
	}
}
