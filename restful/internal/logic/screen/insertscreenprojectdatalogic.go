package screen

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	screenClient "github.com/ch3nnn/goview-gozero/service/screen/client/screen"

	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// InsertScreenProjectDataLogic 保存大屏数据上下文
type InsertScreenProjectDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewInsertScreenProjectDataLogic 新建保存大屏数据上下文
func NewInsertScreenProjectDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertScreenProjectDataLogic {
	return &InsertScreenProjectDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "screen.InsertScreenProjectData"),
		),
	}
}

// InsertScreenProjectData 保存大屏数据
func (l *InsertScreenProjectDataLogic) InsertScreenProjectData(req *types.InsertScreenProjectDataReq) (resp *types.InsertScreenProjectDataResp, err error) {
	_, err = l.svcCtx.ScreenRpc.InsertScreenData(l.ctx, &screenClient.AddScreenDataReq{
		ProjectId: req.ProjectId,
		Content:   req.Content,
	})
	if err != nil {
		return nil, err
	}

	return
}
