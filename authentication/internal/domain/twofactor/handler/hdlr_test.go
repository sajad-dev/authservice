package handler_test

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/twofactorproto"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/dto/gen/response"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/handler"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/mocks"
	"github.com/sajad-dev/authservice/authentication/internal/shared/validation"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/statuscode"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestTwoFactorHandlerSuite struct {
	suite.Suite
	handler  twofactor.TwoFactorHandler
	mocksSVC *mocks.TwoFactorService
}

func (s *TestTwoFactorHandlerSuite) SetupSuite() {
	mocksSVC := new(mocks.TwoFactorService)
	s.mocksSVC = mocksSVC

	s.handler = handler.NewTwoFactorHdlr(mocksSVC, validate.NewValidate(validator.New()))
}

func (s *TestTwoFactorHandlerSuite) TestEmail() {
	tests := []struct {
		name    string
		req     *twofactorproto.EmailRequest
		res     response.TwoFactorResponse
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req:  &twofactorproto.EmailRequest{Code: 1111},
			res: response.TwoFactorResponse{
				Code: statuscode.SUCCESSFUL,
				Msg:  messages.SUCCESS_LOGIN,
			},
			wantErr: false,
			wantMsg: messages.SUCCESS_LOGIN,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSVC.On("Email", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.Email(tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}

			s.Equal(tt.wantMsg, resp.Msg)
		})
	}
}

func (s *TestTwoFactorHandlerSuite) TestGoogle() {
	tests := []struct {
		name    string
		req     *twofactorproto.GoogleRequest
		res     response.TwoFactorResponse
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req:  &twofactorproto.GoogleRequest{Code: 1234, Token: "123456"},
			res: response.TwoFactorResponse{
				Code: statuscode.SUCCESSFUL,
				Msg:  messages.SUCCESS_LOGIN,
			},
			wantErr: false,
			wantMsg: messages.SUCCESS_LOGIN,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSVC.On("Google", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.Google(tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}

			s.Equal(tt.wantMsg, resp.Msg)
		})
	}
}

func TestTwoFactorHandlerSuite_Run(t *testing.T) {
	suite.Run(t, new(TestTwoFactorHandlerSuite))
}

