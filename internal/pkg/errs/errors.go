package errs

import "github.com/sajad-dev/authservice/internal/pkg/errs/stacktrace"


func Err(err error) error {
	return stacktrace.ErrStackTrace(err)
}
