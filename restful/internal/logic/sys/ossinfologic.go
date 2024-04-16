package sys

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// OssInfoLogic 获取oss信息上下文
type OssInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewOssInfoLogic 新建获取oss信息上下文
func NewOssInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OssInfoLogic {
	return &OssInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "sys.OssInfo"),
		),
	}
}

// OssInfo 获取oss信息
func (l *OssInfoLogic) OssInfo(_ *types.OssInfoRep) (resp *types.OssInfoResp, err error) {
	return &types.OssInfoResp{BucketURL: "/static/upload/"}, nil
}
