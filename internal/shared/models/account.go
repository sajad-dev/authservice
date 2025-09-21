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
	GoogleAuthScreatKey string
	LastForgetPassword  time.Time
}

type AccountFiltered struct {
	FirstName string
	LastName  string
	Email     string
	Username  string
	TwoFactor []string
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

func WithGoogleAuthScreatKey(screatKey string) AccountOption {
	return func(a *Accounts) { a.GoogleAuthScreatKey = screatKey }
}
func NewAccounts(opts ...AccountOption) *Accounts {
	account := &Accounts{}
	for _, opt := range opts {
		opt(account)
	}
	return account
}

func AccountInput(inp any) (*Accounts, bool) {
	val := reflect.ValueOf(inp)

	twofactor, ok := val.FieldByName("TwoFactor").Interface().(pq.StringArray)
	if !ok {
		return nil, false
	}

	return &Accounts{
		FirstName: val.FieldByName("FirstName").String(),
		LastName:  val.FieldByName("LastName").String(),
		Email:     val.FieldByName("Email").String(),
		Password:  val.FieldByName("Password").String(),
		Username:  val.FieldByName("Username").String(),
		TwoFactor: twofactor,
	}, true
}

func AccountOutput(row *Accounts) AccountFiltered {
	return AccountFiltered{
		FirstName: row.FirstName,
		LastName:  row.LastName,
		Email:     row.Email,
		Username:  row.Username,
		TwoFactor: row.TwoFactor,
	}
}
