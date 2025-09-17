package response

import "github.com/sajad-dev/authservice/internal/domain/account/accountproto"

type User struct {
	FirstName    string
	LastName     string
	Email        string
	SMS          string
	Username     string
	EmailConfirm bool
	SMSConfirm   bool
}

type ReadResponse struct {
	Code    int
	Message string
	Data    User
}

type ReadResponseOption func(*ReadResponse)

func NewReadResponse(opts ...ReadResponseOption) *ReadResponse {
	at := &ReadResponse{}
	for _, opt := range opts {
		opt(at)
	}
	return at
}

func ReadWithCode(code int) ReadResponseOption {
	return func(r *ReadResponse) {
		r.Code = code
	}
}

func ReadWithMessage(message string) ReadResponseOption {
	return func(r *ReadResponse) {
		r.Message = message
	}
}

func ReadWithFirstName(firstName string) ReadResponseOption {
	return func(r *ReadResponse) {
		r.Data.FirstName = firstName
	}
}

func ReadWithLastName(lastName string) ReadResponseOption {
	return func(r *ReadResponse) {
		r.Data.LastName = lastName
	}
}

func ReadWithEmail(email string) ReadResponseOption {
	return func(r *ReadResponse) {
		r.Data.Email = email
	}
}

func ReadWithSMS(sms string) ReadResponseOption {
	return func(r *ReadResponse) {
		r.Data.SMS = sms
	}
}

func ReadWithUsername(username string) ReadResponseOption {
	return func(r *ReadResponse) {
		r.Data.Username = username
	}
}

func ReadWithEmailConfirm(confirm bool) ReadResponseOption {
	return func(r *ReadResponse) {
		r.Data.EmailConfirm = confirm
	}
}

func ReadWithSMSConfirm(confirm bool) ReadResponseOption {
	return func(r *ReadResponse) {
		r.Data.SMSConfirm = confirm
	}
}

func (r *ReadResponse) ToProto() *accountproto.ReadReply {
	return &accountproto.ReadReply{
		Code:    int32(r.Code),
		Message: r.Message,
		Data: &accountproto.User{
			FirstName:    r.Data.FirstName,
			LastName:     r.Data.LastName,
			Email:        r.Data.Email,
			Sms:          r.Data.SMS,
			Username:     r.Data.Username,
			EmailConfirm: r.Data.EmailConfirm,
			SmsConfirm:   r.Data.SMSConfirm,
		},
	}
}

