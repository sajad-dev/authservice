package postgres_test

import (
	"testing"

	"github.com/sajad-dev/authservice/internal/adaptor/sqldb/postgres"
)

type TestUser struct{}

func TestPostgres_Create(t *testing.T) {
	x := postgres.Postgres{DB: nil}
	x.Create(TestUser{})
}
