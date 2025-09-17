package service

import (
	"github.com/sajad-dev/authservice/internal/domain/account"

)

type AccountService struct {
	Repo account.AccountCURDRepositories
}

func NewAccountService(repo account.AccountCURDRepositories) *AccountService {
	return &AccountService{Repo: repo}
}

func (s *AccountService) CreateService(req createreq.CreateRequest) (createres.CreateResponse, error) {
	var err error
	var account = models.NewAccounts(
		models.WithEmail(req.Email),
		models.WithFirstName(req.FirstName),
		models.WithLastName(req.LastName),
		models.WithSMS(req.SMS),
		models.WithUsername(req.Username),
		models.WithTwoFactor(req.TwoFactor),
	)

	account.Password, err = crypto.SumSHA256([]byte(req.Password))
	if err != nil {
		return *createres.NewCreateResponse(), errs.Err(err)
	}

	err = s.Repo.Create(account)
	if err != nil {
		return *createres.NewCreateResponse(), errs.Err(err)
	}

	return *createres.NewCreateResponse(
		createres.WithMessage(messages.CREATE_ACCOUNT_SUCCESSFUL),
		createres.WithCode(200),
	), nil

}

func (s *AccountService) UpdateService(req updatereq.UpdateRequest) (updateres.UpdateResponse, error) {
	var err error
	var account = models.NewAccounts(
		models.WithEmail(req.Email),
		models.WithFirstName(req.FirstName),
		models.WithLastName(req.LastName),
		models.WithSMS(req.SMS),
		models.WithUsername(req.Username),
		models.WithTwoFactor(req.TwoFactor),
	)

	account.Password, err = crypto.SumSHA256([]byte(req.Password))
	if err != nil {
		return *updateres.NewUpdateResponse(), errs.Err(err)
	}

	err = s.Repo.Update(account)
	if err != nil {
		return *updateres.NewUpdateResponse(), errs.Err(err)
	}

	return *updateres.NewUpdateResponse(
		updateres.WithMessage(messages.UPDATE_ACCOUNT_SUCCESSFUL),
		updateres.WithCode(200),
	), nil

}

func (s *AccountService) DeleteService(req deletereq.DeleteRequest) (deleteres.DeleteResponse, error) {
	err := s.Repo.Delete(req.ID)
	if err != nil {
		return *deleteres.NewDeleteResponse(), errs.Err(err)
	}

	return *deleteres.NewDeleteResponse(
		deleteres.WithMessage(messages.DELETE_ACCOUNT_SUCCESSFUL),
		deleteres.WithCode(200),
	), nil
}

func (s *AccountService) ReadService(req readreq.ReadRequest) (readres.ReadResponse, error) {
	account, err := s.Repo.Read(req.ID)
	if err != nil {
		return *readres.NewReadResponse(), errs.Err(err)
	}
	var res = readres.NewReadResponse()
	if err := adaptor.Adaptor(res, account); err != nil {
		return *readres.NewReadResponse(), errs.Err(err)
	}
	return *res, nil
}

func (s *AccountService) GetAllService() (getallres.GetAllResponse, error) {
	accounts, err := s.Repo.GetAll()
	if err != nil {
		return *getallres.NewGetAllResponse(), errs.Err(err)
	}
	var res = getallres.NewGetAllResponse()
	if err := adaptor.Adaptor(res, accounts); err != nil {
		return *getallres.NewGetAllResponse(), errs.Err(err)
	}
	return *res, nil
}

var _ account.AccountCURDService = &AccountService{}

