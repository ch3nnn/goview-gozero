package sys

import (
	"context"

	userClient "github.com/ch3nnn/goview-gozero/service/user/client/user"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// UserLoginLogic 用户登录上下文
type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUserLoginLogic 新建用户登录上下文
func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "sys.UserLogin"),
		),
	}
}

// UserLogin 用户登录
func (l *UserLoginLogic) UserLogin(req *types.UserLoginRep) (resp *types.UserLoginResp, err error) {
	rpcResp, err := l.svcCtx.UserRpc.UserLogin(l.ctx, &userClient.UserLoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	var userInfo types.UserInfo
	if err := copier.Copy(&userInfo, rpcResp.GetUserInfo()); err != nil {
		return nil, err
	}

	var token types.Token
	if err := copier.Copy(&token, rpcResp.GetToken()); err != nil {
		return nil, err
	}

	return &types.UserLoginResp{
		UserInfo: userInfo,
		Token:    token,
	}, nil
}
