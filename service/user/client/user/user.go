// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package user

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"

	"github.com/ch3nnn/goview-gozero/service/user/pb"
)

// 类型定义
type (
	AddSysUserReq  = pb.AddSysUserReq
	AddSysUserResp = pb.AddSysUserResp
	SysUser        = pb.SysUser
	Token          = pb.Token
	TokenTag       = pb.TokenTag
	UserInfo       = pb.UserInfo
	UserLoginReq   = pb.UserLoginReq
	UserLoginResp  = pb.UserLoginResp

	User interface {
		// 用户登录
		UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error)
		// 创建用户表
		InsertSysUser(ctx context.Context, in *AddSysUserReq, opts ...grpc.CallOption) (*AddSysUserResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

// NewUser 新建 User 客户端
func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

// 用户登录
func (m *defaultUser) UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}

// 创建用户表
func (m *defaultUser) InsertSysUser(ctx context.Context, in *AddSysUserReq, opts ...grpc.CallOption) (*AddSysUserResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.InsertSysUser(ctx, in, opts...)
}
