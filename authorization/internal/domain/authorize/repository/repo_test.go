package repository_test

import (
	"testing"

	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/repository"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/authorize/mocks"
	"github.com/stretchr/testify/suite"
)

type AuthorizeRepoTestSuite struct {
	suite.Suite
	repo      *repository.AuthorizeRepo
	mockAuthz *mocks.Authorize
}

func (s *AuthorizeRepoTestSuite) SetupTest() {
	s.mockAuthz = new(mocks.Authorize)
	s.repo = repository.NewAuthorizeRepo(s.mockAuthz)
}

func (s *AuthorizeRepoTestSuite) TestCheck() {
	tests := []struct {
		name    string
		sub     string
		grp     string
		act     string
		mockRet bool
		mockErr error
		want    bool
		wantErr bool
	}{
		{
			name:    "Success",
			sub:     "user",
			grp:     "resource",
			act:     "all",
			mockRet: true,
			mockErr: nil,
			want:    true,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockAuthz.On("Verify", tt.sub, tt.grp, tt.act).Return(tt.mockRet, tt.mockErr)

			resp, err := s.repo.Check(tt.sub, tt.grp, tt.act)

			if tt.wantErr {
				s.Error(err)
			} else {
				s.NoError(err)
			}
			s.Equal(tt.want, resp)

			s.mockAuthz.AssertExpectations(s.T())
		})
	}
}

func TestAuthorizeRepoTestSuite(t *testing.T) {
	suite.Run(t, new(AuthorizeRepoTestSuite))
}
