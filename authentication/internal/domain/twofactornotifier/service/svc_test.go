package service_test

import (
	"testing"

	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/mocks"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier"

	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/dto/gen/request"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/service"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/crypto"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/crypto/hs256"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/hashing"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/hashing/sha256"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authentication/internal/shared/helpers/timeutil"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestTwoFactorNotifierSuite struct {
	suite.Suite
	twoFactorNotifierService twofactornotifier.TwoFactorNotifierService
	repoMock                 *mocks.TwoFactorNotifierRepository
	crypto                   crypto.Crypto
	hash                     hashing.Hashing
}

func (s *TestTwoFactorNotifierSuite) SetupSuite() {
	s.repoMock = new(mocks.TwoFactorNotifierRepository)

	account := models.NewAccounts(
		models.WithEmail("test@email.com"),
		models.WithUsername("test"),
	)

	s.repoMock.On("FindById", 0).Return(account, nil)
	s.repoMock.On("RemoveExpierd", 0).Return(nil)
	s.repoMock.On("CreateCode", mock.MatchedBy(func(code *models.TwoFactorCode) bool {
		return code.Account.Email == "test@email.com"
	})).Return(nil)

	s.twoFactorNotifierService = service.NewTwoFactorNotifierService(
		s.repoMock,
		s.crypto,
	)
}

func (s *TestTwoFactorNotifierSuite) TestEmailNotifierService() {
	tests := []struct {
		name        string
		token       string
		wantMessage string
		wantErr     bool
	}{
		{
			name:        "Success",
			token:       "valid-token",
			wantMessage: messages.EMAIL_TWO_FACTORT_SUCCESSFUL,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			token, err := s.crypto.Generate(map[string]string{
				"type": "TwoFactor",
				"id":   "0",
			}, timeutil.TokenExpire())
			s.NoError(err)

			reqNotifier := request.NotifierEmailRequest{
				Token: token,
			}

			respNotifier, err := s.twoFactorNotifierService.NotifierEmail(reqNotifier)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}
			s.Equal(tt.wantMessage, respNotifier.Msg)
		})
	}
}

func TestTwoFactorNotifierSuite_Run(t *testing.T) {
	su := &TestTwoFactorNotifierSuite{
		crypto: hs256.NewJWT([]byte("haha this secret")),
		hash:   sha256.NewSha256(),
	}
	suite.Run(t, su)
}

