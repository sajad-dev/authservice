package service_test

import (
	"testing"

	"github.com/sajad-dev/authservice/authorization/internal/domain/account"
	"github.com/sajad-dev/authservice/authorization/internal/domain/account/dto/gen/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/account/mocks"
	"github.com/sajad-dev/authservice/authorization/internal/domain/account/service"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/crypto"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/hashing"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/hashing/sha256"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/authorization/internal/shared/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestAccountSuite struct {
	suite.Suite
	accountService account.AccountCURDService
	repoMock       *mocks.AccountCURDRepository
	crypto         crypto.Crypto
	hash           hashing.Hashing
}

func (s *TestAccountSuite) SetupSuite() {
	s.repoMock = new(mocks.AccountCURDRepository)

	hash, _ := s.hash.Sum([]byte("pass"))
	account := models.NewAccounts(
		models.WithEmail("test@email.com"),
		models.WithUsername("test"),
		models.WithPassword(hash),
	)

	s.repoMock.On("Create", mock.Anything).Return(nil)
	s.repoMock.On("Update", mock.Anything,mock.Anything).Return(nil)
	s.repoMock.On("Read", 1).Return(account, nil)
	s.repoMock.On("Delete", 1).Return(nil)

	s.accountService = service.NewAccountSvc(
		s.repoMock,
		s.hash,
	)
}

func (s *TestAccountSuite) TestCreate() {
	tests := []struct {
		name    string
		req     request.CreateRequest
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req: request.CreateRequest{
				Email:    "test@email.com",
				Username: "test",
				Password: "pass",
			},
			wantErr: false,
			wantMsg: messages.CREATE_ACCOUNT_SUCCESSFUL,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			resp, err := s.accountService.Create(tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}
			s.Equal(tt.wantMsg, resp.Msg)
			s.Equal(int32(statuscode.SUCCESSFUL), resp.Code)
		})
	}
}

func (s *TestAccountSuite) TestUpdate() {
	tests := []struct {
		name    string
		req     request.UpdateRequest
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req: request.UpdateRequest{
				Id:       1,
				Email:    "testupdate@email.com",
				Username: "testupdate",
				Password: "pass",
			},
			wantErr: false,
			wantMsg: messages.UPDATE_ACCOUNT_SUCCESSFUL,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			resp, err := s.accountService.Update(tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}
			s.Equal(tt.wantMsg, resp.Msg)
			s.Equal(int32(statuscode.SUCCESSFUL), resp.Code)
		})
	}
}

func (s *TestAccountSuite) TestRead() {
	tests := []struct {
		name    string
		req     request.ReadRequest
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req: request.ReadRequest{
				Id: 1,
			},
			wantErr: false,
			wantMsg: messages.READ_ACCOUNT_SUCCESSFUL,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			resp, err := s.accountService.Read(tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}
			s.Equal(tt.wantMsg, resp.Msg)
			s.Equal("test@email.com", resp.Data.Email)
			s.Equal("test", resp.Data.Username)
		})
	}
}

func (s *TestAccountSuite) TestDelete() {
	tests := []struct {
		name    string
		req     request.DeleteRequest
		wantErr bool
		wantMsg string
	}{
		{
			name: "Success",
			req: request.DeleteRequest{
				Id: 1,
			},
			wantErr: false,
			wantMsg: messages.DELETE_ACCOUNT_SUCCESSFUL,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			resp, err := s.accountService.Delete(tt.req)
			if !tt.wantErr {
				s.NoError(err)
			} else {
				s.Error(err)
			}
			s.Equal(tt.wantMsg, resp.Msg)
			s.Equal(int32(statuscode.SUCCESSFUL), resp.Code)
		})
	}
}

func TestAccountSuite_Run(t *testing.T) {
	su := &TestAccountSuite{
		hash: sha256.NewSha256(),
	}
	suite.Run(t, su)
}
