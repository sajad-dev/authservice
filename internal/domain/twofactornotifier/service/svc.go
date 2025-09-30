package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier"
	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier/dto/gen/request"
	"github.com/sajad-dev/authservice/internal/domain/twofactornotifier/dto/gen/response"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/crypto"
	"github.com/sajad-dev/authservice/internal/shared/constants"
	"github.com/sajad-dev/authservice/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/internal/shared/errors/errs"
	"github.com/sajad-dev/authservice/internal/shared/helpers/codegenerate"
	"github.com/sajad-dev/authservice/internal/shared/mail"
	"github.com/sajad-dev/authservice/internal/shared/models"
)

var TOKEN_NOT_VALID = errors.New("")

type TwoFactorNotifierSvc struct {
	Repo   twofactornotifier.TwoFactorNotifierRepository
	Crypto crypto.Crypto
}

func NewTwoFactorNotifierService(repo twofactornotifier.TwoFactorNotifierRepository, cry crypto.Crypto) *TwoFactorNotifierSvc {
	return &TwoFactorNotifierSvc{
		Repo:   repo,
		Crypto: cry,
	}
}

func (a TwoFactorNotifierSvc) _verifyJWT(token string) (crypto.DataClaims, error) {
	return a.Crypto.Validate(token)
}

func (s TwoFactorNotifierSvc) _tokenValidation(token string) (int, error) {

	claims, err := s._verifyJWT(token)
	if err != nil {
		return 0, errs.Err(err)
	}

	typeToken, okType := claims["type"]
	id, okID := claims["id"]
	idInt, err := strconv.Atoi(id)
	if !okType || !okID || typeToken != "TwoFactor" || err != nil {
		return 0, errs.Err(TOKEN_NOT_VALID)
	}
	return idInt, nil
}

func (s *TwoFactorNotifierSvc) NotifierEmail(req request.NotifierEmailRequest) (response.TwoFactorNotifierResponse, error) {
	idInt, err := s._tokenValidation(req.Token)
	if err != nil {
		return response.TwoFactorNotifierResponse{}, errs.Err(err)
	}

	account, err := s.Repo.FindById(idInt)
	if err != nil {
		return response.TwoFactorNotifierResponse{}, errs.Err(err)
	}

	err = s.Repo.RemoveExpierd(int(account.ID))
	if err != nil {
		return response.TwoFactorNotifierResponse{}, errs.Err(err)
	}

	code := codegenerate.Generate()
	err = s.Repo.CreateCode(
		models.NewTwoFactorCode(
			models.WithCode(code),
			models.WithType(string(constants.EMAIL_TWO_FACTOR_CODE)),
			models.WithAccount(*account),
		),
	)
	if err != nil {
		return response.TwoFactorNotifierResponse{}, errs.Err(err)
	}

	email := mail.NewMail(
		mail.WithSendAt(time.Now()),
		mail.WithMessageHTML(fmt.Sprintf("Code is : %d", code)),
		mail.WithSendTo(account.Email),
		mail.WithTitle(messages.TWO_FACTORY_TITLE),
	)
	err = email.AddJob()
	if err != nil {
		return response.TwoFactorNotifierResponse{}, errs.Err(err)
	}

	return response.TwoFactorNotifierResponse{
		Msg: messages.SUCCESS_SEND_EMAIL_TWO_FACTOR,
		Code:    statuscode.SUCCESSFUL,
	}, nil
}

var _ twofactornotifier.TwoFactorNotifierService = &TwoFactorNotifierSvc{}
