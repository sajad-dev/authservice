package request

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/groupproto"
)

type CreateRequest struct {
	Subject string `json:"subject" validate:"required,max=64"`
	Group string `json:"group" validate:"required,max=64"`
} 

func ToRequestCreate(pd *groupproto.CreateRequest) *CreateRequest {
	return &CreateRequest{
		Subject: pd.Subject,
		Group: pd.Group,
	}
}


