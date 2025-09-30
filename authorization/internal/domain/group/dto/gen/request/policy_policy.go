package request

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/policyproto"
)

type PolicyRequest struct {
	Subject string `json:"subject" validate:""`
	Group string `json:"group" validate:""`
	Action string `json:"action" validate:""`
} 

func ToRequestPolicy(pd *policyproto.PolicyRequest) *PolicyRequest {
	return &PolicyRequest{
		Subject: pd.Subject,
		Group: pd.Group,
		Action: pd.Action,
	}
}


