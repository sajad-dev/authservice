package postgres

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/setupdb"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs"
)

type SetupPostgres struct {
	Port     string
	Username string
	Password string
	Host     string
	DbName   string
}

func NewSetupPostgres(port string, username string, password string, host string, dbname string) *SetupPostgres {
	return &SetupPostgres{
		Port:     port,
		Username: username,
		Password: password,
		Host:     host,
		DbName:   dbname,
	}
}

func (d *SetupPostgres) Connection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("Connecting with DSN: host=%q user=%q password=%q dbname=%q port=%q sslmode=disable\n",
		d.Host, d.Username, d.Password, "postgres", d.Port)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, errs.Err(err)
	}

	var exists bool
	db.Raw("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = ?)", d.DbName).Scan(&exists)

	if !exists {
		err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", d.DbName)).Error
		if err != nil {
			return nil, errs.Err(err)
		}
	}

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", d.Host, d.Username, d.Password, d.DbName, d.Port)
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		return nil, errs.Err(err)
	}
	return db, nil
}

var _ setupdb.SetupDB[*gorm.DB] = &SetupPostgres{}
