package bootstrap

import (
	authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"github.com/jinzhu/gorm"
	"github.com/sajad-dev/authservice/authorization/internal/config"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/handler"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/repository"
	"github.com/sajad-dev/authservice/authorization/internal/domain/authorize/service"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/authorize"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/authorize/casbinz"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/crypto/hs256"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/setupdb"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/setupdb/postgres"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs"

	"google.golang.org/grpc"
)

type Bootstrap struct {
	Config config.AppConfig
}

func NewBootstrap(cnf config.AppConfig) *Bootstrap {
	return &Bootstrap{Config: cnf}
}

func (b *Bootstrap) _casbinInstanse(db *gorm.DB) (authorize.Authorize, error) {
	en, err := casbinz.CreateInstanse(db, b.Config.MODEL_CONF)
	if err != nil {
		return nil, errs.Err(err)
	}
	return casbinz.NewCasbin(en), nil
}

func (b *Bootstrap) _setupDB() setupdb.SetupDB[*gorm.DB] {
	return postgres.NewSetupPostgres(
		b.Config.DATABASE_PORT,
		b.Config.DATABASE_USER,
		b.Config.DATABASE_PASSWORD,
		b.Config.DATABASE_HOST,
		b.Config.DATABASE_NAME,
	)
}

func registerGrpc(gc *grpc.Server, repoDB authorize.Authorize) {
	authz.RegisterAuthorizationServer(gc, handler.NewAuthorizeHdlr(
		service.NewAuthorizeSvc(
			repository.NewAuthorizeRepo(repoDB),
			hs256.NewJWT([]byte("hi")),
		),
	))
}

func (b *Bootstrap) Boot(gc *grpc.Server) error {
	postgresDb, err := b._setupDB().Connection()
	if err != nil {
		return errs.Err(err)
	}

	en, err := b._casbinInstanse(postgresDb)
	if err != nil {
		return errs.Err(err)
	}

	registerGrpc(gc, en)

	return nil
}
