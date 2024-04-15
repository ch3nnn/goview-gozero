package screen

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// UploadScreenProjectFileLogic 文件上传上下文
type UploadScreenProjectFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUploadScreenProjectFileLogic 新建文件上传上下文
func NewUploadScreenProjectFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadScreenProjectFileLogic {
	return &UploadScreenProjectFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "screen.UploadScreenProjectFile"),
		),
	}
}

// UploadScreenProjectFile 文件上传
func (l *UploadScreenProjectFileLogic) UploadScreenProjectFile(req *types.UploadScreenProjectFileReq) (resp *types.UploadScreenProjectFileResp, err error) {
	return &types.UploadScreenProjectFileResp{Filename: req.Filename}, nil
}
