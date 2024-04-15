package bizerr

import "github.com/ch3nnn/goview-gozero/pkg/errcode"

// RPC 业务错误码
//
// 错误码定义规则: 服务名+Rpc+错误描述
// 根据服务端口号：40200，取 40201 为起始错误码
//
// 例如 "用户服务RPC" 用户或密码错误 UserWrongUserOrPassword = errcode.New(10001, "用户或密码错误")

var UserWrongUserOrPassword = errcode.New(10001, "用户或密码错误")
var ScreenRpcProjectNotFound = errcode.New(10002, "大屏信息不存在")
