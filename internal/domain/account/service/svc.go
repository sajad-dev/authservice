package service

import (
	"github.com/sajad-dev/authservice/internal/domain/account"
	"github.com/sajad-dev/authservice/internal/domain/account/dto/gen/request"
	"github.com/sajad-dev/authservice/internal/domain/account/dto/gen/response"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/hashing"
	"github.com/sajad-dev/authservice/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/internal/shared/errors/errs"
	"github.com/sajad-dev/authservice/internal/shared/models"
)

type AccountSvc struct {
	Repo    account.AccountCURDRepositories
	Hashing hashing.Hashing
}

func NewAccountSvc(repo account.AccountCURDRepositories, hashing hashing.Hashing) *AccountSvc {
	return &AccountSvc{
		Repo:    repo,
		Hashing: hashing,
	}
}

func (s *AccountSvc) Create(req request.CreateRequest) (response.CreateResponse, error) {

	var err error
	account, ok := models.AccountInput(req)
	if !ok {
		return response.CreateResponse{
			Msg:  messages.NOT_VALID_FIELDS_ERR,
			Code: statuscode.VALIDATION_ERR,
		}, nil
	}

	account.Password, err = s.Hashing.Sum([]byte(req.Password))
	if err != nil {
		return response.CreateResponse{}, errs.Err(err)
	}

	err = s.Repo.Create(account)
	if err != nil {
		return response.CreateResponse{}, errs.Err(err)
	}

	return response.CreateResponse{
		Msg:  messages.CREATE_ACCOUNT_SUCCESSFUL,
		Code: statuscode.SUCCESSFUL,
	}, nil

}

func (s *AccountSvc) UpdateService(req request.UpdateRequest) (response.UpdateResponse, error) {
	var err error
	account, ok := models.AccountInput(req)
	if !ok {
		return response.UpdateResponse{
			Msg:  messages.NOT_VALID_FIELDS_ERR,
			Code: statuscode.VALIDATION_ERR,
		}, nil
	}

	account.Password, err = s.Hashing.Sum([]byte(req.Password))
	if err != nil {
		return response.UpdateResponse{}, errs.Err(err)
	}

	err = s.Repo.Update(account)
	if err != nil {
		return response.UpdateResponse{}, errs.Err(err)
	}

	return response.UpdateResponse{
		Msg:  messages.UPDATE_ACCOUNT_SUCCESSFUL,
		Code: statuscode.SUCCESSFUL,
	}, nil

}

func (s *AccountSvc) DeleteService(req request.DeleteRequest) (response.DeleteResponse, error) {
	err := s.Repo.Delete(int(req.Id))
	if err != nil {
		return response.DeleteResponse{}, errs.Err(err)
	}

	return response.DeleteResponse{
		Msg:  messages.DELETE_ACCOUNT_SUCCESSFUL,
		Code: statuscode.SUCCESSFUL,
	}, nil
}

func (s *AccountSvc) ReadService(req request.ReadRequest) (response.ReadResponse, error) {
	account, err := s.Repo.Read(int(req.Id))
	if err != nil {
		return response.ReadResponse{}, errs.Err(err)
	}

	return response.ReadResponse{
		Msg:  messages.READ_ACCOUNT_SUCCESSFUL,
		Code: statuscode.SUCCESSFUL,
		Data: models.AccountOutput(account),
	}, nil
}
