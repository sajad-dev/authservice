package handler_test

import (
	"context"
	"testing"

	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	rq "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/dto/response"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/handler"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/mocks"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/statuscode"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/genproto/googleapis/rpc/code"
)

type AuthorizeHandlerSuite struct {
	suite.Suite
	handler  *handler.AuthorizeHdlr
	mocksSvc *mocks.AuthorizeService
}

func (s *AuthorizeHandlerSuite) SetupSuite() {
	s.mocksSvc = new(mocks.AuthorizeService)
	s.handler = handler.NewAuthorizeHdlr(s.mocksSvc)
}

func (s *AuthorizeHandlerSuite) TestCheck() {
	tests := []struct {
		name    string
		resp    response.AuthorizeResponse
		wantErr bool
		wantOK  bool
	}{
		{
			name: "Success",
			resp: response.AuthorizeResponse{
				Code: statuscode.SUCCESSFUL,
			},
			wantErr: false,
			wantOK:  true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mocksSvc.On("Check", mock.AnythingOfType("request.AuthorizeRequest")).Return(tt.resp, nil).Once()

			req := &authz.CheckRequest{
				Attributes: &rq.AttributeContext{
					Request: &rq.AttributeContext_Request{
						Http: &rq.AttributeContext_HttpRequest{
							Method:   "GET",
							Path:     "/test",
							Host:     "example.com",
							Body:     "body",
							Protocol: "HTTP/1.1",
							Headers: map[string]string{
								"Authorization": "Bearer token",
							},
						},
					},
					Source: &rq.AttributeContext_Peer{
						Address: &core.Address{
							Address: &core.Address_SocketAddress{
								SocketAddress: &core.SocketAddress{
									Address: "127.0.0.1",
									PortSpecifier: &core.SocketAddress_PortValue{
										PortValue: 8080,
									},
								},
							},
						},
					},
				},
			}

			resp, err := s.handler.Check(context.Background(), req)

			if tt.wantErr {
				s.Error(err)
				return
			}
			s.NoError(err)
			s.NotNil(resp)

			if tt.wantOK {
				s.Equal(int32(code.Code_OK), resp.Status.Code)
			} else {
				s.Equal(int32(code.Code_PERMISSION_DENIED), resp.Status.Code)
			}
		})
	}
}

func TestAuthorizeHandlerSuite_Run(t *testing.T) {
	suite.Run(t, new(AuthorizeHandlerSuite))
}
