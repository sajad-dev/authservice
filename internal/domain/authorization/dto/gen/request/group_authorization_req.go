package request

import (
	"github.com/sajad-dev/authservice/internal/domain/authorization/authorizationproto"
)

type GroupRequest struct {
	sub string `json:"sub" validate:""`
	Grp string `json:"grp" validate:""`
} 

func ToRequestGroup(pd *authorizationproto.GroupRequest) *GroupRequest {
	return &GroupRequest{
		Sub: pd.sub,
		Grp: pd.Grp,
	}
}


