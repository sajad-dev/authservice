package service

import (
	"encoding/json"
	"log"
	"os"
	"slices"
	"strings"

	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/dto/request"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/dto/response"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/crypto"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/messages"
	"github.com/sajad-dev/authservice/authorization/internal/shared/constants/statuscode"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs"
)

type AuthorizeSvc struct {
	Repo   authorize.AuthorizeRepository
	Crypto crypto.Crypto
	Guest  string
}

type guest struct {
	Path   string `json:"path"`
	Action string `json:"action"`
}

func NewAuthorizeSvc(repo authorize.AuthorizeRepository, crp crypto.Crypto, gst string) *AuthorizeSvc {
	return &AuthorizeSvc{
		Repo:   repo,
		Crypto: crp,
		Guest:  gst,
	}
}

func (a AuthorizeSvc) _verifyJWT(token string) (crypto.DataClaims, error) {
	return a.Crypto.Validate(token)
}

func _readJson(gst string) ([]string, error) {
	file, err := os.Open(gst)
	if err != nil {
		return []string{}, errs.Err(err)
	}
	defer file.Close()

	var gstJson []string
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&gstJson)
	if err != nil {
		return []string{}, errs.Err(err)
	}
	return gstJson, nil
}

func (a *AuthorizeSvc) Check(req request.AuthorizeRequest) (response.AuthorizeResponse, error) {
	gst, err := _readJson(a.Guest)
	if err != nil {
		return response.AuthorizeResponse{}, err
	}

	ok := slices.Contains(gst, req.Path)
	if ok {
		return response.AuthorizeResponse{
			Code: statuscode.SUCCESSFUL,
			Msg:  messages.SUCCESS_VERIFY,
		}, nil
	}

	authorization, ok := req.Headers["authorization"]
	log.Println(authorization,req.Headers)
	if !ok {
		return response.AuthorizeResponse{
			Code: statuscode.PERMISSION_DENIED,
			Msg:  messages.ERR_FORBIDDEN,
		}, nil
	}

	extracted := strings.Fields(authorization)
	if len(extracted) != 2 || extracted[0] != "Bearer" {
		return response.AuthorizeResponse{
			Code: statuscode.PERMISSION_DENIED,
			Msg:  messages.ERR_FORBIDDEN,
		}, nil
	}

	claims, err := a._verifyJWT(extracted[1])
	if err != nil {
		return response.AuthorizeResponse{}, errs.Err(err)
	}

	user, ok := claims["user"]
	if !ok {
		return response.AuthorizeResponse{
			Code: statuscode.PERMISSION_DENIED,
			Msg:  messages.ERR_FORBIDDEN,
		}, nil
	}

	ok, err = a.Repo.Check(user, req.Path, req.Method)
	if !ok {
		return response.AuthorizeResponse{
			Code: statuscode.PERMISSION_DENIED,
			Msg:  messages.ERR_FORBIDDEN,
		}, nil

	}

	if err != nil {
		return response.AuthorizeResponse{}, errs.Err(err)
	}

	return response.AuthorizeResponse{
		Code: statuscode.SUCCESSFUL,
		Msg:  messages.SUCCESS_VERIFY,
	}, nil

}

var _ authorize.AuthorizeService = &AuthorizeSvc{}
