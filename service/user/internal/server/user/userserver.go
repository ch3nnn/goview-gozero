// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"github.com/ch3nnn/goview-gozero/service/user/internal/logic/user"
	"github.com/ch3nnn/goview-gozero/service/user/internal/svc"
	"github.com/ch3nnn/goview-gozero/service/user/pb"
)

// UserServer User 服务器结构
type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

// NewUserServer 新建 User 服务器
func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// 用户登录
func (s *UserServer) UserLogin(ctx context.Context, in *pb.UserLoginReq) (*pb.UserLoginResp, error) {
	l := userlogic.NewUserLoginLogic(ctx, s.svcCtx)
	return l.UserLogin(in)
}

// 创建用户表
func (s *UserServer) InsertSysUser(ctx context.Context, in *pb.AddSysUserReq) (*pb.AddSysUserResp, error) {
	l := userlogic.NewInsertSysUserLogic(ctx, s.svcCtx)
	return l.InsertSysUser(in)
}
