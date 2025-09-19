package response

import (
	"github.com/sajad-dev/authservice/internal/domain/account/accountproto"
)

type CreateResponse struct {
	Code int32 `json:"code"`
	Message string `json:"message"`
} 

func (a *CreateResponse) ToProto () *accountproto.CreateResponse {
	return &accountproto.CreateResponse{
		Code: a.Code,
		Message: a.Message,
	}
}
