package service_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/dto/gen/request"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/mocks"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/service"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/crypto"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/crypto/hs256"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/hashing"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/hashing/sha256"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authentication/internal/shared/helpers/timeutil"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
	"github.com/stretchr/testify/suite"
)

type TestTwoFactorSuite struct {
	suite.Suite
	twoFactorService twofactor.TwoFactorService
	repoMock         *mocks.TwoFactorRepository
	googleAuthSecret *otp.Key
	crypto           crypto.Crypto
	hash             hashing.Hashing
}

func (s *TestTwoFactorSuite) SetupSuite() {
	s.repoMock = new(mocks.TwoFactorRepository)

	var err error
	s.googleAuthSecret, err = totp.Generate(totp.GenerateOpts{
		Issuer:      "My butty app",
		AccountName: "test@email.com",
	})
	s.NoError(err)

	account := models.NewAccounts(
		models.WithEmail("test@email.com"),
		models.WithUsername("test"),
		models.WithGoogleAuthSecretKey(s.googleAuthSecret.Secret()),
	)

	s.repoMock.On("FindById", 0).Return(account, nil)

	s.twoFactorService = service.NewTwoFactorSvc(
		s.repoMock,
		s.crypto,
		s.hash,
	)
}

func (s *TestTwoFactorSuite) TestEmailService() {
	tests := []struct {
		name    string
		req     request.EmailRequest
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req: request.EmailRequest{
				Code: 1111,
			},
			wantErr: false,
			wantMsg: messages.SUCCESS_LOGIN,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			account := models.NewAccounts(models.WithEmail("test@email.com"), models.WithUsername("test"))

			twoFactorCode := models.NewTwoFactorCode(models.WithCode(1111), models.WithType(string(constants.EMAIL_TWO_FACTOR_CODE)), models.WithAccount(*account))
			
			s.repoMock.On("FindByCode", 1111, string(constants.EMAIL_TWO_FACTOR_CODE)).Return([]*models.TwoFactorCode{twoFactorCode}, nil)

			resp, err := s.twoFactorService.Email(tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}
			s.Equal(tt.wantMsg, resp.Msg)
		})
	}
}

// func (s *TestTwoFactorSuite) TestGoogleService() {
// 	googleCode, err := totp.GenerateCode(s.googleAuthSecret.Secret(), time.Now())
// 	s.NoError(err)
	
// 	googleCodeInt, err := strconv.Atoi(googleCode)
// 	s.NoError(err)

// 	token, err := s.crypto.Generate(map[string]string{ "type": "TwoFactor", "id": "0", }, timeutil.TokenExpire())
// 	s.NoError(err)

// 	tests := []struct {
// 		name    string
// 		req     request.GoogleRequest
// 		wantErr bool
// 		wantMsg string
// 	}{
// 		{
// 			name: "Success",
// 			req: request.GoogleRequest{
// 				Code:  int32(googleCodeInt),
// 				Token: token,
// 			},
// 			wantErr: false,
// 			wantMsg: messages.SUCCESS_LOGIN,
// 		},
// 	}

// 	for _, tt := range tests {
// 		s.Run(tt.name, func() {
// 			resp, err := s.twoFactorService.Google(tt.req)
// 			if !tt.wantErr {
// 				s.NoError(err)
// 			} else {
// 				s.Error(err)
// 			}
// 			s.Equal(tt.wantMsg, resp.Msg)
// 		})
// 	}
// }

func TestTwoFactorSuite_Run(t *testing.T) {
	su := &TestTwoFactorSuite{
		crypto: hs256.NewJWT([]byte("haha this secret")),
		hash:   sha256.NewSha256(),
	}
	suite.Run(t, su)
}
