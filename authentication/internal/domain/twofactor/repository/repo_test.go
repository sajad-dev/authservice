package repository_test

import (
	"testing"
	"time"

	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/repository"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb/mocks"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestTwoFactorSuite struct {
	suite.Suite
	repo      twofactor.TwoFactorRepository
	mockDBAcc *mocks.SqlDB[*models.Accounts]
	mockDBCode *mocks.SqlDB[*models.TwoFactorCode]
}

func (s *TestTwoFactorSuite) SetupSuite() {
	s.mockDBAcc = &mocks.SqlDB[*models.Accounts]{}
	s.mockDBCode = &mocks.SqlDB[*models.TwoFactorCode]{}
	s.repo = repository.NewTwoFactorRepo(s.mockDBAcc, s.mockDBCode)
}

func (s *TestTwoFactorSuite) TestFindByCode() {
	tests := []struct {
		name    string
		code    *models.TwoFactorCode
		wantErr bool
	}{
		{
			name: "Find two factor code by email type",
			code: models.NewTwoFactorCode(
				models.WithCode(654321),
				models.WithAccount(
					*models.NewAccounts(
						models.WithEmail("test@gmail.com"),
						models.WithUsername("test"),
						models.WithPassword("pass"),
					),
				),
				models.WithType("email"),
				models.WithExpiredAt(time.Now().Add(time.Minute*5)),
			),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockDBCode.On("Create", tt.code).Return(nil).Once()
			dbErr := s.repo.(*repository.TwoFactorRepo).DBCode.Create(tt.code)
			s.NoError(dbErr)
			s.mockDBCode.On("Where", mock.Anything).Return([]*models.TwoFactorCode{tt.code}, nil).Once()
			found, err := s.repo.FindByCode(tt.code.Code, tt.code.Type)
			if !tt.wantErr {
				s.NoError(err)
				s.GreaterOrEqual(len(found), 1)
				s.Equal(tt.code.Code, found[0].Code)
			} else {
				s.Error(err)
			}
			s.mockDBCode.AssertCalled(s.T(), "Where", mock.Anything)
		})
	}
}

func (s *TestTwoFactorSuite) TestFindById() {
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
			dbErr := s.repo.(*repository.TwoFactorRepo).DBAccount.Create(tt.account)
			s.NoError(dbErr)
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

func TestTwoFactorSuite_Run(t *testing.T) {
	suite.Run(t, new(TestTwoFactorSuite))
}

