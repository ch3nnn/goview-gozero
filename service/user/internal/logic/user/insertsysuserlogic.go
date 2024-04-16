package userlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ch3nnn/goview-gozero/service/user/internal/dal/model"
	"github.com/ch3nnn/goview-gozero/service/user/internal/dal/query"
	"github.com/ch3nnn/goview-gozero/service/user/internal/svc"
	"github.com/ch3nnn/goview-gozero/service/user/internal/utils"
	"github.com/ch3nnn/goview-gozero/service/user/pb"
)

// InsertSysUserLogic 业务逻辑上下文
type InsertSysUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewInsertSysUserLogic 新建业务逻辑上下文
func NewInsertSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertSysUserLogic {
	return &InsertSysUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "User.InsertSysUser"),
		),
	}
}

// InsertSysUser 创建用户表
func (l *InsertSysUserLogic) InsertSysUser(in *pb.AddSysUserReq) (*pb.AddSysUserResp, error) {
	if err := query.SysUser.WithContext(l.ctx).Create(&model.SysUser{
		Username: in.GetUsername(),
		Password: utils.EncryptPassword(in.GetPassword()),
		Nickname: in.GetPassword(),
	}); err != nil {
		return nil, err
	}

	return &pb.AddSysUserResp{}, nil
}
