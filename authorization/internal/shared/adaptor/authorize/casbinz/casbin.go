package casbinz

import (
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs"
)

var instanse *casbin.Enforcer
var once sync.Once

func _newInstanse(adapter *gormadapter.Adapter,model string) (*casbin.Enforcer, error) {


	enforcer, err := casbin.NewEnforcer(model, adapter)
	if err != nil {
		return nil, errs.Err(err)
	}

	return enforcer, nil

}

func CreateInstanse(adapter *gormadapter.Adapter,model string) (*casbin.Enforcer, error) {
	var err error
	once.Do(func() {
		instanse, err = _newInstanse(adapter,model)
	})

	return instanse, err
}
