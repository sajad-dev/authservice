package service_test

import (
	"testing"

	"github.com/lib/pq"
	"github.com/sajad-dev/authservice/internal/domain/authentication"
	"github.com/sajad-dev/authservice/internal/domain/authentication/dto/gen/request"
	"github.com/sajad-dev/authservice/internal/domain/authentication/mocks"
	"github.com/sajad-dev/authservice/internal/domain/authentication/service"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/crypto"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/crypto/hs256"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/hashing"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/hashing/sha256"
	"github.com/sajad-dev/authservice/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/internal/shared/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestAuthenticationSuite struct {
	suite.Suite
	authService authentication.AuthenticatorService
	repoMock    *mocks.AuthenticatorRepository
	crypto      crypto.Crypto
	hash        hashing.Hashing
}

func (s *TestAuthenticationSuite) SetupSuite() {
	s.repoMock = new(mocks.AuthenticatorRepository)

	hash, _ := s.hash.Sum([]byte("pass"))
	account := models.NewAccounts(
		models.WithEmail("test@email.com"),
		models.WithUsername("testuser"),
		models.WithPassword(hash),
	)
	account2FA := models.NewAccounts(
		models.WithEmail("2fa@email.com"),
		models.WithUsername("user2fa"),
		models.WithPassword(hash),
		models.WithTwoFactor(pq.StringArray{
			"email",
		}),
	)

	s.repoMock.On("Create", mock.Anything).Return(nil)
	s.repoMock.On("Find", "email", "test@email.com").Return(account, nil)
	s.repoMock.On("Find", "username", "testuser").Return(account, nil)
	s.repoMock.On("Find", "email", "2fa@email.com").Return(account2FA, nil)
	s.repoMock.On("Find", "username", "user2fa").Return(account2FA, nil)

	s.authService = service.NewAuthService(
		s.repoMock,
		s.crypto,
		s.hash,
	)
}

func (s *TestAuthenticationSuite) TestRegisterService() {
	req := request.RegisterRequest{
		Email:                "test@email.com",
		Username:             "testuser",
		Password:             "pass",
		PasswordConfirmation: "pass",
	}

	resp, err := s.authService.Register(req)
	s.NoError(err)
	s.Equal(messages.LOGIN_IS_SUCCESSFUL, resp.Msg)
	s.NotEmpty(resp.Token)
	s.Equal("test@email.com", resp.Data.Email)
}

func (s *TestAuthenticationSuite) TestLoginService_Success() {
	req := request.LoginRequest{
		Username: "testuser",
		Password: "pass",
	}

	resp, err := s.authService.Login(req)
	s.NoError(err)
	s.Equal(messages.LOGIN_IS_SUCCESSFUL, resp.Msg)
	s.NotEmpty(resp.Token)
	s.Equal("testuser", resp.Data.Username)
}

func (s *TestAuthenticationSuite) TestLoginService_WrongPassword() {
	req := request.LoginRequest{
		Username: "testuser",
		Password: "worngpass",
	}

	resp, err := s.authService.Login(req)
	s.NoError(err)
	s.Equal(resp.Msg, messages.USERNAME_OR_PASSWORD_IS_WORNG)
	s.Empty(resp.Token)
}

func (s *TestAuthenticationSuite) TestLoginService_TwoFactor() {
	req := request.LoginRequest{
		Username: "user2fa",
		Password: "pass",
	}

	resp, err := s.authService.Login(req)

	s.NoError(err)
	s.Equal(messages.LOGIN_WITH_TWO_FACTOR, resp.Msg)
	s.NotEmpty(resp.Token)
	s.ElementsMatch([]string{"email"}, resp.Data.TwoFactor)
}

func TestAuthenticationSuite_Run(t *testing.T) {
	su := &TestAuthenticationSuite{
		crypto: hs256.NewJWT([]byte("haha this secret")),
		hash:   sha256.NewSha256(),
	}
	suite.Run(t, su)
}
