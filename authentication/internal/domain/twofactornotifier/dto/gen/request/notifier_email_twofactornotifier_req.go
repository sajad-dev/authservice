package request

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/twofactornotifierproto"
)

type NotifierEmailRequest struct {
	Token string `json:"token" validate:"required"`
} 

func ToRequestNotifierEmail(pd *twofactornotifierproto.NotifierEmailRequest) *NotifierEmailRequest {
	return &NotifierEmailRequest{
		Token: pd.Token,
	}
}


