package service

import (
	"testing"

	"github.com/sajad-dev/authservice/authorization/internal/domain/group/dto/gen/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/dto/gen/response"
	"github.com/sajad-dev/authservice/authorization/internal/domain/group/mocks"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/statuscode"
	"github.com/stretchr/testify/suite"
)

type GroupSvcTestSuite struct {
	suite.Suite
	service  *GroupSvc
	mockRepo *mocks.GroupRepository
}

func (s *GroupSvcTestSuite) SetupTest() {
	s.mockRepo = new(mocks.GroupRepository)
	s.service = NewGroupSvc(s.mockRepo)
}

func (s *GroupSvcTestSuite) TestCreate() {
	tests := []struct {
		name    string
		req     request.CreateRequest
		mockRet bool
		mockErr error
		want    response.Response
		wantErr bool
	}{
		{
			name:    "Success",
			req:     request.CreateRequest{Subject: "user", Group: "resource"},
			mockRet: true,
			mockErr: nil,
			want: response.Response{
				Msg:  messages.SUCCESS_POLICY_ADDED,
				Code: statuscode.SUCCESSFUL,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockRepo.On("Create", tt.req.Subject, tt.req.Group).Return(tt.mockRet, tt.mockErr)

			resp, err := s.service.Create(tt.req)

			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
			}
			s.Equal(tt.want, resp)

			s.mockRepo.AssertExpectations(s.T())
		})
	}
}

func (s *GroupSvcTestSuite) TestDelete() {
	tests := []struct {
		name    string
		req     request.DeleteRequest
		mockRet bool
		mockErr error
		want    response.Response
		wantErr bool
	}{
		{
			name:    "Success",
			req:     request.DeleteRequest{Subject: "user", Group: "resource"},
			mockRet: true,
			mockErr: nil,
			want: response.Response{
				Msg:  messages.SUCCESS_POLICY_ADDED,
				Code: statuscode.SUCCESSFUL,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockRepo.On("Delete", tt.req.Subject, tt.req.Group).Return(tt.mockRet, tt.mockErr)

			resp, err := s.service.Delete(tt.req)

			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
			}
			s.Equal(tt.want, resp)

			s.mockRepo.AssertExpectations(s.T())
		})
	}
}

// func (s *GroupSvcTestSuite) TestGetAll() {
// 	tests := []struct {
// 		name    string
// 		mockRet []interface{}
// 		mockErr error
// 		want    response.GetAllResponse
// 		wantErr bool
// 	}{
// 		{
// 			name:    "Success",
// 			mockRet: []interface{}{},
// 			mockErr: nil,
// 			want: response.GetAllResponse{
// 				Msg:  messages.SUCCESS_POLICY_ADDED,
// 				Code: statuscode.SUCCESSFUL,
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		s.Run(tt.name, func() {
// 			s.mockRepo.On("GetAll").Return(tt.mockRet, tt.mockErr)

// 			resp, err := s.service.GetAll()

// 			if tt.wantErr {
// 				s.Error(err)
// 			} else {
// 				s.NoError(err)
// 			}
// 			s.Equal(tt.want, resp)

// 			s.mockRepo.AssertExpectations(s.T())
// 		})
// 	}
// }

func TestGroupSvcTestSuite(t *testing.T) {
	suite.Run(t, new(GroupSvcTestSuite))
}
