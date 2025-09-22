package service

import (
	"strconv"
	"time"

	"github.com/sajad-dev/authservice/internal/domain/forgetpassword"
	"github.com/sajad-dev/authservice/internal/domain/forgetpassword/dto/gen/request"
	"github.com/sajad-dev/authservice/internal/domain/forgetpassword/dto/gen/response"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/crypto"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/hashing"
	"github.com/sajad-dev/authservice/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/internal/shared/errors/errs"
	"github.com/sajad-dev/authservice/internal/shared/helpers/timeutil"
	"github.com/sajad-dev/authservice/internal/shared/mail"
)

type TwoFactorSvc struct {
	Repo    forgetpassword.TwoFactorRepo
	Crypto  crypto.Crypto
	Hashing hashing.Hashing
}

func NewTwoFactorService(cry crypto.Crypto, hashing hashing.Hashing, repo forgetpassword.TwoFactorRepo) *TwoFactorSvc {
	return &TwoFactorSvc{
		Repo:    repo,
		Crypto:  cry,
		Hashing: hashing,
	}
}

func (a TwoFactorSvc) _createJWT(claims map[string]string) (string, error) {
	return a.Crypto.Generate(claims, timeutil.TokenExpire())
}

func (a TwoFactorSvc) _verifyJWT(token string) (crypto.DataClaims, error) {
	return a.Crypto.Validate(token)
}

func (a TwoFactorSvc) Forget(req request.ForgetRequest) (response.ForgetResponse, error) {

	table, err := a.Repo.Find("email", req.Email)
	if err != nil {
		return response.ForgetResponse{}, errs.Err(err)
	}

	claims := map[string]string{
		"type": "TwoFactor",
		"id":   strconv.Itoa(int(table.ID)),
	}

	token, err := a._createJWT(claims)
	if err != nil {
		return response.ForgetResponse{}, errs.Err(err)
	}

	mail := mail.NewMail(
		mail.WithSendAt(time.Now()),
		mail.WithMessageHTML(token),
		mail.WithSendTo(table.Email),
		mail.WithTitle(messages.FORGET_PASSWORD_TITLE),
	)

	err = mail.AddJob()
	if err != nil {
		return response.ForgetResponse{}, errs.Err(err)
	}

	return response.ForgetResponse{
		Msg:   messages.FORGET_PASSWORD_SEND_MAIL_IS_SUCCESSFUL,
		Code:  statuscode.SUCCESSFUL,
		Token: token,
	}, nil
}
func (a TwoFactorSvc) Reset(req request.ResetRequest) (response.ResetResponse, error) {
	claims, err := a._verifyJWT(req.Token)
	if err != nil {
		return response.ResetResponse{}, errs.Err(err)
	}

	typeToken, okType := claims["type"]
	id, okID := claims["id"]
	idInt, err := strconv.Atoi(id)
	if !okType || !okID || typeToken != "TwoFactor" || err != nil {
		return response.ResetResponse{
			Msg:  messages.TOKEN_NOT_VALID,
			Code: statuscode.VALIDATION_ERR,
		}, nil
	}

	row, err := a.Repo.FindById(idInt)
	if err != nil {
		return response.ResetResponse{}, errs.Err(err)
	}

	hashPassword, err := a.Hashing.Sum([]byte(req.Password))
	if err != nil {
		return response.ResetResponse{}, errs.Err(err)
	}

	row.Password = hashPassword

	err = a.Repo.Update(row)
	if err != nil {
		return response.ResetResponse{}, errs.Err(err)
	}

	return response.ResetResponse{
			Code: statuscode.SUCCESSFUL,
			Msg:  messages.RESET_PASSWORD_IS_SUCCESSFUL,
		},
		nil
}

var _ forgetpassword.TwoFactorService = &TwoFactorSvc{}
