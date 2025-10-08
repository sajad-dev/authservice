package request

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/twofactorproto"
)

type GoogleRequest struct {
	Token string `json:"token" validate:"required"`
	Code int32 `json:"code" validate:"required"`
} 

func ToRequestGoogle(pd *twofactorproto.GoogleRequest) *GoogleRequest {
	return &GoogleRequest{
		Token: pd.Token,
		Code: pd.Code,
	}
}


