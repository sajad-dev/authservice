package response

import (
	"github.com/sajad-dev/authservice/internal/domain/account/accountproto"
	"github.com/sajad-dev/authservice/internal/shered/account/modeldto"
)

type ReadResponse struct {
	Code    int
	Message string
	Data    modeldto.AccountFiltered
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

func (r *ReadResponse) ToProto() *accountproto.ReadReply {
	return &accountproto.ReadReply{
		Code: int32(r.Code),
		Msg:  r.Message,
		Data: &accountproto.User{
			FirstName: r.Data.FirstName,
			LastName:  r.Data.LastName,
			Email:     r.Data.Email,
			SMS:       r.Data.SMS,
			Username:  r.Data.Username,
		},
	}
}
