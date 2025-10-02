package authorize

import (
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/dto/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/dto/response"
)

type AuthorizeService interface {
	Check(req request.AuthorizeRequest) (response.AuthorizeResponse, error)
}

type AuthorizeRepository interface {
	Check(sub string, obj string, act string) (bool, error)
}
