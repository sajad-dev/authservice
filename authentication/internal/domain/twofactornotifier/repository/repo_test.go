package repository_test

import (
	"testing"
	"time"

	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/repository"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb/mocks"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
	"github.com/stretchr/testify/suite"
)

type TestTwoFactorNotifierSuite struct {
	suite.Suite
	repo       twofactornotifier.TwoFactorNotifierRepository
	mockDBAcc  *mocks.SqlDB[*models.Accounts]
	mockDBCode *mocks.SqlDB[*models.TwoFactorCode]
}

func (s *TestTwoFactorNotifierSuite) SetupSuite() {
	s.mockDBAcc = &mocks.SqlDB[*models.Accounts]{}
	s.mockDBCode = &mocks.SqlDB[*models.TwoFactorCode]{}
	s.repo = repository.NewTwoFactorNotifierRepo(s.mockDBAcc, s.mockDBCode)
}

func (s *TestTwoFactorNotifierSuite) TestCreateCode() {
	tests := []struct {
		name string
		code *models.TwoFactorCode
	}{
		{
			name: "Create new two factor code successfully",
			code: models.NewTwoFactorCode(
				models.WithCode(123456),
				models.WithType("email"),
				models.WithAccount(
					*models.NewAccounts(
						models.WithEmail("test@gmail.com"),
						models.WithUsername("test"),
						models.WithPassword("pass"),
					),
				),
				models.WithExpiredAt(time.Now().Add(time.Minute*5)),
			),
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockDBCode.On("Create", tt.code).Return(nil).Once()
			err := s.repo.CreateCode(tt.code)
			s.NoError(err)
			s.mockDBCode.AssertCalled(s.T(), "Create", tt.code)
		})
	}
}

func (s *TestTwoFactorNotifierSuite) TestFindById() {
	tests := []struct {
		name    string
		account *models.Accounts
		wantErr bool
	}{
		{
			name: "Find account by ID",
			account: models.NewAccounts(
				models.WithUsername("twofactoruser"),
				models.WithEmail("twofactor@example.com"),
				models.WithPassword("password123"),
			),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockDBAcc.On("Create", tt.account).Return(nil).Once()
			_ = s.repo.(*repository.TwoFactorNotifierRepo).DBAccount.Create(tt.account)
			s.mockDBAcc.On("GetByID", int(tt.account.ID)).Return(tt.account, nil).Once()
			found, err := s.repo.FindById(int(tt.account.ID))
			if !tt.wantErr {
				s.NoError(err)
				s.Equal(tt.account.ID, found.ID)
			} else {
				s.Error(err)
			}
			s.mockDBAcc.AssertCalled(s.T(), "GetByID", int(tt.account.ID))
		})
	}
}

func (s *TestTwoFactorNotifierSuite) TestRemoveExpierd() {
	tests := []struct {
		name    string
		account *models.Accounts
		code    *models.TwoFactorCode
		wantErr bool
	}{
		{
			name: "Remove expired code successfully",
			account: models.NewAccounts(
				models.WithUsername("expireuser"),
				models.WithEmail("expire@example.com"),
				models.WithPassword("password123"),
			),
			code: models.NewTwoFactorCode(
				models.WithCode(111112),
				models.WithType("email"),
				models.WithExpiredAt(time.Now().Add(-time.Minute)),
			),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockDBAcc.On("Create", tt.account).Return(nil).Once()
			_ = s.repo.(*repository.TwoFactorNotifierRepo).DBAccount.Create(tt.account)
			tt.code.Account = *tt.account
			s.mockDBCode.On("Create", tt.code).Return(nil).Once()
			err := s.repo.CreateCode(tt.code)
			s.NoError(err)
			s.mockDBCode.On("RemoveExpierd", "account_id", int(tt.account.ID)).Return(nil).Once()
			err = s.repo.RemoveExpierd(int(tt.account.ID))
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}
			s.mockDBCode.AssertCalled(s.T(), "RemoveExpierd", "account_id", int(tt.account.ID))
		})
	}
}

func TestTwoFactorNotifierSuite_Run(t *testing.T) {
	suite.Run(t, new(TestTwoFactorNotifierSuite))
}
