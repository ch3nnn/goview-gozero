package config

import (
	"github.com/ch3nnn/goview-gozero/pkg/config"
	"github.com/zeromicro/go-zero/zrpc"
)

// Config 全局相关配置
type Config struct {
	zrpc.RpcServerConf
	MysqlDBConf config.DatabaseConf
}
