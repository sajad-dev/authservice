package modeldto_test

import (
	"testing"

	"github.com/sajad-dev/authservice/internal/domain/account/dto/modeldto"
)

type TestAccount struct{
	FirstName string
}

func TestAccountdto (t *testing.T) {
	x := modeldto.AccountInput(TestAccount{FirstName: "hi"})
	t.Log(x.LastName)
}
