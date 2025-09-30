package request

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/policyproto"
)

type CreateRequest struct {
	Subject string `json:"subject" validate:""`
	Object string `json:"object" validate:""`
	Action string `json:"action" validate:""`
} 

func ToRequestCreate(pd *policyproto.CreateRequest) *CreateRequest {
	return &CreateRequest{
		Subject: pd.Subject,
		Object: pd.Object,
		Action: pd.Action,
	}
}


