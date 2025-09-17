package account

import (
	"context"

	"github.com/sajad-dev/authservice/internal/domain/account/accountproto"
)

type AccountCURDHandler interface {
	CreateGRPC(ctx context.Context, req *accountproto.CreateRequest) (*accountproto.CreateReply, error)
	UpdateGRPC(ctx context.Context, req *accountproto.UpdateRequest) (*accountproto.UpdateReply, error)
	DeleteGRPC(ctx context.Context, req *accountproto.DeleteRequest) (*accountproto.DeleteReply, error)
	ReadGRPC(ctx context.Context, req *accountproto.ReadRequest) (*accountproto.ReadReply, error)
	GetAllGRPC(ctx context.Context, req *accountproto.GetAllRequest) (*accountproto.GetAllReply, error)
}
