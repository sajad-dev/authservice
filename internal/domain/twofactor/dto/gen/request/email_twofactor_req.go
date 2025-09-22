package request

import (
	"github.com/sajad-dev/authservice/internal/domain/twofactor/twofactorproto"
)

type EmailRequest struct {
	Code int32 `json:"code" validate:""`
} 

func ToRequestEmail(pd *twofactorproto.EmailRequest) *EmailRequest {
	return &EmailRequest{
		Code: pd.Code,
	}
}


