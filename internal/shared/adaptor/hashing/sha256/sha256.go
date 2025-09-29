package sha256

import (
	al "crypto/sha256"
	"encoding/hex"

	"github.com/sajad-dev/authservice/internal/shared/adaptor/hashing"
	"github.com/sajad-dev/authservice/internal/shared/errors/errs"
)

type Sha256 struct{}

func NewSha256 () *Sha256{
	return &Sha256{}
}

func (s *Sha256) Sum(data []byte) (string, error) {
	hashSha256 := al.New()
	_, err := hashSha256.Write(data)
	if err != nil {
		return "", errs.Err(err)
	}
	hashHex := hashSha256.Sum(nil)
	return hex.EncodeToString(hashHex[:]), nil
}

var _ hashing.Hashing = &Sha256{}
