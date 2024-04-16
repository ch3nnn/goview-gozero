package screen

import (
	"context"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/zeromicro/go-zero/core/logx"

	screenClient "github.com/ch3nnn/goview-gozero/service/screen/client/screen"

	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// UpdateScreenProjectLogic 更新大屏信息上下文
type UpdateScreenProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewUpdateScreenProjectLogic 新建更新大屏信息上下文
func NewUpdateScreenProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateScreenProjectLogic {
	return &UpdateScreenProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "screen.UpdateScreenProject"),
		),
	}
}

// UpdateScreenProject 更新大屏信息
func (l *UpdateScreenProjectLogic) UpdateScreenProject(req *types.UpdateScreenProjectReq) (resp *types.UpdateScreenProjectResp, err error) {
	toInt, err := convertor.ToInt(req.Id)
	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.ScreenRpc.UpdateScreenProject(l.ctx, &screenClient.UpdateScreenProjectReq{
		Id:       toInt,
		Name:     req.Name,
		IndexImg: req.IndexImg,
		Remark:   req.Remark,
	})
	if err != nil {
		return nil, err
	}

	return
}
