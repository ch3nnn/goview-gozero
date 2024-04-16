package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	screenClient "github.com/ch3nnn/goview-gozero/service/screen/client/screen"
	userClient "github.com/ch3nnn/goview-gozero/service/user/client/user"

	"github.com/ch3nnn/goview-gozero/restful/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserRpc   userClient.User
	ScreenRpc screenClient.Screen
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		UserRpc:   userClient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ScreenRpc: screenClient.NewScreen(zrpc.MustNewClient(c.ScreenRpc)),
	}
}
