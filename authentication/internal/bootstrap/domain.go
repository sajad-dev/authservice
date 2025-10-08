package bootstrap

import (
	"github.com/sajad-dev/authservice/authentication/internal/config"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/hashing/sha256"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/setupdb"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/setupdb/setuppostgres"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb/postgres"
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
	"github.com/sajad-dev/authservice/authentication/internal/shared/validation"

	"github.com/sajad-dev/authservice/authentication/internal/domain/account/accountproto"
	accounthdlr "github.com/sajad-dev/authservice/authentication/internal/domain/account/handler"
	accountrepo "github.com/sajad-dev/authservice/authentication/internal/domain/account/repository"
	accountsvc "github.com/sajad-dev/authservice/authentication/internal/domain/account/service"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Bootstrap struct {
	Config config.Config
}

func NewBootstrap(cnf config.Config) *Bootstrap {
	return &Bootstrap{Config: cnf}
}

func (b *Bootstrap) _setupDB() setupdb.SetupDB[*gorm.DB] {
	return setuppostgres.NewSetupPostgres(
		b.Config.Database.Port,
		b.Config.Database.User,
		b.Config.Database.Password,
		b.Config.Database.Host,
		b.Config.Database.DbName,
	)
}

func (b *Bootstrap) _registerGrpc(gc *grpc.Server) error {

	postgresDb, err := b._setupDB().Connection()
	if err != nil {
		return errs.Err(err)
	}
	db := postgres.NewPostgres[*models.Accounts](postgresDb)

	vld := validation.NewValidator[*models.Accounts](db)

	hashing := sha256.NewSha256()

	accountproto.RegisterAccountServer(gc, accounthdlr.NewAccountHdlr(
		accountsvc.NewAccountSvc(
			accountrepo.NewAccountRepo(db),
			hashing,
		),
		vld,
	))

	return nil


}

func (b *Bootstrap) Boot(gc *grpc.Server) error {

	err := b._registerGrpc(gc)
	if err != nil {
		return errs.Err(err)
	}

	return nil
}
