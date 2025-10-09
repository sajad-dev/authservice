package seeder

import (
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/hashing"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
)

func AccountSeeder(db sqldb.SqlDB[*models.Accounts], hash hashing.Hashing) error {
	password, err := hash.Sum([]byte("I am best programmer in the world"))
	if err != nil {
		return err
	}
	db.Create(models.NewAccounts(
		models.WithFirstName("Mohammad sajjad"),
		models.WithLastName("Pourajam"),
		models.WithEmail("prj.sajad85@gmail.com"),
		models.WithUsername("sajad"),
		models.WithPassword(password),
	))
	return nil
}
