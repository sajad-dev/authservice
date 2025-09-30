package handler_test

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword"
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/dto/gen/response"
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/forgetpasswordproto"
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/handler"
	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/mocks"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/validation/validate"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/statuscode"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestForgetPasswordHandlerSuite struct {
	suite.Suite
	handler  forgetpassword.ForgetPasswordHandler
	mocksSVC *mocks.ForgetPasswordService
}

func (s *TestForgetPasswordHandlerSuite) SetupSuite() {
	mocksSVC := new(mocks.ForgetPasswordService)
	s.mocksSVC = mocksSVC

	s.handler = handler.NewForgetPasswordHandler(mocksSVC, validate.NewValidate(validator.New()))
}

func (s *TestForgetPasswordHandlerSuite) TestForget() {
	tests := []struct {
		name    string
		req     *forgetpasswordproto.ForgetRequest
		res     response.ForgetResponse
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req:  &forgetpasswordproto.ForgetRequest{Email: "test@example.com"},
			res: response.ForgetResponse{
				Code: statuscode.SUCCESSFUL,
				Msg:  messages.FORGET_PASSWORD_SEND_MAIL_IS_SUCCESSFUL,
			},
			wantErr: false,
			wantMsg: messages.FORGET_PASSWORD_SEND_MAIL_IS_SUCCESSFUL,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSVC.On("Forget", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.Forget(tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}

			s.Equal(tt.wantMsg, resp.Msg)
		})
	}
}

func (s *TestForgetPasswordHandlerSuite) TestReset() {
	tests := []struct {
		name    string
		req     *forgetpasswordproto.ResetRequest
		res     response.ResetResponse
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req:  &forgetpasswordproto.ResetRequest{Token: "token", Password: "123456", PasswordConfirmation: "123456"},
			res: response.ResetResponse{
				Code: statuscode.SUCCESSFUL,
				Msg:  messages.RESET_PASSWORD_IS_SUCCESSFUL,
			},
			wantErr: false,
			wantMsg: messages.RESET_PASSWORD_IS_SUCCESSFUL,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSVC.On("Reset", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.Reset(tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}

			s.Equal(tt.wantMsg, resp.Msg)
		})
	}
}

func TestForgetPasswordHandlerSuite_Run(t *testing.T) {
	suite.Run(t, new(TestForgetPasswordHandlerSuite))
}
