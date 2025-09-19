package request

import "github.com/sajad-dev/authservice/internal/domain/account/accountproto"

type DeleteRequest struct {
	ID int32 `json:"id" validate:""`
} 

func ToRequestDelete(pd accountproto.DeleteRequest) *DeleteRequest {
	return &DeleteRequest{
		ID: pd.ID,
	}
}

