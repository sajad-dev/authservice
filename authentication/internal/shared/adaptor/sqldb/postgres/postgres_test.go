package postgres_test

import (
	"testing"

	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb/postgres"
	"gorm.io/gorm"
)

type TestModel struct {
	ID    int
	Name  string
	Value string
}

type mockDB struct {
	createErr error
	whereErr  error
	firstErr  error
	saveErr   error
	deleteErr error
	findRows  []TestModel
}

func (m *mockDB) Create(value interface{}) *gorm.DB {
	return &gorm.DB{Error: m.createErr}
}

func (m *mockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	db := &gorm.DB{Error: m.whereErr}
	db.Statement = &gorm.Statement{Dest: &m.findRows}
	return db
}

func (m *mockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return &gorm.DB{Error: m.whereErr}
}

func (m *mockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return &gorm.DB{Error: m.firstErr}
}

func (m *mockDB) Save(value interface{}) *gorm.DB {
	return &gorm.DB{Error: m.saveErr}
}

func (m *mockDB) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return &gorm.DB{Error: m.deleteErr}
}

func TestPostgres_CRUD(t *testing.T) {
	mock := &gorm.DB{}
	pg := postgres.NewPostgres[TestModel](mock)

	tests := []struct {
		name        string
		action      string
		expectError bool
	}{
		{"Create", "create", false},
		{"Where", "where", false},
		{"WhereField", "wherefield", false},
		{"Save", "save", false},
		{"GetByID", "getbyid", false},
		{"Delete", "delete", false},
		{"RemoveExpired", "removeexpired", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.action {
			case "create":
				err := pg.Create(TestModel{Name: "A"})
				if (err != nil) != tt.expectError {
					t.Errorf("Create() error = %v, expectError %v", err, tt.expectError)
				}
			case "where":
				_, err := pg.Where(TestModel{Name: "A"})
				if (err != nil) != tt.expectError {
					t.Errorf("Where() error = %v, expectError %v", err, tt.expectError)
				}
			case "wherefield":
				_, err := pg.WhereField("name", "A")
				if (err != nil) != tt.expectError {
					t.Errorf("WhereField() error = %v, expectError %v", err, tt.expectError)
				}
			case "save":
				err := pg.Save(TestModel{Name: "A"})
				if (err != nil) != tt.expectError {
					t.Errorf("Save() error = %v, expectError %v", err, tt.expectError)
				}
			case "getbyid":
				_, err := pg.GetByID(1)
				if (err != nil) != tt.expectError {
					t.Errorf("GetByID() error = %v, expectError %v", err, tt.expectError)
				}
			case "delete":
				err := pg.Delete(1)
				if (err != nil) != tt.expectError {
					t.Errorf("Delete() error = %v, expectError %v", err, tt.expectError)
				}
			case "removeexpired":
				err := pg.RemoveExpierd("expired_at", 1)
				if (err != nil) != tt.expectError {
					t.Errorf("RemoveExpierd() error = %v, expectError %v", err, tt.expectError)
				}
			}
		})
	}
}
