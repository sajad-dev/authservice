package models

import (
	"reflect"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model
	FirstName           string
	LastName            string
	Email               string         `gorm:"unique;not null"`
	Username            string         `gorm:"unique;not null"`
	TwoFactor           pq.StringArray `gorm:"type:text[]"`
	Password            string         `gorm:"not null"`
	GoogleAuthSecretKey string
	LastForgetPassword  time.Time
}

type AccountFiltered struct {
	Id        int32
	FirstName string
	LastName  string
	Email     string
	Username  string
	TwoFactor pq.StringArray
}

type AccountOption func(*Accounts)

func WithFirstName(firstName string) AccountOption {
	return func(a *Accounts) { a.FirstName = firstName }
}

func WithLastName(lastName string) AccountOption {
	return func(a *Accounts) { a.LastName = lastName }
}

func WithEmail(email string) AccountOption {
	return func(a *Accounts) { a.Email = email }
}

func WithUsername(username string) AccountOption {
	return func(a *Accounts) { a.Username = username }
}

func WithTwoFactor(twoFactor pq.StringArray) AccountOption {
	return func(a *Accounts) { a.TwoFactor = twoFactor }
}

func WithPassword(password string) AccountOption {
	return func(a *Accounts) { a.Password = password }
}

func WithLastForgetPassword(t time.Time) AccountOption {
	return func(a *Accounts) { a.LastForgetPassword = t }
}

func WithGoogleAuthSecretKey(secretKey string) AccountOption {
	return func(a *Accounts) { a.GoogleAuthSecretKey = secretKey }
}
func NewAccounts(opts ...AccountOption) *Accounts {
	account := &Accounts{}
	for _, opt := range opts {
		opt(account)
	}
	return account
}

func getStringField(val reflect.Value, fieldName string) string {
	fieldVal := val.FieldByName(fieldName)
	if fieldVal.IsValid() && fieldVal.Kind() == reflect.String {
		return fieldVal.String()
	}
	return ""
}

func getStringArrayField(val reflect.Value, fieldName string) (pq.StringArray, bool) {
	fieldVal := val.FieldByName(fieldName)
	if fieldVal.IsValid() && fieldVal.Kind() == reflect.Slice {
		if strArray, ok := fieldVal.Interface().(pq.StringArray); ok {
			return strArray, true
		}
	}

	return nil, false
}
func AccountInput(inp any) (*Accounts, bool) {
	val := reflect.ValueOf(inp)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	firstName := getStringField(val, "FirstName")
	lastName := getStringField(val, "LastName")
	email := getStringField(val, "Email")
	password := getStringField(val, "Password")
	username := getStringField(val, "Username")
	twofactor, _ := getStringArrayField(val, "TwoFactor")

	return NewAccounts(
		WithFirstName(firstName),
		WithLastName(lastName),
		WithEmail(email),
		WithPassword(password),
		WithUsername(username),
		WithTwoFactor(twofactor),
	), true
}

func AccountOutput(row *Accounts) AccountFiltered {
	return AccountFiltered{
		FirstName: row.FirstName,
		Id:        int32(row.ID),
		LastName:  row.LastName,
		Email:     row.Email,
		Username:  row.Username,
		TwoFactor: row.TwoFactor,
	}
}
