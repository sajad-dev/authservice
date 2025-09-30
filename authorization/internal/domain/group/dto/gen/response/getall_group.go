package response

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/groupproto"
)

type GetAllResponse struct {
	Code int32 `json:"code"`
	Msg string `json:"msg"`
} 

func (a *GetAllResponse) ToProto () *groupproto.GetAllResponse {
	return &groupproto.GetAllResponse{
		Code: a.Code,
		Msg: a.Msg,
	}
}
