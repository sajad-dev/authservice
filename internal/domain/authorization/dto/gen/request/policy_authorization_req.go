package request

import (
	"github.com/sajad-dev/authservice/internal/domain/authorization/authorizationproto"
)

type PolicyRequest struct {
	sub string `json:"sub" validate:""`
	obj string `json:"obj" validate:""`
	act string `json:"act" validate:""`
} 

func ToRequestPolicy(pd *authorizationproto.PolicyRequest) *PolicyRequest {
	return &PolicyRequest{
		Sub: pd.sub,
		Obj: pd.obj,
		Act: pd.act,
	}
}


