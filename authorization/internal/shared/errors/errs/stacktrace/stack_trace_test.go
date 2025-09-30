package stacktrace_test

import (
	"errors"
	"testing"

	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs/stacktrace"
	"github.com/stretchr/testify/assert"
)

func TestErr(t *testing.T) {
	err := errors.New("This test error haha haha haha :)")
	stacketrace := stacktrace.ErrStackTrace(err)
	assert.Contains(t, stacketrace.Error(), "This test error haha haha haha :)")
}
