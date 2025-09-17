package request

import "github.com/sajad-dev/authservice/internal/domain/account/accountproto"

type DeleteRequest struct {
	ID int
}

func ToRequestDelete(pd accountproto.DeleteRequest) *DeleteRequest {
	return &DeleteRequest{
		ID: int(pd.ID),
	}
}
