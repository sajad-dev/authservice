package repository_test

import (
	"testing"

	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword"
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/repository"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb/mocks"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
	"github.com/stretchr/testify/suite"
)

type TestForgetPasswordSuite struct {
	suite.Suite
	repo   forgetpassword.ForgetPasswordRepository
	mockDB *mocks.SqlDB[*models.Accounts]
}

func (s *TestForgetPasswordSuite) SetupSuite() {
	s.mockDB = &mocks.SqlDB[*models.Accounts]{}
	s.repo = repository.NewForgetPasswordRepo(s.mockDB)
}

func (s *TestForgetPasswordSuite) TestFind() {
	tests := []struct {
		name    string
		account *models.Accounts
		column  string
		value   string
		wantErr bool
	}{
		{
			name: "Find account by email",
			account: models.NewAccounts(
				models.WithUsername("finduser"),
				models.WithEmail("find@example.com"),
				models.WithPassword("password123"),
			),
			column:  "email",
			value:   "find@example.com",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockDB.On("Save", tt.account).Return(nil).Once()
			err := s.repo.Update(tt.account)
			s.NoError(err)
			s.mockDB.On("WhereField", tt.column, tt.value).Return(tt.account, nil).Once()
			found, err := s.repo.Find(tt.column, tt.value)
			if !tt.wantErr {
				s.NoError(err)
				s.Equal(tt.account.Email, found.Email)
			} else {
				s.Error(err)
			}
			s.mockDB.AssertCalled(s.T(), "WhereField", tt.column, tt.value)
		})
	}
}

func (s *TestForgetPasswordSuite) TestFindById() {
	tests := []struct {
		name    string
		account *models.Accounts
		wantErr bool
	}{
		{
			name: "Find account by ID",
			account: models.NewAccounts(
				models.WithUsername("iduser"),
				models.WithEmail("id@example.com"),
				models.WithPassword("password123"),
			),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockDB.On("Save", tt.account).Return(nil).Once()
			_ = s.repo.Update(tt.account)
			s.mockDB.On("GetByID", int(tt.account.ID)).Return(tt.account, nil).Once()
			found, err := s.repo.FindById(int(tt.account.ID))
			if !tt.wantErr {
				s.NoError(err)
				s.Equal(tt.account.Username, found.Username)
			} else {
				s.Error(err)
			}
			s.mockDB.AssertCalled(s.T(), "GetByID", int(tt.account.ID))
		})
	}
}

func (s *TestForgetPasswordSuite) TestUpdate() {
	tests := []struct {
		name      string
		account   *models.Accounts
		updatedTo string
		wantErr   bool
	}{
		{
			name: "Update account password successfully",
			account: models.NewAccounts(
				models.WithUsername("updateuser"),
				models.WithEmail("update@example.com"),
				models.WithPassword("oldpassword"),
			),
			updatedTo: "newpassword",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockDB.On("Save", tt.account).Return(nil).Once()
			_ = s.repo.Update(tt.account)
			tt.account.Password = tt.updatedTo
			s.mockDB.On("Save", tt.account).Return(nil).Once()
			err := s.repo.Update(tt.account)
			if !tt.wantErr {
				s.NoError(err)
				s.mockDB.On("GetByID", int(tt.account.ID)).Return(tt.account, nil).Once()
				updated, _ := s.repo.FindById(int(tt.account.ID))
				s.Equal(tt.updatedTo, updated.Password)
			} else {
				s.Error(err)
			}
			s.mockDB.AssertCalled(s.T(), "Save", tt.account)
		})
	}
}

func TestForgetPasswordSuite_Run(t *testing.T) {
	suite.Run(t, new(TestForgetPasswordSuite))
}
