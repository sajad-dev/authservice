package request

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/policyproto"
)

type DeleteRequest struct {
	Subject string `json:"subject" validate:""`
	Group string `json:"group" validate:""`
	Action string `json:"action" validate:""`
} 

func ToRequestDelete(pd *policyproto.DeleteRequest) *DeleteRequest {
	return &DeleteRequest{
		Subject: pd.Subject,
		Group: pd.Group,
		Action: pd.Action,
	}
}


