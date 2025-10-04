package postgres

import (
	"fmt"

	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/lib/pq"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/setupdb"
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

func (d *SetupPostgres) Connection() (*gormadapter.Adapter, error) {
	fmt.Printf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable \n", d.Username, d.Password, d.Host, d.Port, d.DbName)
	adapter, err := gormadapter.NewAdapter(
		"postgres",
		fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", d.Username, d.Password, d.Host, d.Port, d.DbName),
	)

	if err != nil {
		return nil, err
	}
	return adapter, nil
}

var _ setupdb.SetupDB[*gormadapter.Adapter] = &SetupPostgres{}
