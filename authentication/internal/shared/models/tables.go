package models

import "gorm.io/gorm"

func Migration(db *gorm.DB) error {
	err := db.AutoMigrate(&Accounts{}, &TwoFactorCode{})
	if err != nil {
		return err
	}
	return nil
}
