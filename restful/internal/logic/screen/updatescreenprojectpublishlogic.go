package screen

import (
	"context"

	screenClient "github.com/ch3nnn/goview-gozero/service/screen/client/screen"
	"github.com/duke-git/lancet/v2/pointer"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// UpdateScreenProjectPublishLogic 更新大屏信息发布状态上下文
type UpdateScreenProjectPublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUpdateScreenProjectPublishLogic 新建更新大屏信息发布状态上下文
func NewUpdateScreenProjectPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateScreenProjectPublishLogic {
	return &UpdateScreenProjectPublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "screen.UpdateScreenProjectPublish"),
		),
	}
}

// UpdateScreenProjectPublish 更新大屏信息发布状态
func (l *UpdateScreenProjectPublishLogic) UpdateScreenProjectPublish(req *types.UpdateScreenProjectPublishReq) (resp *types.UpdateScreenProjectPublishResp, err error) {
	_, err = l.svcCtx.ScreenRpc.UpdateScreenProject(l.ctx, &screenClient.UpdateScreenProjectReq{
		Id:    req.Id,
		State: pointer.Of(req.State),
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
