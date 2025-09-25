package request

import (
	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier/twofactornotifierproto"
)

type NotifierEmailRequest struct {
	Token string `json:"token" validate:""`
} 

func ToRequestNotifierEmail(pd *twofactornotifierproto.NotifierEmailRequest) *NotifierEmailRequest {
	return &NotifierEmailRequest{
		Token: pd.Token,
	}
}


