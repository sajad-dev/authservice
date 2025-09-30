package service

import (
	"github.com/sajad-dev/authservice/authentication/internal/domain/account"
	"github.com/sajad-dev/authservice/authentication/internal/domain/account/dto/gen/request"
	"github.com/sajad-dev/authservice/authentication/internal/domain/account/dto/gen/response"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/hashing"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
)

type AccountSvc struct {
	Repo    account.AccountCURDRepository
	Hashing hashing.Hashing
}

func NewAccountSvc(repo account.AccountCURDRepository, hashing hashing.Hashing) *AccountSvc {
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
			Msg:  messages.ERR_INVALID_FIELDS,
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
		Msg:  messages.SUCCESS_ACCOUNT_CREATED,
		Code: statuscode.SUCCESSFUL,
	}, nil

}

func (s *AccountSvc) Update(req request.UpdateRequest) (response.UpdateResponse, error) {
	var err error
	account, ok := models.AccountInput(req)
	if !ok {
		return response.UpdateResponse{
			Msg:  messages.ERR_INVALID_FIELDS,
			Code: statuscode.VALIDATION_ERR,
		}, nil
	}

	account.Password, err = s.Hashing.Sum([]byte(req.Password))
	if err != nil {
		return response.UpdateResponse{}, errs.Err(err)
	}

	err = s.Repo.Update(account, int(req.Id))
	if err != nil {
		return response.UpdateResponse{}, errs.Err(err)
	}

	return response.UpdateResponse{
		Msg:  messages.SUCCESS_ACCOUNT_UPDATED,
		Code: statuscode.SUCCESSFUL,
	}, nil

}

func (s *AccountSvc) Delete(req request.DeleteRequest) (response.DeleteResponse, error) {
	err := s.Repo.Delete(int(req.Id))
	if err != nil {
		return response.DeleteResponse{}, errs.Err(err)
	}

	return response.DeleteResponse{
		Msg:  messages.SUCCESS_ACCOUNT_DELETED,
		Code: statuscode.SUCCESSFUL,
	}, nil
}

func (s *AccountSvc) Read(req request.ReadRequest) (response.ReadResponse, error) {
	account, err := s.Repo.Read(int(req.Id))
	if err != nil {
		return response.ReadResponse{}, errs.Err(err)
	}

	return response.ReadResponse{
		Msg:  messages.SUCCESS_ACCOUNT_RETRIEVED,
		Code: statuscode.SUCCESSFUL,
		Data: models.AccountOutput(account),
	}, nil
}

var _ account.AccountCURDService = &AccountSvc{}
