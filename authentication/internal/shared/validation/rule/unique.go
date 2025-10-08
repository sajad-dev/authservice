package rule

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb"
)

func Unique[T any](fl validator.FieldLevel, db sqldb.SqlDB[T]) bool {
	fieldName := fl.FieldName()
	params := strings.Split(fl.Param(), ";")

	var exists bool
	db.Table(params[0]).Where(fmt.Sprintf("%s = ?", fieldName), params[1]).Find(&exists)

	return exists
}
