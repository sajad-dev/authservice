package rule

import (
	"github.com/go-playground/validator/v10"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb"
)

func Unique(fl validator.FieldLevel, db sqldb.SqlDBGlobal) bool {
	fieldName := fl.FieldName()
	table := fl.Param()

	return !db.Exists(table, fieldName, fl.Field().String())

}

func Exists(fl validator.FieldLevel, db sqldb.SqlDBGlobal) bool {
	fieldName := fl.FieldName()
	table := fl.Param()

	return db.Exists(table, fieldName, fl.Field().String())
}
