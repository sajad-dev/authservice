package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model
	FirstName           string
	LastName            string
	Email               string         `gorm:"unique;not null"`
	SMS                 *string        `gorm:"uniqueIndex"`
	Username            string         `gorm:"unique;not null"`
	TwoFactor           pq.StringArray `gorm:"type:text[]"`
	Password            string         `gorm:"not null"`
	EmailConfirm        bool           `gorm:"default:false"`
	SMSConfirm          bool           `gorm:"default:false"`
	GoogleAuthScreatKey string
	LastForgetPassword  time.Time
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

func WithSMS(sms string) AccountOption {
	return func(a *Accounts) { a.SMS = &sms }
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

func WithEmailConfirm(emailConfirm bool) AccountOption {
	return func(a *Accounts) { a.EmailConfirm = emailConfirm }
}

func WithSMSConfirm(smsConfirm bool) AccountOption {
	return func(a *Accounts) { a.SMSConfirm = smsConfirm }
}

func WithLastForgetPassword(t time.Time) AccountOption {
	return func(a *Accounts) { a.LastForgetPassword = t }
}

func WithGoogleAuthScreatKey(screatKey string) AccountOption {
	return func(a *Accounts) { a.GoogleAuthScreatKey = screatKey }
}
func NewAccounts(opts ...AccountOption) *Accounts {
	account := &Accounts{
		EmailConfirm: false,
		SMSConfirm:   false,
	}
	for _, opt := range opts {
		opt(account)
	}
	return account
}
