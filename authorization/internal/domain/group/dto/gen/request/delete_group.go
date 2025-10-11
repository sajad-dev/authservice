package request

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/groupproto"
)

type DeleteRequest struct {
	Subject string `json:"subject" validate:"required,max=64"`
	Group string `json:"group" validate:"required,max=64"`
} 

func ToRequestDelete(pd *groupproto.DeleteRequest) *DeleteRequest {
	return &DeleteRequest{
		Subject: pd.Subject,
		Group: pd.Group,
	}
}


