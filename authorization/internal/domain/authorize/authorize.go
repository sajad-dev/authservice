package authorize

import "context"

type AuthorizeService interface {
	Check (ctx *context.Context)
}


