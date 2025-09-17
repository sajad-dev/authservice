package request

import "github.com/sajad-dev/authservice/internal/domain/account/accountproto"

type ReadRequest struct {
	ID int
}

func ToRequestRead(pd accountproto.ReadRequest) *ReadRequest {
	return &ReadRequest{
		ID: int(pd.ID),
	}
}

