package request

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/policyproto"
)

type CreateRequest struct {
	Subject string `json:"subject" validate:"required,max=64"`
	Object string `json:"object" validate:"required,max=64"`
	Action string `json:"action" validate:"required,max=64"`
} 

func ToRequestCreate(pd *policyproto.CreateRequest) *CreateRequest {
	return &CreateRequest{
		Subject: pd.Subject,
		Object: pd.Object,
		Action: pd.Action,
	}
}


