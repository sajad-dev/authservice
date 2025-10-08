package setuppostgres

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SetupPostgres struct {
	Port     int
	Username string
	Password string
	Host     string
	DbName   string
}

func NewSetupPostgres(port int, username, password, host, dbname string) *SetupPostgres {
	return &SetupPostgres{
		Port:     port,
		Username: username,
		Password: password,
		Host:     host,
		DbName:   dbname,
	}
}

var (
	conn *gorm.DB
	once sync.Once
)

func (d *SetupPostgres) _connection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		d.Host, d.Username, d.Password, d.DbName, d.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		return db, nil
	}

	log.Printf("Database '%s' not found, trying to create it...", d.DbName)
	defaultDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%d sslmode=disable",
		d.Host, d.Username, d.Password, d.Port)

	tempDB, err := gorm.Open(postgres.Open(defaultDSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("cannot connect to default database: %v", err)
	}

	createDBQuery := fmt.Sprintf("CREATE DATABASE %s;", d.DbName)
	if err := tempDB.Exec(createDBQuery).Error; err != nil {
		return nil, fmt.Errorf("failed to create database '%s': %v", d.DbName, err)
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to new database '%s': %v", d.DbName, err)
	}

	log.Printf("Database '%s' created and connected successfully!", d.DbName)
	return db, nil
}

func (d *SetupPostgres) Connection() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		conn, err = d._connection()
	})
	return conn, err
}

