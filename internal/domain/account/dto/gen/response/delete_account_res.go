package response

import (
	"github.com/sajad-dev/authservice/internal/domain/account/accountproto"
)

type DeleteResponse struct {
	Code int32 `json:"code"`
	Message string `json:"message"`
} 

func (a *DeleteResponse) ToProto () *accountproto.DeleteResponse {
	return &accountproto.DeleteResponse{
		Code: a.Code,
		Message: a.Message,
	}
}
