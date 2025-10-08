package handler

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/dto/gen/response"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/mocks"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/policyproto"
	"github.com/sajad-dev/authservice/authorization/internal/shared/validation"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/statuscode"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type PolicyHandlerSuite struct {
	suite.Suite
	handler  *PolicyHdlr
	mocksSvc *mocks.PolicyService
}

func (s *PolicyHandlerSuite) SetupSuite() {
	mocksSvc := new(mocks.PolicyService)
	s.mocksSvc = mocksSvc

	s.handler = NewPolicyHdlr(mocksSvc, validate.NewValidate(validator.New()))
}

func (s *PolicyHandlerSuite) TestCreate() {
	tests := []struct {
		name    string
		req     *policyproto.CreateRequest
		res     response.Response
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req:  &policyproto.CreateRequest{Subject: "user", Object: "resource", Action: "create"},
			res: response.Response{
				Code: statuscode.SUCCESSFUL,
				Msg:  "Policy created successfully",
			},
			wantErr: false,
			wantMsg: "Policy created successfully",
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

func (s *PolicyHandlerSuite) TestDelete() {
	tests := []struct {
		name    string
		req     *policyproto.DeleteRequest
		res     response.Response
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req:  &policyproto.DeleteRequest{Subject: "user", Object: "resource", Action: "delete"},
			res: response.Response{
				Code: statuscode.SUCCESSFUL,
				Msg:  "Policy deleted successfully",
			},
			wantErr: false,
			wantMsg: "Policy deleted successfully",
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

func (s *PolicyHandlerSuite) TestGetAll() {
	tests := []struct {
		name    string
		req     *policyproto.GetAllRequest
		res     response.GetAllResponse
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req:  &policyproto.GetAllRequest{},
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

func TestPolicyHandlerSuite_Run(t *testing.T) {
	suite.Run(t, new(PolicyHandlerSuite))
}
