package request

import (
	"github.com/sajad-dev/authservice/internal/domain/account/accountproto"
)

type CreateRequest struct {
	FirstName string `json:"first_name" validate:""`
	LastName string `json:"last_name" validate:""`
	Email string `json:"email" validate:""`
	Username string `json:"username" validate:""`
	TwoFactor []string `json:"two_factor" validate:""`
	Password string `json:"password" validate:""`
} 

func ToRequestCreate(pd *accountproto.CreateRequest) *CreateRequest {
	return &CreateRequest{
		FirstName: pd.FirstName,
		LastName: pd.LastName,
		Email: pd.Email,
		Username: pd.Username,
		TwoFactor: pd.TwoFactor,
		Password: pd.Password,
	}
}


