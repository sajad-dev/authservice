package request

import "github.com/sajad-dev/authservice/internal/domain/account/accountproto"

type CreateRequest struct {
	FirstName string ‍‍`json:"first_name" validate:""`
	LastName  string
	Email     string
	SMS       string
	Username  string
	TwoFactor []string
	Password  string
}

func ToRequestCreate(pd accountproto.CreateRequest) *CreateRequest {
	return &CreateRequest{
		FirstName: pd.FirstName,
		LastName:  pd.LastName,
		Email:     pd.Email,
		SMS:       pd.SMS,
		Username:  pd.Username,
		TwoFactor: pd.TwoFactor,
		Password:  pd.Password,
	}
}
