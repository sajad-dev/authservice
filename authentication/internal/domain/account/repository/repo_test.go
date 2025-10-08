package repository_test

import (
	"errors"
	"testing"

	"github.com/sajad-dev/authservice/authentication/internal/domain/account"
	"github.com/sajad-dev/authservice/authentication/internal/domain/account/repository"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb/mocks"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
	"github.com/stretchr/testify/suite"
)

type TestAccountSuite struct {
	suite.Suite
	accountRepo account.AccountCURDRepository
	mockDB      *mocks.SqlDB[*models.Accounts]
}

func (s *TestAccountSuite) SetupSuite() {
	s.mockDB = &mocks.SqlDB[*models.Accounts]{}
	s.accountRepo = repository.NewAccountRepo(s.mockDB)
}

func (s *TestAccountSuite) TestCreate() {
	tests := []struct {
		name    string
		account *models.Accounts
		wantErr bool
	}{
		{
			name: "Create new account successfully",
			account: models.NewAccounts(
				models.WithUsername("testuser"),
				models.WithEmail("test@example.com"),
				models.WithPassword("password123"),
				models.WithTwoFactor([]string{"email"}),
			),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockDB.On("Create", tt.account).Return(nil).Once()
			err := s.accountRepo.Create(tt.account)

			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}

			s.mockDB.AssertCalled(s.T(), "Create", tt.account)
		})
	}
}

func (s *TestAccountSuite) TestRead() {
	tests := []struct {
		name    string
		account *models.Accounts
		wantErr bool
	}{
		{
			name: "Read created account by ID",
			account: models.NewAccounts(
				models.WithUsername("readuser"),
				models.WithEmail("read@example.com"),
				models.WithPassword("readpass"),
				models.WithTwoFactor([]string{}),
			),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockDB.On("Create", tt.account).Return(nil).Once()
			s.mockDB.On("GetByID", int(tt.account.ID)).Return(tt.account, nil).Once()

			err := s.accountRepo.Create(tt.account)
			s.NoError(err)

			read, err := s.accountRepo.Read(int(tt.account.ID))
			if !tt.wantErr {
				s.NoError(err)
				s.Equal(tt.account.Username, read.Username)
				s.Equal(tt.account.Email, read.Email)
			} else {
				s.Error(err)
			}

			s.mockDB.AssertCalled(s.T(), "GetByID", int(tt.account.ID))
		})
	}
}

func (s *TestAccountSuite) TestUpdate() {
	tests := []struct {
		name      string
		account   *models.Accounts
		updatedTo string
		wantErr   bool
	}{
		{
			name: "Update account username successfully",
			account: models.NewAccounts(
				models.WithUsername("updateuser"),
				models.WithEmail("update@example.com"),
				models.WithPassword("updatepass"),
				models.WithTwoFactor([]string{}),
			),
			updatedTo: "updateduser",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockDB.On("Create", tt.account).Return(nil).Once()
			s.mockDB.On("Save", tt.account).Return(nil).Once()
			s.mockDB.On("GetByID", int(tt.account.ID)).Return(tt.account, nil).Once()

			_ = s.accountRepo.Create(tt.account)
			tt.account.Username = tt.updatedTo
			err := s.accountRepo.Update(tt.account, int(tt.account.ID))

			if !tt.wantErr {
				s.NoError(err)
				updated, _ := s.accountRepo.Read(int(tt.account.ID))
				s.Equal(tt.updatedTo, updated.Username)
			} else {
				s.Error(err)
			}

			s.mockDB.AssertCalled(s.T(), "Save", tt.account)
		})
	}
}

func (s *TestAccountSuite) TestDelete() {
	tests := []struct {
		name    string
		account *models.Accounts
		wantErr bool
	}{
		{
			name: "Delete account successfully",
			account: models.NewAccounts(
				models.WithUsername("deleteuser"),
				models.WithEmail("delete@example.com"),
				models.WithPassword("deletepass"),
				models.WithTwoFactor([]string{"email"}),
			),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockDB.On("Create", tt.account).Return(nil).Once()
			s.mockDB.On("Delete", int(tt.account.ID)).Return(nil).Once()
			s.mockDB.On("GetByID", int(tt.account.ID)).Return(nil, errors.New("not found")).Once()

			_ = s.accountRepo.Create(tt.account)
			err := s.accountRepo.Delete(int(tt.account.ID))

			if !tt.wantErr {
				s.NoError(err)
				_, err = s.accountRepo.Read(int(tt.account.ID))
				s.Error(err)
			} else {
				s.Error(err)
			}

			s.mockDB.AssertCalled(s.T(), "Delete", int(tt.account.ID))
		})
	}
}

func TestAccountSuite_Run(t *testing.T) {
	suite.Run(t, new(TestAccountSuite))
}
