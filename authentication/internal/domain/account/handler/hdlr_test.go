package handler_test

import (
	"testing"

	"github.com/sajad-dev/authservice/authentication/internal/domain/account"
	"github.com/sajad-dev/authservice/authentication/internal/domain/account/accountproto"
	"github.com/sajad-dev/authservice/authentication/internal/domain/account/dto/gen/response"
	"github.com/sajad-dev/authservice/authentication/internal/domain/account/handler"
	"github.com/sajad-dev/authservice/authentication/internal/domain/account/mocks"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"

	mockdb "github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb/mocks"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/authentication/internal/shared/validation"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestAccountCURDHandlerSuite struct {
	suite.Suite
	handler  account.AccountCURDHandler
	mocksSVC *mocks.AccountCURDService
}

func (s *TestAccountCURDHandlerSuite) SetupSuite() {
	mocksSVC := new(mocks.AccountCURDService)
	s.mocksSVC = mocksSVC

	mockDB := &mockdb.SqlDBGlobal{}

	mockDB.On("Exists", mock.Anything, "Id", mock.Anything).Return(true)
	mockDB.On("Exists", mock.Anything, mock.Anything, mock.Anything).Return(false)

	s.handler = handler.NewAccountHdlr(mocksSVC, validation.NewValidator(mockDB))
}

func (s *TestAccountCURDHandlerSuite) TestCreate() {
	tests := []struct {
		name    string
		req     *accountproto.CreateRequest
		res     response.CreateResponse
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req: &accountproto.CreateRequest{
				Username:             "test1",
				Email:                "test@example.com",
				FirstName:            "Sajad",
				LastName:             "Poorajam",
				TwoFactor:            []string{},
				Password:             "MyStrongPassword123!",
				PasswordConfirmation: "MyStrongPassword123!",
			},
			res: response.CreateResponse{
				Code: statuscode.SUCCESSFUL,
				Msg:  messages.SUCCESS_ACCOUNT_CREATED,
			},
			wantErr: false,
			wantMsg: messages.SUCCESS_ACCOUNT_CREATED,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSVC.On("Create", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.Create(nil, tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}

			s.Equal(tt.wantMsg, resp.Msg)
		})
	}
}

func (s *TestAccountCURDHandlerSuite) TestUpdate() {
	tests := []struct {
		name    string
		req     *accountproto.UpdateRequest
		res     response.UpdateResponse
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req: &accountproto.UpdateRequest{
				Id:                   1,
				Username:             "test1",
				Email:                "test@example.com",
				FirstName:            "Sajad",
				LastName:             "Poorajam",
				TwoFactor:            []string{},
				Password:             "MyStrongPassword123!",
				PasswordConfirmation: "MyStrongPassword123!"},
			res: response.UpdateResponse{
				Code: statuscode.SUCCESSFUL,
				Msg:  messages.SUCCESS_ACCOUNT_UPDATED,
			},
			wantErr: false,
			wantMsg: messages.SUCCESS_ACCOUNT_UPDATED,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSVC.On("Update", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.Update(nil, tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}

			s.Equal(tt.wantMsg, resp.Msg)
		})
	}
}

func (s *TestAccountCURDHandlerSuite) TestDelete() {
	tests := []struct {
		name    string
		req     *accountproto.DeleteRequest
		res     response.DeleteResponse
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req:  &accountproto.DeleteRequest{Id: 1},
			res: response.DeleteResponse{
				Code: statuscode.SUCCESSFUL,
				Msg:  messages.SUCCESS_ACCOUNT_DELETED,
			},
			wantErr: false,
			wantMsg: messages.SUCCESS_ACCOUNT_DELETED,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSVC.On("Delete", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.Delete(nil, tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}

			s.Equal(tt.wantMsg, resp.Msg)
		})
	}
}

func (s *TestAccountCURDHandlerSuite) TestRead() {
	tests := []struct {
		name    string
		req     *accountproto.ReadRequest
		res     response.ReadResponse
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req:  &accountproto.ReadRequest{Id: 1},
			res: response.ReadResponse{
				Code: statuscode.SUCCESSFUL,
				Msg:  messages.SUCCESS_ACCOUNT_RETRIEVED,
				Data: models.AccountFiltered{
					Username: "test",
					Email:    "test@example.com",
				},
			},
			wantErr: false,
			wantMsg: messages.SUCCESS_ACCOUNT_RETRIEVED,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSVC.On("Read", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.Read(nil, tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}

			s.Equal(tt.wantMsg, resp.Msg)
		})
	}
}

func TestAccountCURDHandlerSuite_Run(t *testing.T) {
	suite.Run(t, new(TestAccountCURDHandlerSuite))
}
