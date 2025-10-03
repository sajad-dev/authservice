package casbinz

import (
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/jinzhu/gorm"
	"github.com/sajad-dev/authservice/authorization/internal/shared/errors/errs"
)

var instanse *casbin.Enforcer
var once sync.Once

func _newInstanse(db *gorm.DB, model string) (*casbin.Enforcer, error) {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, errs.Err(err)
	}

	enforcer, err := casbin.NewEnforcer(model, adapter)
	if err != nil {
		return nil, errs.Err(err)
	}

	return enforcer, nil

}

func CreateInstanse(db *gorm.DB, model string) (*casbin.Enforcer, error) {
	var err error
	once.Do(func() {
		instanse, err = _newInstanse(db, model)
	})

	return instanse, err
}
