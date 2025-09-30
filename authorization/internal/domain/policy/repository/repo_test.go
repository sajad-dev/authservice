package repository_test

import (
	"testing"

	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/repository"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/authorize/mocks"
	"github.com/stretchr/testify/suite"
)

type PolicyRepoTestSuite struct {
	suite.Suite
	repo      *repository.PolicyRepo
	mockAuthz *mocks.Authorize
}

func (s *PolicyRepoTestSuite) SetupTest() {
	s.mockAuthz = new(mocks.Authorize)
	s.repo = repository.NewPolicyRepo(s.mockAuthz)
}

func (s *PolicyRepoTestSuite) TestCreate() {
	tests := []struct {
		name    string
		sub     string
		obj     string
		act     string
		mockRet bool
		mockErr error
		want    bool
		wantErr bool
	}{
		{
			name:    "Success",
			sub:     "user",
			obj:     "resource",
			act:     "create",
			mockRet: true,
			mockErr: nil,
			want:    true,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockAuthz.On("AddPolicy", tt.sub, tt.obj, tt.act).Return(tt.mockRet, tt.mockErr)

			resp, err := s.repo.Create(tt.sub, tt.obj, tt.act)

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

func (s *PolicyRepoTestSuite) TestDelete() {
	tests := []struct {
		name    string
		sub     string
		obj     string
		act     string
		mockRet bool
		mockErr error
		want    bool
		wantErr bool
	}{
		{
			name:    "Success",
			sub:     "user",
			obj:     "resource",
			act:     "delete",
			mockRet: true,
			mockErr: nil,
			want:    true,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockAuthz.On("RemovePolicy", tt.sub, tt.obj, tt.act).Return(tt.mockRet, tt.mockErr)

			resp, err := s.repo.Delete(tt.sub, tt.obj, tt.act)

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

func (s *PolicyRepoTestSuite) TestGetAll() {
	tests := []struct {
		name    string
		mockRet [][]string
		mockErr error
		want    [][]string
		wantErr bool
	}{
		{
			name:    "Success",
			mockRet: [][]string{{"user", "resource", "create"}},
			mockErr: nil,
			want:    [][]string{{"user", "resource", "create"}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockAuthz.On("GetAllGroup").Return(tt.mockRet, tt.mockErr)

			resp, err := s.repo.GetAll()

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

func TestPolicyRepoTestSuite(t *testing.T) {
	suite.Run(t, new(PolicyRepoTestSuite))
}
