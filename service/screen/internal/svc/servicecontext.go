package svc

import (
	"github.com/ch3nnn/goview-gozero/service/screen/internal/config"
	"github.com/ch3nnn/goview-gozero/service/screen/internal/dal/query"
)

// ServiceContext 服务上下文
type ServiceContext struct {
	Config config.Config
}

// NewServiceContext 新建服务上下文
func NewServiceContext(c config.Config) *ServiceContext {
	query.SetDefault(c.MysqlDBConf.NewDriver())

	return &ServiceContext{
		Config: c,
	}
}
