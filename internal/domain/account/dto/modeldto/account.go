package modeldto

import (
	"reflect"

	"github.com/lib/pq"
	"github.com/sajad-dev/authservice/internal/domain/account/models"
)

type AccountFiltered struct {
	FirstName string
	LastName  string
	Email     string
	SMS       string
	Username  string
	TwoFactor []string
}

func AccountInput(inp any) (*models.Accounts, bool) {
	val := reflect.ValueOf(inp)

	twofactor, ok := val.FieldByName("TwoFactor").Interface().(pq.StringArray)
	if !ok {
		return nil, false
	}

	return &models.Accounts{
		FirstName: val.FieldByName("FirstName").String(),
		LastName:  val.FieldByName("LastName").String(),
		Email:     val.FieldByName("Email").String(),
		Password:  val.FieldByName("Password").String(),
		SMS:       func(s string) *string { return &s }(val.FieldByName("SMS").String()),
		Username:  val.FieldByName("Username").String(),
		TwoFactor: twofactor,
	}, true
}

func AccountOutput(row *models.Accounts) AccountFiltered {
	return AccountFiltered{
		FirstName: row.FirstName,
		LastName:  row.LastName,
		Email:     row.Email,
		SMS:       *row.SMS,
		Username:  row.Username,
		TwoFactor: row.TwoFactor,
	}
}
