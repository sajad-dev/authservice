package response

import "github.com/sajad-dev/authservice/internal/domain/account/accountproto"

type UpdateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UpdateResponseOption func(*UpdateResponse)

func NewUpdateResponse(opts ...UpdateResponseOption) *UpdateResponse {
	at := &UpdateResponse{}
	for _, opt := range opts {
		opt(at)
	}
	return at
}

func UpdateWithCode(code int) UpdateResponseOption {
	return func(lr *UpdateResponse) {
		lr.Code = code
	}
}

func UpdateWithMessage(msg string) UpdateResponseOption {
	return func(lr *UpdateResponse) {
		lr.Message = msg
	}
}

func (u *UpdateResponse) ToProto() *accountproto.UpdateReply {
	return &accountproto.UpdateReply{
		Code:    int32(u.Code),
		Message: u.Message,
	}
}

