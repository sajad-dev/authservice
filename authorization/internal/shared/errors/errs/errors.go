package errs

import "github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs/stacktrace"


func Err(err error) error {
	return stacktrace.ErrStackTrace(err)
}
