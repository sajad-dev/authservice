package service

import (
	"strconv"

	"github.com/pquerna/otp/totp"
	"github.com/sajad-dev/authservice/internal/domain/twofactor"
	"github.com/sajad-dev/authservice/internal/domain/twofactor/dto/gen/request"
	"github.com/sajad-dev/authservice/internal/domain/twofactor/dto/gen/response"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/crypto"
	"github.com/sajad-dev/authservice/internal/shared/adaptor/hashing"
	"github.com/sajad-dev/authservice/internal/shared/constants"
	"github.com/sajad-dev/authservice/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/internal/shared/errors/errs"
	"github.com/sajad-dev/authservice/internal/shared/errors/errvar"
	"github.com/sajad-dev/authservice/internal/shared/helpers/timeutil"
	"github.com/sajad-dev/authservice/internal/shared/models"
)

type TwoFactorSvc struct {
	Repo    twofactor.TwoFactorRepo
	Crypto  crypto.Crypto
	Hashing hashing.Hashing
}

func NewTwoFactorService(cry crypto.Crypto, hashing hashing.Hashing, repo twofactor.TwoFactorRepo) *TwoFactorSvc {
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

func (s *TwoFactorSvc) Google(req request.GoogleRequest) (response.TwoFactorResponse, error) {
	claims, err := s._verifyJWT(req.Token)
	if err != nil {
		return response.TwoFactorResponse{}, errs.Err(err)
	}

	typeToken, okType := claims["type"]
	id, okID := claims["id"]
	idInt, err := strconv.Atoi(id)
	if !okType || !okID || typeToken != "TwoFactor" || err != nil {
		return response.TwoFactorResponse{
			Msg:  messages.TOKEN_NOT_VALID,
			Code: statuscode.VALIDATION_ERR,
		}, nil
	}

	account, err := s.Repo.FindById(idInt)
	if err != nil {
		return response.TwoFactorResponse{}, errs.Err(err)
	}

	if account.GoogleAuthScreatKey == "" {
		return response.TwoFactorResponse{}, errs.Err(errvar.SCREAT_KEY_IS_NOT_VALID)
	}

	valid := totp.Validate(strconv.Itoa(int(req.Code)), account.GoogleAuthScreatKey)
	if !valid {
		return response.TwoFactorResponse{
			Msg:  messages.GOOGLE_CODE_IS_NOT_VALID,
			Code: statuscode.VALIDATION_ERR,
		}, nil
	}

	claimsCreate := map[string]string{
		"id": strconv.Itoa(int(idInt)),
	}

	cry, err := s._createJWT(claimsCreate)
	if err != nil {
		return response.TwoFactorResponse{}, errs.Err(err)
	}

	return response.TwoFactorResponse{
		Msg:   messages.GOOGLE_TWO_FACTORT_SUCCESSFUL,
		Code:  statuscode.SUCCESSFUL,
		Token: cry,
		Data:  models.AccountOutput(account),
	}, nil
}

func (s *TwoFactorSvc) Email(req request.EmailRequest) (response.TwoFactorResponse, error) {
	table, err := s.Repo.FindByCode(int(req.Code), string(constants.EMAIL_TWO_FACTORY_CODE))
	if err != nil {
		return response.TwoFactorResponse{}, errs.Err(err)
	}

	claims := map[string]string{
		"id": strconv.Itoa(int(table[0].AccountID)),
	}

	cry, err := s._createJWT(claims)
	if err != nil {
		return response.TwoFactorResponse{}, errs.Err(err)
	}

	return response.TwoFactorResponse{
		Msg:   messages.GOOGLE_TWO_FACTORT_SUCCESSFUL,
		Code:  statuscode.SUCCESSFUL,
		Token: cry,
		Data:  models.AccountOutput(&table[0].Account),
	}, nil
}

var _ twofactor.TwoFactorService = &TwoFactorSvc{}
