package bootstrap

import (
	"github.com/sajad-dev/authservice/authentication/internal/config"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/crypto/hs256"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/hashing/sha256"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/setupdb"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/setupdb/setuppostgres"
	"github.com/sajad-dev/authservice/authentication/internal/shared/adaptor/sqldb/postgres"
	"github.com/sajad-dev/authservice/authentication/internal/shared/errors/errs"
	"github.com/sajad-dev/authservice/authentication/internal/shared/job"
	"github.com/sajad-dev/authservice/authentication/internal/shared/models"
	"github.com/sajad-dev/authservice/authentication/internal/shared/validation"

	"github.com/sajad-dev/authservice/authentication/internal/domain/account/accountproto"
	accounthdlr "github.com/sajad-dev/authservice/authentication/internal/domain/account/handler"
	accountrepo "github.com/sajad-dev/authservice/authentication/internal/domain/account/repository"
	accountsvc "github.com/sajad-dev/authservice/authentication/internal/domain/account/service"

	"github.com/sajad-dev/authservice/authentication/internal/domain/authentication/authenticationproto"
	authenticationhdlr "github.com/sajad-dev/authservice/authentication/internal/domain/authentication/handler"
	authenticationrepo "github.com/sajad-dev/authservice/authentication/internal/domain/authentication/repository"
	authenticationsvc "github.com/sajad-dev/authservice/authentication/internal/domain/authentication/service"

	"github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/forgetpasswordproto"
	forgetpasswordhdlr "github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/handler"
	forgetpasswordrepo "github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/repository"
	forgetpasswordsvc "github.com/sajad-dev/authservice/authentication/internal/domain/forgetpassword/service"

	twofactorhdlr "github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/handler"
	twofactorrepo "github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/repository"
	twofactorsvc "github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/service"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactor/twofactorproto"

	twofactornotifierhdlr "github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/handler"
	twofactornotifierrepo "github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/repository"
	twofactornotifiersvc "github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/service"
	"github.com/sajad-dev/authservice/authentication/internal/domain/twofactornotifier/twofactornotifierproto"

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
	dbCode := postgres.NewPostgres[*models.TwoFactorCode](postgresDb)

	validationDB := postgres.NewGlobalPostgres(postgresDb)

	vld := validation.NewValidator(validationDB)

	hashing := sha256.NewSha256()
	crp := hs256.NewJWT([]byte(b.Config.SecretKey))

	jb := job.NewJobs()

	accountproto.RegisterAccountServer(gc, accounthdlr.NewAccountHdlr(
		accountsvc.NewAccountSvc(
			accountrepo.NewAccountRepo(db),
			hashing,
		),
		vld,
	))

	authenticationproto.RegisterAuthenticationServer(gc, authenticationhdlr.NewAuthenticationHdlr(
		authenticationsvc.NewAuthenticationSvc(
			authenticationrepo.NewAuthenticationRepo(db),
			crp,
			hashing,
		),
		vld,
	))

	forgetpasswordproto.RegisterForgetPasswordServer(gc, forgetpasswordhdlr.NewForgetPasswordHdlr(
		forgetpasswordsvc.NewForgetPasswordSvc(
			forgetpasswordrepo.NewForgetPasswordRepo(db),
			crp,
			hashing,
			jb,
		),
		vld,
	))

	twofactorproto.RegisterTwofactorServer(gc, twofactorhdlr.NewTwoFactorHdlr(
		twofactorsvc.NewTwoFactorSvc(
			twofactorrepo.NewTwoFactorRepo(db, dbCode),
			crp,
			hashing,
		),
		vld,
	))

	twofactornotifierproto.RegisterTwofactorNotifierServer(gc, twofactornotifierhdlr.NewTwoFactorNotifierHdlr(
		twofactornotifiersvc.NewTwoFactorNotifierSvc(
			twofactornotifierrepo.NewTwoFactorNotifierRepo(db, dbCode),
			crp,
			jb,
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
