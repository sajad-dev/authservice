package errs

import (
	"errors"

	"gorm.io/gorm"
)

var (
	USERNAME_NOT_VALID        = errors.New("User name not valid")
	INCORRECT_PASSWORD        = errors.New("incorrect password")
	TOKEN_JWT_NOT_VALID       = errors.New("Token not valid")
	SCREAT_KEY_IS_NOT_VALID   = errors.New("Screat key is not valid")
	CODE_OTP_GOOGLE_NOT_VALID = errors.New("Code not valid")
	POLICY_NOT_TRUE           = errors.New("Policy not true")
	NOT_ACCSES                = errors.New("Not acces ")
)

var (
	PROTO_NOT_VALID_MESSAGE_ERROR = errors.New("errParams is not a proto.Message")
)

var (
	RECORD_NOT_FOUND_ORM = gorm.ErrRecordNotFound
)
