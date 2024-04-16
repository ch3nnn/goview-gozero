package userlogic

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/ch3nnn/goview-gozero/common/bizerr"

	"github.com/ch3nnn/goview-gozero/service/user/internal/dal/query"
	"github.com/ch3nnn/goview-gozero/service/user/internal/svc"
	"github.com/ch3nnn/goview-gozero/service/user/internal/utils"
	"github.com/ch3nnn/goview-gozero/service/user/pb"
	"github.com/ch3nnn/goview-gozero/service/user/types"
)

// UserLoginLogic 业务逻辑上下文
type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewUserLoginLogic 新建业务逻辑上下文
func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "User.UserLogin"),
		),
	}
}

// UserLogin 用户登录
func (l *UserLoginLogic) UserLogin(in *pb.UserLoginReq) (*pb.UserLoginResp, error) {
	sysUser, err := query.SysUser.WithContext(l.ctx).
		Where(
			query.SysUser.Username.Eq(in.GetUsername()),
			query.SysUser.Password.Eq(utils.EncryptPassword(in.GetPassword())),
		).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, bizerr.UserWrongUserOrPassword
		}
		return nil, err
	}

	jwt, err := utils.GenJWT("AccessSecret", 3600*24,
		utils.WithJWTOption("id", sysUser.ID),
		utils.WithJWTOption("username", sysUser.Username),
	)

	return &pb.UserLoginResp{
		Token:    types.ToJWT(sysUser.ID, jwt),
		UserInfo: types.ToUserInfo(sysUser),
	}, nil
}
