package service_test

import (
	"testing"
	"time"

	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/dto/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/dto/response"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/mocks"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/service"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/crypto"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/crypto/hs256"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/statuscode"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthorizeSvcTestSuite struct {
	suite.Suite
	service  *service.AuthorizeSvc
	mockRepo *mocks.AuthorizeRepository
	crp      crypto.Crypto
}

func (s *AuthorizeSvcTestSuite) SetupTest() {
	s.crp = hs256.NewJWT([]byte("(:"))
	s.mockRepo = new(mocks.AuthorizeRepository)
	s.service = service.NewAuthorizeSvc(s.mockRepo, s.crp, []string{"/"})
}

func (s *AuthorizeSvcTestSuite) TestCreate() {
	token, _ := s.crp.Generate(map[string]string{}, time.Now().Add(time.Hour))
	tests := []struct {
		name    string
		req     request.AuthorizeRequest
		mockRet bool
		mockErr error
		want    response.AuthorizeResponse
		wantErr bool
	}{
		{
			name: "Success",
			req: request.AuthorizeRequest{Headers: map[string]string{
				"authorization": token,
			},
				Method: "GET",
				Path:   "/",
			},
			mockRet: true,
			mockErr: nil,
			want: response.AuthorizeResponse{
				Msg:  messages.SUCCESS_VERIFY,
				Code: statuscode.SUCCESSFUL,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockRepo.On("Check", mock.Anything, mock.Anything, mock.Anything).Return(tt.mockRet, tt.mockErr)

			resp, err := s.service.Check(tt.req)

			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
			}
			s.Equal(tt.want, resp)

		})
	}
}

func TestAuthorizeSvcTestSuite(t *testing.T) {
	suite.Run(t, new(AuthorizeSvcTestSuite))
}
