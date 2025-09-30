package request

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/groupproto"
)

type DeleteRequest struct {
	Subject string `json:"subject" validate:""`
	Group string `json:"group" validate:""`
} 

func ToRequestDelete(pd *groupproto.DeleteRequest) *DeleteRequest {
	return &DeleteRequest{
		Subject: pd.Subject,
		Group: pd.Group,
	}
}


