package service

import (
	"testing"

	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/dto/gen/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/dto/gen/response"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/mocks"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/statuscode"
	"github.com/stretchr/testify/suite"
)

type PolicySvcTestSuite struct {
	suite.Suite
	service  *PolicySvc
	mockRepo *mocks.PolicyRepository
}

func (s *PolicySvcTestSuite) SetupTest() {
	s.mockRepo = new(mocks.PolicyRepository)
	s.service = NewPolicySvc(s.mockRepo)
}

func (s *PolicySvcTestSuite) TestCreate() {
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
			req:     request.CreateRequest{Subject: "user", Object: "resource", Action: "create"},
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
			s.mockRepo.On("Create", tt.req.Subject, tt.req.Object, tt.req.Action).Return(tt.mockRet, tt.mockErr)

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

func (s *PolicySvcTestSuite) TestDelete() {
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
			req:     request.DeleteRequest{Subject: "user", Object: "resource", Action: "delete"},
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
			s.mockRepo.On("Delete", tt.req.Subject, tt.req.Object, tt.req.Action).Return(tt.mockRet, tt.mockErr)

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

// func (s *PolicySvcTestSuite) TestGetAll() {
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

func TestPolicySvcTestSuite(t *testing.T) {
	suite.Run(t, new(PolicySvcTestSuite))
}
