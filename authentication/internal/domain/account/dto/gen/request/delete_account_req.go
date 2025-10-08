package request

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/account/accountproto"
)

type DeleteRequest struct {
	Id int32 `json:"id" validate:"required,exists=accounts"`
} 

func ToRequestDelete(pd *accountproto.DeleteRequest) *DeleteRequest {
	return &DeleteRequest{
		Id: pd.Id,
	}
}


