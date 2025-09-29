package service_test

import (
	"testing"

	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier/mocks"
	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier"

	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier/dto/gen/request"
	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier/service"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/crypto"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/crypto/hs256"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/hashing"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/hashing/sha256"
	"github.com/sajad-dev/authservice/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/internal/shared/helpers/timeutil"
	"github.com/sajad-dev/authservice/internal/shared/models"
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
	token, err := s.crypto.Generate(map[string]string{
		"type": "TwoFactor",
		"id":   "0",
	},timeutil.TokenExpire())
	s.NoError(err)

	reqNotifier := request.NotifierEmailRequest{
		Token: token,
	}
	respNotifier, err := s.twoFactorNotifierService.NotifierEmail(reqNotifier)
	s.NoError(err)
	s.Equal(messages.EMAIL_TWO_FACTORT_SUCCESSFUL, respNotifier.Message)
}

func TestTwoFactorNotifierSuite_Run(t *testing.T) {
	su := &TestTwoFactorNotifierSuite{
		crypto: hs256.NewJWT([]byte("haha this secret")),
		hash:   sha256.NewSha256(),
	}
	suite.Run(t, su)
}
