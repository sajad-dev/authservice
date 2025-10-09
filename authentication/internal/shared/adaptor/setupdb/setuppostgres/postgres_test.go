package setuppostgres_test

import (
	"errors"
	"testing"

	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/setupdb/setuppostgres"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type mockDB struct {
	shouldFailOpen bool
	shouldFailExec bool
}

func (m *mockDB) Open(dsn string) (*gorm.DB, error) {
	if m.shouldFailOpen {
		return nil, errors.New("open failed")
	}
	return &gorm.DB{}, nil
}

func TestSetupPostgres_Connection(t *testing.T) {
	tests := []struct {
		name        string
		port        int
		username    string
		password    string
		host        string
		dbName      string
		expectError bool
	}{
		{
			name:        "valid connection",
			port:        5432,
			username:    "root",
			password:    "root",
			host:        "localhost",
			dbName:      "testdb",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup := setuppostgres.NewSetupPostgres(tt.port, tt.username, tt.password, tt.host, tt.dbName)
			_, err := setup.Connection()
			assert.NoError(t, err)

		})
	}
}
