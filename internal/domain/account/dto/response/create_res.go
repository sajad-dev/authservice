package response

import "github.com/sajad-dev/authservice/internal/domain/account/accountproto"

type CreateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CreateResponseOption func(*CreateResponse)

func NewCreateResponse(opts ...CreateResponseOption) *CreateResponse {
	at := &CreateResponse{}
	for _, opt := range opts {
		opt(at)
	}
	return at
}

func CreateWithCode(code int) CreateResponseOption {
	return func(lr *CreateResponse) {
		lr.Code = code
	}
}

func CreateWithMessage(msg string) CreateResponseOption {
	return func(lr *CreateResponse) {
		lr.Message = msg
	}
}

func (c *CreateResponse) ToProto () *accountproto.CreateReply {
	return &accountproto.CreateReply{
		Code: int32(c.Code),
		Message: c.Message,
	}
}
