package request

import "github.com/sajad-dev/authservice/internal/domain/account/accountproto"

func ToRequestCreate(pd accountproto.CreateRequest) *CreateRequest {
	return &CreateRequest{
		FirstName: pd.FirstName,
		LastName: pd.LastName,
		Email: pd.Email,
		SMS: pd.SMS,
		Username: pd.Username,
		TwoFactor: pd.TwoFactor,
		Password: pd.Password,
	}
}

