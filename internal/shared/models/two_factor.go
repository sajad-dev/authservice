package models

import (
	"time"

	"gorm.io/gorm"
)

type TwoFactorCode struct {
	gorm.Model
	Code      int       `gorm:"unique;not null"`
	Type      string    `gorm:"not null"`
	ExpiredAt time.Time `gorm:"column:expired_at"`
	AccountID uint
	Account   Accounts `gorm:"foreignKey:AccountID"`
}

type TwoFactorCodeOption func(*TwoFactorCode)

func NewTwoFactorCode(opts ...TwoFactorCodeOption) *TwoFactorCode {
	tfc := &TwoFactorCode{}
	for _, opt := range opts {
		opt(tfc)
	}
	return tfc
}

func WithCode(code int) TwoFactorCodeOption {
	return func(tfc *TwoFactorCode) {
		tfc.Code = code
	}
}

func WithType(typ string) TwoFactorCodeOption {
	return func(tfc *TwoFactorCode) {
		tfc.Type = typ
	}
}

func WithExpiredAt(exp time.Time) TwoFactorCodeOption {
	return func(tfc *TwoFactorCode) {
		tfc.ExpiredAt = exp
	}
}

func WithAccount(account Accounts) TwoFactorCodeOption {
	return func(tfc *TwoFactorCode) {
		tfc.Account = account
	}
}
