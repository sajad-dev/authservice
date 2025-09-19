package request

import "github.com/sajad-dev/authservice/internal/domain/account/accountproto"

type ReadRequest struct {
	ID int32 `json:"id" validate:""`
} 

func ToRequestRead(pd accountproto.ReadRequest) *ReadRequest {
	return &ReadRequest{
		ID: pd.ID,
	}
}

