package casbin

import (
	"go_web_demo/models"
	"go_web_demo/pkg/setting"
	"os"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
)

var (
	Enforcer *casbin.Enforcer
)

// 初始化 casbin
func InitCasbin() (*casbin.Enforcer, error) {

	// 判断是否有缓存
	if Enforcer != nil {
		return Enforcer, nil
	}
	adapter, err := gormadapter.NewAdapterByDB(models.Db)
	if err != nil {
		return nil, err
	}

	if setting.CasbinConf.Path == "" {
		wd, _ := os.Getwd()
		setting.CasbinConf.Path = wd + "/conf/rbac_model.conf"
	}
	Enforcer, err = casbin.NewEnforcer(setting.CasbinConf.Path, adapter)
	if err != nil {
		return nil, err
	}

	if err = Enforcer.LoadPolicy(); err != nil {
		return nil, err
	}

	return Enforcer, nil
}

// 清空缓存
func ClearEnforcer() {
	Enforcer = nil
}
