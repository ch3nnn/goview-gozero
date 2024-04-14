package screenlogic

import (
	"context"

	"github.com/ch3nnn/goview-gozero/service/screen/internal/dal/query"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ch3nnn/goview-gozero/service/screen/internal/svc"
	"github.com/ch3nnn/goview-gozero/service/screen/pb"
)

// UpdateScreenProjectLogic 业务逻辑上下文
type UpdateScreenProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewUpdateScreenProjectLogic 新建业务逻辑上下文
func NewUpdateScreenProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateScreenProjectLogic {
	return &UpdateScreenProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "Screen.UpdateScreenProject"),
		),
	}
}

// UpdateScreenProject 更新大屏信息
func (l *UpdateScreenProjectLogic) UpdateScreenProject(in *pb.UpdateScreenProjectReq) (*pb.UpdateScreenProjectResp, error) {
	if _, err := query.ScreenProject.WithContext(l.ctx).
		Where(query.ScreenProject.ID.Eq(in.GetId())).
		Updates(in); err != nil {
		return nil, err
	}

	return &pb.UpdateScreenProjectResp{}, nil
}
