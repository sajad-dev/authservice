package bootstrap

import (
	"log"

	gormadapter "github.com/casbin/gorm-adapter/v2"
	authz "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	"github.com/go-playground/validator"
	"github.com/sajad-dev/authservice/authorization/internal/config"

	authorizehdlr "github.com/sajad-dev/authservice/authorization/internal/domain/authorize/handler"
	authorizerepo "github.com/sajad-dev/authservice/authorization/internal/domain/authorize/repository"
	authorizesvc "github.com/sajad-dev/authservice/authorization/internal/domain/authorize/service"

	grouphdlr "github.com/sajad-dev/authservice/authorization/internal/domain/group/handler"
	grouprepo "github.com/sajad-dev/authservice/authorization/internal/domain/group/repository"
	groupsvc "github.com/sajad-dev/authservice/authorization/internal/domain/group/service"

	policyhdlr "github.com/sajad-dev/authservice/authorization/internal/domain/policy/handler"
	"github.com/sajad-dev/authservice/authorization/internal/domain/policy/policyproto"
	policyrepo "github.com/sajad-dev/authservice/authorization/internal/domain/policy/repository"
	policysvc "github.com/sajad-dev/authservice/authorization/internal/domain/policy/service"

	"github.com/sajad-dev/authservice/authorization/internal/domain/group/groupproto"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/authorize"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/authorize/casbinz"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/crypto/hs256"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/setupdb"
	"github.com/sajad-dev/authservice/authorization/internal/shared/adaptor/setupdb/postgres"
	"github.com/sajad-dev/authservice/authorization/internal/shared/validation"
	"github.com/sajad-dev/authservice/authorization/internal/shared/validation"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs"

	"google.golang.org/grpc"
)

type Bootstrap struct {
	Config config.Config
}

func NewBootstrap(cnf config.Config) *Bootstrap {
	return &Bootstrap{Config: cnf}
}

func (b *Bootstrap) _casbinInstanse(adapter *gormadapter.Adapter) (authorize.Authorize, error) {
	log.Println("Boot GRPC")

	en, err := casbinz.CreateInstanse(adapter, b.Config.ConfigModel)
	if err != nil {
		return nil, errs.Err(err)
	}
	return casbinz.NewCasbin(en), nil
}

func (b *Bootstrap) _setupDB() setupdb.SetupDB[*gormadapter.Adapter] {
	return postgres.NewSetupPostgres(
		b.Config.Database.Port,
		b.Config.Database.User,
		b.Config.Database.Password,
		b.Config.Database.Host,
		b.Config.Database.DbName,
	)
}

func (b *Bootstrap) _registerGrpc(gc *grpc.Server, repoDB authorize.Authorize, vld validation.Validation) {
	authz.RegisterAuthorizationServer(gc, authorizehdlr.NewAuthorizeHdlr(
		authorizesvc.NewAuthorizeSvc(
			authorizerepo.NewAuthorizeRepo(repoDB),
			hs256.NewJWT([]byte(b.Config.SecretKey)),
		),
	))

	groupproto.RegisterGroupServer(gc, grouphdlr.NewGroupHdlr(
		groupsvc.NewGroupSvc(
			grouprepo.NewGroupRepo(repoDB),
		),
		vld,
	))

	policyproto.RegisterPolicyServer(gc, policyhdlr.NewPolicyHdlr(
		policysvc.NewPolicySvc(
			policyrepo.NewPolicyRepo(repoDB),
		),
		vld,
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

	vld := validate.NewValidate(validator.New())

	b._registerGrpc(gc, en, vld)

	return nil
}
