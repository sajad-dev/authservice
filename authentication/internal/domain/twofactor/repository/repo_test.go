package repository_test

import (
	"testing"
	"time"

	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/repository"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb/sqlite"
	"github.com/sajad-dev/authservice/authentication/internal/shared/helpers/testhelper/testdb"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
	"github.com/stretchr/testify/suite"
)

type TestTwoFactorSuite struct {
	suite.Suite
	repo twofactor.TwoFactorRepository
}

func (s *TestTwoFactorSuite) SetupSuite() {
	db, err := testdb.Conn(&models.Accounts{}, &models.TwoFactorCode{})
	s.NoError(err)

	s.repo = repository.NewTwoFactorRepo(sqlite.NewSqlite[*models.Accounts](db), sqlite.NewSqlite[*models.TwoFactorCode](db))
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
			dbErr := s.repo.(*repository.TwoFactorRepo).DBCode.Create(tt.code)
			s.NoError(dbErr)

			found, err := s.repo.FindByCode(tt.code.Code, tt.code.Type)
			if !tt.wantErr {
				s.NoError(err)
				s.GreaterOrEqual(len(found), 1)
				s.Equal(tt.code.Code, found[0].Code)
			} else {
				s.Error(err)
			}
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
			dbErr := s.repo.(*repository.TwoFactorRepo).DBAccount.Create(tt.account)
			s.NoError(dbErr)

			found, err := s.repo.FindById(int(tt.account.ID))
			if !tt.wantErr {
				s.NoError(err)
				s.Equal(tt.account.ID, found.ID)
			} else {
				s.Error(err)
			}
		})
	}
}

func TestTwoFactorSuite_Run(t *testing.T) {
	suite.Run(t, new(TestTwoFactorSuite))
}

