package handler_test

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/sajad-dev/authservice/internal/domain/account"
	"github.com/sajad-dev/authservice/internal/domain/account/accountproto"
	"github.com/sajad-dev/authservice/internal/domain/account/dto/gen/response"
	"github.com/sajad-dev/authservice/internal/domain/account/handler"
	"github.com/sajad-dev/authservice/internal/domain/account/mocks"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/validation/validate"
	"github.com/sajad-dev/authservice/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/internal/shared/models"
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

	s.handler = handler.NewAccountHandler(mocksSVC, validate.NewValidate(validator.New()))
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
			req:  &accountproto.CreateRequest{Username: "test", Email: "test@example.com"},
			res: response.CreateResponse{
				Code: statuscode.SUCCESSFUL,
				Msg:  messages.CREATE_ACCOUNT_SUCCESSFUL,
			},
			wantErr: false,
			wantMsg: messages.CREATE_ACCOUNT_SUCCESSFUL,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSVC.On("Create", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.Create(tt.req)
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
			req:  &accountproto.UpdateRequest{Username: "test", Email: "update@example.com"},
			res: response.UpdateResponse{
				Code: statuscode.SUCCESSFUL,
				Msg:  messages.UPDATE_ACCOUNT_SUCCESSFUL,
			},
			wantErr: false,
			wantMsg: messages.UPDATE_ACCOUNT_SUCCESSFUL,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSVC.On("Update", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.Update(tt.req)
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
			req:  &accountproto.DeleteRequest{Id: 0},
			res: response.DeleteResponse{
				Code: statuscode.SUCCESSFUL,
				Msg:  messages.DELETE_ACCOUNT_SUCCESSFUL,
			},
			wantErr: false,
			wantMsg: messages.DELETE_ACCOUNT_SUCCESSFUL,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSVC.On("Delete", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.Delete(tt.req)
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
			req:  &accountproto.ReadRequest{Id: 0},
			res: response.ReadResponse{
				Code: statuscode.SUCCESSFUL,
				Msg:  messages.READ_ACCOUNT_SUCCESSFUL,
				Data: models.AccountFiltered{
					Username: "test",
					Email:    "test@example.com",
				},
			},
			wantErr: false,
			wantMsg: messages.READ_ACCOUNT_SUCCESSFUL,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSVC.On("Read", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.Read(tt.req)
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
