package validation_test

import (
	"testing"

	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb/postgres"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
	"github.com/sajad-dev/authservice/authentication/internal/shared/validation"
)

func TestValidation (t *testing.T) {
	 validation.NewValidator(postgres.NewPostgres[*models.Accounts](nil))

}
