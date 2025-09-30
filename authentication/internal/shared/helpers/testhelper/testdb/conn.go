package testdb

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Conn(tables ...any) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(tables...)
	if err != nil {
		return nil, err
	}
	return db, nil
}
