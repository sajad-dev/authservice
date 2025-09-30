package request

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/policyproto"
)

type GroupRequest struct {
	Subject string `json:"subject" validate:""`
	Group string `json:"group" validate:""`
} 

func ToRequestGroup(pd *policyproto.GroupRequest) *GroupRequest {
	return &GroupRequest{
		Subject: pd.Subject,
		Group: pd.Group,
	}
}


