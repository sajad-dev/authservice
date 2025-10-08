package request

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/twofactorproto"
)

type EmailRequest struct {
	Code int32 `json:"code" validate:"required"`
} 

func ToRequestEmail(pd *twofactorproto.EmailRequest) *EmailRequest {
	return &EmailRequest{
		Code: pd.Code,
	}
}


