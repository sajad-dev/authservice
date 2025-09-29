package repository_test

import (
	"testing"

	"github.com/sajad-dev/authservice/internal/domain/authentication"
	"github.com/sajad-dev/authservice/internal/domain/authentication/repository"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/sqldb/sqlite"
	"github.com/sajad-dev/authservice/internal/shared/helpers/testhelper/testdb"
	"github.com/sajad-dev/authservice/internal/shared/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestAuthenticationSuite struct {
	suite.Suite
	authRepo authentication.AuthenticatorRepository
}

func (s *TestAuthenticationSuite) SetupSuite() {
	db, err := testdb.Conn(&models.Accounts{})
	assert.NoError(s.T(), err)

	s.authRepo = repository.NewAuthenticationRepo(sqlite.NewSqlite[*models.Accounts](db))
}

func (s *TestAuthenticationSuite) TestCreate() {
	tests := []struct {
		name    string
		account *models.Accounts
		wantErr bool
	}{
		{
			name: "Create new account successfully",
			account: models.NewAccounts(
				models.WithUsername("authuser"),
				models.WithEmail("auth@example.com"),
				models.WithPassword("password123"),
			),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := s.authRepo.Create(tt.account)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}
		})
	}
}

func (s *TestAuthenticationSuite) TestFind() {
	tests := []struct {
		name    string
		account *models.Accounts
		wantErr bool
	}{
		{
			name: "Find account by email",
			account: models.NewAccounts(
				models.WithUsername("fielduser"),
				models.WithEmail("field@example.com"),
				models.WithPassword("password123"),
			),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := s.authRepo.Create(tt.account)
			s.NoError(err)

			found, err := s.authRepo.Find("email", tt.account.Email)
			if !tt.wantErr {
				s.NoError(err)
				s.Equal(tt.account.Username, found.Username)
			} else {
				s.Error(err)
			}
		})
	}
}

func TestAuthenticationSuite_Run(t *testing.T) {
	suite.Run(t, new(TestAuthenticationSuite))
}
