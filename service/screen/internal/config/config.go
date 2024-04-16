package config

import (
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/ch3nnn/goview-gozero/pkg/config"
)

// Config 全局相关配置
type Config struct {
	zrpc.RpcServerConf
	MysqlDBConf config.DatabaseConf
}
