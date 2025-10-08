package rule

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb"
)

func Unique[T any](fl validator.FieldLevel, db sqldb.SqlDB[T]) bool {
	fieldName := fl.FieldName()
	table := fl.Param()

	var exists bool
	db.Table(table).Where(fmt.Sprintf("%s = ?", fieldName), fl.Field().String()).Find(&exists)

	return !exists
}


func Exists[T any](fl validator.FieldLevel, db sqldb.SqlDB[T]) bool {
	fieldName := fl.FieldName()
	table := fl.Param()

	var exists bool
	db.Table(table).Where(fmt.Sprintf("%s = ?", fieldName), fl.Field().String()).Find(&exists)

	return exists
}


