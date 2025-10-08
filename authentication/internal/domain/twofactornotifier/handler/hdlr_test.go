package handler_test

import (
	"context"
	"testing"

	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/dto/gen/response"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/handler"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/mocks"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/twofactornotifierproto"
	mockdb "github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb/mocks"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/authentication/internal/shared/validation"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestTwoFactorNotifierHandlerSuite struct {
	suite.Suite
	handler  twofactornotifier.TwoFactorNotifierHandler
	mocksSVC *mocks.TwoFactorNotifierService
}

func (s *TestTwoFactorNotifierHandlerSuite) SetupSuite() {
	mocksSVC := new(mocks.TwoFactorNotifierService)
	s.mocksSVC = mocksSVC

	mockDB := &mockdb.SqlDBGlobal{}

	s.handler = handler.NewTwoFactorNotifierHdlr(mocksSVC, validation.NewValidator(mockDB))
}

func (s *TestTwoFactorNotifierHandlerSuite) TestNotifierEmail() {
	tests := []struct {
		name    string
		req     *twofactornotifierproto.NotifierEmailRequest
		res     response.TwoFactorNotifierResponse
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req:  &twofactornotifierproto.NotifierEmailRequest{Token: "111"},
			res: response.TwoFactorNotifierResponse{
				Code: statuscode.SUCCESSFUL,
				Msg:  messages.SUCCESS_SEND_EMAIL_TWO_FACTOR,
			},
			wantErr: false,
			wantMsg: messages.SUCCESS_SEND_EMAIL_TWO_FACTOR,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSVC.On("NotifierEmail", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.NotifierEmail(context.Background(),tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}

			s.Equal(tt.wantMsg, resp.Msg)
		})
	}
}

func TestTwoFactorNotifierHandlerSuite_Run(t *testing.T) {
	suite.Run(t, new(TestTwoFactorNotifierHandlerSuite))
}
