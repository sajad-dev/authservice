package service_test

import (
	"testing"

	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword"
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/dto/gen/request"
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/mocks"
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/service"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/crypto"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/crypto/hs256"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/hashing"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/hashing/sha256"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/authentication/internal/shared/helpers/timeutil"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestForgetPasswordSuite struct {
	suite.Suite
	service  forgetpassword.ForgetPasswordService
	repoMock *mocks.ForgetPasswordRepository
	crypto   crypto.Crypto
	hash     hashing.Hashing
}

func (s *TestForgetPasswordSuite) SetupSuite() {
	s.repoMock = new(mocks.ForgetPasswordRepository)

	hash, _ := s.hash.Sum([]byte("pass"))
	account := models.NewAccounts(
		models.WithEmail("test@email.com"),
		models.WithUsername("test"),
		models.WithPassword(hash),
	)

	s.repoMock.On("Find", "email", "test@email.com").Return(account, nil)
	s.repoMock.On("FindById", 0).Return(account, nil)
	s.repoMock.On("Update", mock.Anything).Return(nil)

	s.service = service.NewForgetPasswordSvc(
		s.repoMock,
		s.hash,
		s.crypto,
	)
}

func (s *TestForgetPasswordSuite) TestForgetPassword() {
	tests := []struct {
		name    string
		req     request.ForgetRequest
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req: request.ForgetRequest{
				Email: "test@email.com",
			},
			wantErr: false,
			wantMsg: messages.SUCCESS_PASSWORD_RESET_EMAIL,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			resp, err := s.service.Forget(tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}
			s.Equal(tt.wantMsg, resp.Msg)
			s.Equal(int32(statuscode.SUCCESSFUL), resp.Code)
		})
	}
}

func (s *TestForgetPasswordSuite) TestResetPassword() {
	token, err := s.crypto.Generate(
		map[string]string{
			"type": "ForgetPassword",
			"id":   "0",
		},
		timeutil.TokenExpire(),
	)
	s.NoError(err)
	tests := []struct {
		name    string
		req     request.ResetRequest
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req: request.ResetRequest{
				Password: "newpassword",
				Token:    token},
			wantErr: false,
			wantMsg: messages.SUCCESS_PASSWORD_RESET,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			resp, err := s.service.Reset(tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}
			s.Equal(tt.wantMsg, resp.Msg)
			s.Equal(int32(statuscode.SUCCESSFUL), resp.Code)
		})
	}
}

func TestForgetPasswordSuite_Run(t *testing.T) {
	su := &TestForgetPasswordSuite{
		crypto: hs256.NewJWT([]byte("this is a secret")),
		hash:   sha256.NewSha256(),
	}
	suite.Run(t, su)
}
