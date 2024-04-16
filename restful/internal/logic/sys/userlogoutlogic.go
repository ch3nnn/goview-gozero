package sys

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// UserLogoutLogic 用户登出上下文
type UserLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUserLogoutLogic 新建用户登出上下文
func NewUserLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogoutLogic {
	return &UserLogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "sys.UserLogout"),
		),
	}
}

// UserLogout 用户登出
func (l *UserLogoutLogic) UserLogout(req *types.UserLogoutRep) (resp *types.UserLogoutResp, err error) {
	// todo: add your logic here and delete this line

	return
}
