package hs256

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/crypto"
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs"
)

type JWT struct {
	Key []byte
}

type Claims struct {
	Data crypto.DataClaims
	jwt.RegisteredClaims
}

func NewJWT(key []byte) *JWT {
	return &JWT{Key: key}
}

func (j JWT) Generate(data crypto.DataClaims, expire time.Time) (string, error) {
	claims := Claims{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenSig, err := token.SignedString(j.Key)
	if err != nil {
		return "", errs.Err(err)
	}
	return tokenSig, nil
}

func (j JWT) Validate(tokenString string) (crypto.DataClaims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.Key, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims.Data, nil
}

var _ crypto.Crypto = &JWT{}
