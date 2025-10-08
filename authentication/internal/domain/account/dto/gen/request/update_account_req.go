package request

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/account/accountproto"
)

type UpdateRequest struct {
	Id int32 `json:"id" validate:"required,exists=accounts"`
	FirstName string `json:"first_name" validate:"required,max=64"`
	LastName string `json:"last_name" validate:"required,max=64"`
	Email string `json:"email" validate:"required,max=256,unique=accounts"`
	Username string `json:"username" validate:"required,max=256,unique=accounts"`
	TwoFactor []string `json:"two_factor" validate:"required,max=3,dive"`
	Password string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
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
		PasswordConfirmation: pd.PasswordConfirmation,
	}
}


