package postgres

import (
	"fmt"
	"sync"

	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/lib/pq"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/setupdb"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs"
)

type SetupPostgres struct {
	Port     int
	Username string
	Password string
	Host     string
	DbName   string
}

func NewSetupPostgres(port int, username string, password string, host string, dbname string) *SetupPostgres {
	return &SetupPostgres{
		Port:     port,
		Username: username,
		Password: password,
		Host:     host,
		DbName:   dbname,
	}
}

var conn *gormadapter.Adapter
var once sync.Once

func (d *SetupPostgres) _connection() (*gormadapter.Adapter, error) {
	fmt.Printf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable \n", d.Username, d.Password, d.Host, d.Port, d.DbName)
	adapter, err := gormadapter.NewAdapter(
		"postgres",
		fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", d.Username, d.Password, d.Host, d.Port, d.DbName),
	)

	if err != nil {
		return nil, errs.Err(err)
	}
	return adapter, nil
}

func (d *SetupPostgres) Connection() (*gormadapter.Adapter, error) {
	var err error
	
	once.Do(func() {
		conn, err = d._connection()
	})

	return conn, err

}

var _ setupdb.SetupDB[*gormadapter.Adapter] = &SetupPostgres{}
