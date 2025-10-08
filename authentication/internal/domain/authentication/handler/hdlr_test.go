package handler_test

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication"
	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/authenticationproto"
	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/dto/gen/response"
	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/handler"
	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/mocks"
	"github.com/sajad-dev/authservice/authentication/internal/shared/validation"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestAuthenticationHandlerSuite struct {
	suite.Suite
	handler  authentication.AuthenticatorHandler
	mocksSVC *mocks.AuthenticatorService
}

func (s *TestAuthenticationHandlerSuite) SetupSuite() {
	mocksSVC := new(mocks.AuthenticatorService)
	s.mocksSVC = mocksSVC

	s.handler = handler.NewAuthenticationHdlr(mocksSVC, validate.NewValidate(validator.New()))
}

func (s *TestAuthenticationHandlerSuite) TestLogin() {
	tests := []struct {
		name    string
		req     *authenticationproto.LoginRequest
		res     response.LoginResponse
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req:  &authenticationproto.LoginRequest{Username: "test@example.com", Password: "123456"},
			res: response.LoginResponse{
				Token: "token",

				Code: statuscode.SUCCESSFUL,
				Msg:  messages.SUCCESS_LOGIN,
				Data: models.AccountFiltered{
					Email:     "test@email.com",
					FirstName: "test",
					LastName:  "test",
					Username:  "test",
				},
			},

			wantErr: false,
			wantMsg: messages.SUCCESS_LOGIN,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSVC.On("Login", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.Login(tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}

			s.Equal(tt.wantMsg, resp.Msg)
		})
	}
}

func (s *TestAuthenticationHandlerSuite) TestRegister() {
	tests := []struct {
		name    string
		req     *authenticationproto.RegisterRequest
		res     response.RegisterResponse
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req:  &authenticationproto.RegisterRequest{Email: "test@example.com", Password: "123456", PasswordConfirmation: "123456"},
			res: response.RegisterResponse{
				Token: "token",
				Code:  statuscode.SUCCESSFUL,
				Msg:   messages.SUCCESS_LOGIN,
				Data: models.AccountFiltered{
					Email:     "test@email.com",
					FirstName: "test",
					LastName:  "test",
					Username:  "test",
					TwoFactor: []string{},
				},
			},
			wantErr: false,
			wantMsg: messages.SUCCESS_LOGIN,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSVC.On("Register", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.Register(tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}

			s.Equal(tt.wantMsg, resp.Msg)
		})
	}
}

func TestAuthenticationHandlerSuite_Run(t *testing.T) {
	suite.Run(t, new(TestAuthenticationHandlerSuite))
}
