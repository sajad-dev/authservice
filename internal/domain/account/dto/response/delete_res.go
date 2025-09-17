package response

import "github.com/sajad-dev/authservice/internal/domain/account/accountproto"

type DeleteResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DeleteResponseOption func(*DeleteResponse)

func NewDeleteResponse(opts ...DeleteResponseOption) *DeleteResponse {
	at := &DeleteResponse{}
	for _, opt := range opts {
		opt(at)
	}
	return at
}

func DeleteWithCode(code int) DeleteResponseOption {
	return func(lr *DeleteResponse) {
		lr.Code = code
	}
}

func DeleteWithMessage(msg string) DeleteResponseOption {
	return func(lr *DeleteResponse) {
		lr.Message = msg
	}
}

func (d *DeleteResponse) ToProto() *accountproto.DeleteReply {
	return &accountproto.DeleteReply{
		Code:    int32(d.Code),
		Message: d.Message,
	}
}

