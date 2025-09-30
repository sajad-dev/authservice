package repository_test

import (
	"testing"

	"github.com/sajad-dev/authservice/authorization/internal/domain/group/repository"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/authorize/mocks"
	"github.com/stretchr/testify/suite"
)

type GroupRepoTestSuite struct {
	suite.Suite
	repo      *repository.GroupRepo
	mockAuthz *mocks.Authorize
}

func (s *GroupRepoTestSuite) SetupTest() {
	s.mockAuthz = new(mocks.Authorize)
	s.repo = repository.NewGroupRepo(s.mockAuthz)
}

func (s *GroupRepoTestSuite) TestCreate() {
	tests := []struct {
		name    string
		sub     string
		grp     string
		mockRet bool
		mockErr error
		want    bool
		wantErr bool
	}{
		{
			name:    "Success",
			sub:     "user",
			grp:     "resource",
			mockRet: true,
			mockErr: nil,
			want:    true,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockAuthz.On("AddGroup", tt.sub, tt.grp).Return(tt.mockRet, tt.mockErr)

			resp, err := s.repo.Create(tt.sub, tt.grp)

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

func (s *GroupRepoTestSuite) TestDelete() {
	tests := []struct {
		name    string
		sub     string
		grp     string
		mockRet bool
		mockErr error
		want    bool
		wantErr bool
	}{
		{
			name:    "Success",
			sub:     "user",
			grp:     "resource",
			mockRet: true,
			mockErr: nil,
			want:    true,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			s.mockAuthz.On("RemoveGroup", tt.sub, tt.grp).Return(tt.mockRet, tt.mockErr)

			resp, err := s.repo.Delete(tt.sub, tt.grp)

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

func (s *GroupRepoTestSuite) TestGetAll() {
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

func TestGroupRepoTestSuite(t *testing.T) {
	suite.Run(t, new(GroupRepoTestSuite))
}

