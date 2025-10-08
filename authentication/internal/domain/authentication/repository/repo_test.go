package repository_test

import (
	"testing"

	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication"
	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/repository"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb/mocks"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
	"github.com/stretchr/testify/suite"
)

type TestAuthenticationSuite struct {
	suite.Suite
	authRepo authentication.AuthenticatorRepository
	mockDB   *mocks.SqlDB[*models.Accounts]
}

func (s *TestAuthenticationSuite) SetupSuite() {
	s.mockDB = &mocks.SqlDB[*models.Accounts]{}
	s.authRepo = repository.NewAuthenticationRepo(s.mockDB)
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
			s.mockDB.On("Create", tt.account).Return(nil).Once()
			err := s.authRepo.Create(tt.account)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}
			s.mockDB.AssertCalled(s.T(), "Create", tt.account)
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
			s.mockDB.On("Create", tt.account).Return(nil).Once()
			_ = s.authRepo.Create(tt.account)
			s.mockDB.On("WhereField", "email", tt.account.Email).Return(tt.account, nil).Once()
			found, err := s.authRepo.Find("email", tt.account.Email)
			if !tt.wantErr {
				s.NoError(err)
				s.Equal(tt.account.Username, found.Username)
			} else {
				s.Error(err)
			}
			s.mockDB.AssertCalled(s.T(), "WhereField", "email", tt.account.Email)
		})
	}
}

func TestAuthenticationSuite_Run(t *testing.T) {
	suite.Run(t, new(TestAuthenticationSuite))
}
