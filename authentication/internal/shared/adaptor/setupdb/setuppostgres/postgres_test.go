package setuppostgres_test

import (
	"errors"
	"testing"

	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/setupdb/setuppostgres"
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
			username:    "user",
			password:    "pass",
			host:        "localhost",
			dbName:      "testdb",
			expectError: false,
		},
		{
			name:        "invalid port",
			port:        -1,
			username:    "user",
			password:    "pass",
			host:        "localhost",
			dbName:      "testdb",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup := setuppostgres.NewSetupPostgres(tt.port, tt.username, tt.password, tt.host, tt.dbName)
			db, err := setup.Connection()
			if tt.expectError && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !tt.expectError && db == nil {
				t.Errorf("expected db instance but got nil")
			}
		})
	}
}
