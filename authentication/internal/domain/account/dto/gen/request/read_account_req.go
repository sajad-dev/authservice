package request

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/account/accountproto"
)

type ReadRequest struct {
	Id int32 `json:"id" validate:"required,exists=accounts"`
} 

func ToRequestRead(pd *accountproto.ReadRequest) *ReadRequest {
	return &ReadRequest{
		Id: pd.Id,
	}
}


