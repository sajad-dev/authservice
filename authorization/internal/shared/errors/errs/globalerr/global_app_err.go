package globalerr

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs/grpcerr"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs/grpcerr/errorsproto"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs/logging"
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

		errMessage, err := grpcerr.ErrorsWithDetails(&errorsproto.ErrorDetail{Message: messages.ERR_INTERNAL_SERVER}, messages.ERR_INTERNAL_SERVER, codes.Internal)

		logging.ErrLog(errParametr)

		if err != nil {
			return WithDetailsErr(err, errors.New(messages.ERR_INTERNAL_SERVER))
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
				messages.ERR_VALIDATION,
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
