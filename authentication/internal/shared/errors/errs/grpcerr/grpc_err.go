package grpcerr

import (
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs/grpcerr/errorsproto"
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs/logging"
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs/stacktrace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorsWithDetails(errParams *errorsproto.ErrorDetail, title string, code codes.Code) (*status.Status, error) {

	errMessage, err := status.New(code, title).WithDetails(errParams)
	if err != nil {
		logging.ErrLog(stacktrace.ErrStackTrace(err))
		return nil, stacktrace.ErrStackTrace(err)
	}
	return errMessage, err
}
