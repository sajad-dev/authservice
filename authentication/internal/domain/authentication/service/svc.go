package service

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"

	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication"
	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/dto/gen/request"
	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/dto/gen/response"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/crypto"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/hashing"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authentication/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs"
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errvar"
	"github.com/sajad-dev/authservice/authentication/internal/shared/helpers/timeutil"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
)

const (
	EMAIL    = "email"
	USERNAME = "username"
)

const (
	EMAILREGEX    = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	PHONEREGEX    = `^\+[1-9]\d{1,14}$`
	USERNAMEREGEX = `^[a-zA-Z0-9\-]{3,20}$`
)

type AuthenticationSvc struct {
	Repo    authentication.AuthenticatorRepository
	crypto  crypto.Crypto
	Hashing hashing.Hashing
}

func NewAuthenticationSvc(repo authentication.AuthenticatorRepository, cry crypto.Crypto, hashing hashing.Hashing) *AuthenticationSvc {
	return &AuthenticationSvc{
		Repo:    repo,
		crypto:  cry,
		Hashing: hashing,
	}
}

func _checkUsernameOrEmail(field string) (string, error) {
	email, err := regexp.MatchString(EMAILREGEX, field)
	if err != nil {
		return "", errs.Err(err)
	}

	if email {
		return EMAIL, nil
	}

	username, err := regexp.MatchString(USERNAMEREGEX, field)
	if err != nil {
		return "", errs.Err(err)
	}
	if username {
		return USERNAME, nil
	}
	return "", errvar.USERNAME_NOT_VALID
}

func (a *AuthenticationSvc) _createJWT(claims map[string]string) (string, error) {
	token, err := a.crypto.Generate(claims, timeutil.TokenExpire())
	return token, errs.Err(err)
}

func (a *AuthenticationSvc) Login(req request.LoginRequest) (response.LoginResponse, error) {
	hashPassword, err := a.Hashing.Sum([]byte(req.Password))
	if err != nil {
		return response.LoginResponse{}, errs.Err(err)
	}
	usernameType, err := _checkUsernameOrEmail(req.Username)
	if err != nil {
		return response.LoginResponse{}, errs.Err(err)
	}
	table, err := a.Repo.Find(usernameType, req.Username)
	if err != nil {
		return response.LoginResponse{}, errs.Err(err)
	}

	if table.Password != hashPassword {
		return response.LoginResponse{
			Msg:  messages.ERR_INVALID_CREDENTIALS,
			Code: statuscode.VALIDATION_ERR,
		}, nil
	}

	claims := map[string]string{}
	msg := ""

	if len(table.TwoFactor) > 0 {
		jsonTwoFactory, err := json.Marshal(table.TwoFactor)
		if err != nil {
			return response.LoginResponse{}, errs.Err(err)
		}
		msg = messages.SUCCESS_LOGIN_TWO_FACTOR
		claims = map[string]string{
			"type":    "TwoFactor",
			"options": string(jsonTwoFactory),
		}
	} else {
		msg = messages.SUCCESS_LOGIN
		claims = map[string]string{
			"user": strconv.Itoa(int(table.ID)),
		}
	}

	crp, err := a._createJWT(claims)
	if err != nil {
		return response.LoginResponse{}, errs.Err(err)
	}

	return response.LoginResponse{
		Data:  models.AccountOutput(table),
		Msg:   msg,
		Token: crp,
		Code:  statuscode.SUCCESSFUL,
	}, nil

}
func (a *AuthenticationSvc) Register(req request.RegisterRequest) (response.RegisterResponse, error) {
	var err error

	req.Password, err = a.Hashing.Sum([]byte(req.Password))
	if err != nil {
		return response.RegisterResponse{}, errs.Err(err)
	}
	table, ok := models.AccountInput(req)
	if !ok {
		return response.RegisterResponse{
			Msg:  messages.ERR_INVALID_FIELDS,
			Code: statuscode.VALIDATION_ERR,
		}, nil
	}

	log.Println(table)
	err = a.Repo.Create(table)
	if err != nil {
		return response.RegisterResponse{}, errs.Err(err)
	}

	claims := map[string]string{
		"id": strconv.Itoa(int(table.ID)),
	}

	crp, err := a._createJWT(claims)
	if err != nil {
		return response.RegisterResponse{}, errs.Err(err)
	}

	return response.RegisterResponse{
		Data:  models.AccountOutput(table),
		Msg:   messages.SUCCESS_REGISTER,
		Token: crp,
		Code:  statuscode.SUCCESSFUL,
	}, nil
}

var _ authentication.AuthenticatorService = &AuthenticationSvc{}
