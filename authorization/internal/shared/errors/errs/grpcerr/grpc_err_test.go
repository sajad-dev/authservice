package grpcerr_test

import (
	"testing"

	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs/grpcerr"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs/grpcerr/errorsproto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
)

func TestErr_GRPXErr_ErrorsWithDetails(t *testing.T) {

	errD, err := grpcerr.ErrorsWithDetails(&errorsproto.ErrorDetail{Message: "This very very very very bad error"}, ":((", codes.InvalidArgument)
	assert.NoError(t, err)

	assert.Equal(t, 1, len(errD.Details()))
	errDetails, ok := errD.Details()[0].(*errorsproto.ErrorDetail)
	assert.True(t, ok)
	assert.Equal(t, "This very very very very bad error", errDetails.Message)
}
