package handler_test

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/dto/gen/response"
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/groupproto"
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/handler"
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/mocks"
	"github.com/sajad-dev/authservice/authorization/internal/shared/validation/validate"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/statuscode"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type GroupHandlerSuite struct {
	suite.Suite
	handler  *handler.GroupHdlr
	mocksSvc *mocks.GroupService
}

func (s *GroupHandlerSuite) SetupSuite() {
	mocksSvc := new(mocks.GroupService)
	s.mocksSvc = mocksSvc

	s.handler = handler.NewGroupHdlr(mocksSvc, validate.NewValidate(validator.New()))
}

func (s *GroupHandlerSuite) TestCreate() {
	tests := []struct {
		name    string
		req     *groupproto.CreateRequest
		res     response.Response
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req:  &groupproto.CreateRequest{Subject: "user", Group: "resource"},
			res: response.Response{
				Code: statuscode.SUCCESSFUL,
				Msg:  "Group created successfully",
			},
			wantErr: false,
			wantMsg: "Group created successfully",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSvc.On("Create", mock.Anything).Return(tt.res, nil)

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

func (s *GroupHandlerSuite) TestDelete() {
	tests := []struct {
		name    string
		req     *groupproto.DeleteRequest
		res     response.Response
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req:  &groupproto.DeleteRequest{Subject: "user", Group: "resource"},
			res: response.Response{
				Code: statuscode.SUCCESSFUL,
				Msg:  "Group deleted successfully",
			},
			wantErr: false,
			wantMsg: "Group deleted successfully",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSvc.On("Delete", mock.Anything).Return(tt.res, nil)

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

func (s *GroupHandlerSuite) TestGetAll() {
	tests := []struct {
		name    string
		req     *groupproto.GetAllRequest
		res     response.GetAllResponse
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req:  &groupproto.GetAllRequest{},
			res: response.GetAllResponse{
				Code: statuscode.SUCCESSFUL,
				Msg:  "Fetched all policies successfully",
				// Data: [][]string{[]string{}},
			},
			wantErr: false,
			wantMsg: "Fetched all policies successfully",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSvc.On("GetAll", mock.Anything).Return(tt.res, nil)

			resp, err := s.handler.GetAll(tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}

			s.Equal(tt.wantMsg, resp.Msg)
		})
	}
}

func TestGroupHandlerSuite_Run(t *testing.T) {
	suite.Run(t, new(GroupHandlerSuite))
}

