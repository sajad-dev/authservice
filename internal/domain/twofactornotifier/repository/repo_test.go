package repository_test

import (
	"testing"
	"time"

	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier"
	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier/repository"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/sqldb/sqlite"
	"github.com/sajad-dev/authservice/internal/shared/helpers/testhelper/testdb"
	"github.com/sajad-dev/authservice/internal/shared/models"
	"github.com/stretchr/testify/suite"
)

type TestTwoFactorNotifierSuite struct {
	suite.Suite
	repo twofactornotifier.TwoFactorNotifierRepository
}

func (s *TestTwoFactorNotifierSuite) SetupSuite() {
	db, err := testdb.Conn(&models.Accounts{}, &models.TwoFactorCode{})
	s.NoError(err)

	s.repo = repository.NewTwoFactorNotifierRepo(
		sqlite.NewSqlite[*models.Accounts](db),
		sqlite.NewSqlite[*models.TwoFactorCode](db),
	)
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
			err := s.repo.CreateCode(tt.code)
			s.NoError(err)
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
			dbErr := s.repo.(*repository.TwoFactorNotifierRepo).DBAccount.Create(tt.account)
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
				models.WithAccount(
					*models.NewAccounts(
						models.WithUsername("expireuser"),
						models.WithEmail("expire@example.com"),
						models.WithPassword("password123"),
					),
				),
				models.WithExpiredAt(time.Now().Add(-time.Minute)),
			),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			dbErr := s.repo.(*repository.TwoFactorNotifierRepo).DBAccount.Create(tt.account)
			s.NoError(dbErr)

			tt.code.Account = *tt.account
			err := s.repo.CreateCode(tt.code)
			s.NoError(err)

			err = s.repo.RemoveExpierd(int(tt.account.ID))
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}
		})
	}
}

func TestTwoFactorNotifierSuite_Run(t *testing.T) {
	suite.Run(t, new(TestTwoFactorNotifierSuite))
}
