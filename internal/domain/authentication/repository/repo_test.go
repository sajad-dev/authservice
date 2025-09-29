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
	account := models.NewAccounts(
		models.WithUsername("authuser"),
		models.WithEmail("auth@example.com"),
		models.WithPassword("password123"),
	)
	err := s.authRepo.Create(account)
	assert.NoError(s.T(), err)
}

func (s *TestAuthenticationSuite) TestGetByFields() {
	account := models.NewAccounts(
		models.WithUsername("fielduser"),
		models.WithEmail("field@example.com"),
		models.WithPassword("password123"),
	)
	err := s.authRepo.Create(account)
	assert.NoError(s.T(), err)

	found, err := s.authRepo.Find("email", account.Email)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), account.Username, found.Username)
}

func TestAuthenticationSuite_Run(t *testing.T) {
	suite.Run(t, new(TestAuthenticationSuite))
}
