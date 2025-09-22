package crypto

import "time"

type DataClaims map[string]string

type Crypto interface {
	Generate(data DataClaims, expire time.Time) (string, error)
	Validate(tokenString string) (DataClaims, error)
}
