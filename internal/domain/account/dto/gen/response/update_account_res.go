package response

import (
	"github.com/sajad-dev/authservice/internal/domain/account/accountproto"
)

type UpdateResponse struct {
	Code int32 `json:"code"`
	Message string `json:"message"`
} 

func (a *UpdateResponse) ToProto () *accountproto.UpdateResponse {
	return &accountproto.UpdateResponse{
		Code: a.Code,
		Message: a.Message,
	}
}
