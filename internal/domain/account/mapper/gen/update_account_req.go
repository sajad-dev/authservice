package request

import "github.com/sajad-dev/authservice/internal/domain/account/accountproto"

type UpdateRequest struct {
	ID        int32    `json:"id" validate:""`
	FirstName string   `json:"first_name" validate:""`
	LastName  string   `json:"last_name" validate:""`
	Email     string   `json:"email" validate:""`
	SMS       string   `json:"sms" validate:""`
	Username  string   `json:"username" validate:""`
	TwoFactor []string `json:"two_factor" validate:""`
	Password  string   `json:"password" validate:""`
}

func ToRequestUpdate(pd accountproto.UpdateRequest) *UpdateRequest {
	return &UpdateRequest{
		ID:        pd.ID,
		FirstName: pd.FirstName,
		LastName:  pd.LastName,
		Email:     pd.Email,
		SMS:       pd.SMS,
		Username:  pd.Username,
		TwoFactor: pd.TwoFactor,
		Password:  pd.Password,
	}
}
