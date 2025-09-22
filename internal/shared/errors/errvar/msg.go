package errvar

import "errors"

var (
	USER_EXIST = errors.New("User exist")
	USERNAME_NOT_VALID = errors.New("error not valid")
)
