package globalerr

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/sajad-dev/authservice/internal/constants/messages"
	"github.com/sajad-dev/authservice/internal/pkg/errs/grpc/errorsproto"
	"github.com/sajad-dev/authservice/internal/pkg/errs/grpc/grpcerr"
	"github.com/sajad-dev/authservice/internal/pkg/errs/logging"
	"google.golang.org/grpc/codes"
)

func WithDetailsErr(err error, errParametr error) error {
	if errParametr != nil {
		logging.ErrLog(errors.New("Error Detail :(( " + errParametr.Error()))
	}
	return err
}

func ServerErr(errParametr error) error {
	if errParametr != nil {
		errMessage, err := grpcerr.ErrorsWithDetails(&errorsproto.ErrorDetail{Message: messages.SERVER_ERR}, messages.SERVER_ERR, codes.Internal)

		logging.ErrLog(errParametr)

		if err != nil {
			return WithDetailsErr(err, errors.New(messages.SERVER_ERR))
		}

		return errMessage.Err()
	} else {
		return nil
	}
}
func ValidationErr(errParametr error) error {
	if errParametr != nil {
		if validationErr, ok := errParametr.(validator.ValidationErrors); ok {
			errMessage, err := grpcerr.ErrorsWithDetails(
				&errorsproto.ErrorDetail{Message: validationErr.Error()},
				messages.VALIDATION_ERR_TITLE,
				codes.InvalidArgument,
			)
			if err != nil {
				return WithDetailsErr(err, errParametr)
			}
			return errMessage.Err()
		}

		return ServerErr(errParametr)
	}
	return nil
}
