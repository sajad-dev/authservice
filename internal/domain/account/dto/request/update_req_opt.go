package request

import "github.com/sajad-dev/authservice/internal/domain/account/accountproto"

type UpdateRequest struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	SMS       string
	Username  string
	TwoFactor []string
	Password  string
}

func ToRequestUpdate(pd accountproto.UpdateRequest) *UpdateRequest {
	return &UpdateRequest{
		ID:        int(pd.ID),
		FirstName: pd.FirstName,
		LastName:  pd.LastName,
		Email:     pd.Email,
		SMS:       pd.SMS,
		Username:  pd.Username,
		TwoFactor: pd.TwoFactor,
		Password:  pd.Password,
	}
}
