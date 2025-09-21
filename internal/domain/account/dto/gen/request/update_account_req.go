package request

import (
	"github.com/sajad-dev/authservice/internal/domain/account/accountproto"
)

type UpdateRequest struct {
	Id int32 `json:"id" validate:""`
	FirstName string `json:"first_name" validate:""`
	LastName string `json:"last_name" validate:""`
	Email string `json:"email" validate:""`
	Username string `json:"username" validate:""`
	TwoFactor []string `json:"two_factor" validate:""`
	Password string `json:"password" validate:""`
} 

func ToRequestUpdate(pd *accountproto.UpdateRequest) *UpdateRequest {
	return &UpdateRequest{
		Id: pd.Id,
		FirstName: pd.FirstName,
		LastName: pd.LastName,
		Email: pd.Email,
		Username: pd.Username,
		TwoFactor: pd.TwoFactor,
		Password: pd.Password,
	}
}


