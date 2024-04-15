package screen

import (
	"context"

	screenClient "github.com/ch3nnn/goview-gozero/service/screen/client/screen"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// DeleteScreenProjectLogic 根据大屏信息ID删除上下文
type DeleteScreenProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewDeleteScreenProjectLogic 新建根据大屏信息ID删除上下文
func NewDeleteScreenProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteScreenProjectLogic {
	return &DeleteScreenProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "screen.DeleteScreenProject"),
		),
	}
}

// DeleteScreenProject 根据大屏信息ID删除
func (l *DeleteScreenProjectLogic) DeleteScreenProject(req *types.DelScreenProjectReq) (resp *types.DelScreenProjectResp, err error) {
	if _, err := l.svcCtx.ScreenRpc.DeleteScreenProject(l.ctx, &screenClient.DelScreenProjectReq{Id: req.Ids}); err != nil {
		return nil, err
	}

	return
}
