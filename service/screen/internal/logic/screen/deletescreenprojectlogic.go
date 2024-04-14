package screenlogic

import (
	"context"

	"github.com/ch3nnn/goview-gozero/service/screen/internal/dal/query"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ch3nnn/goview-gozero/service/screen/internal/svc"
	"github.com/ch3nnn/goview-gozero/service/screen/pb"
)

// DeleteScreenProjectLogic 业务逻辑上下文
type DeleteScreenProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewDeleteScreenProjectLogic 新建业务逻辑上下文
func NewDeleteScreenProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteScreenProjectLogic {
	return &DeleteScreenProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "Screen.DeleteScreenProject"),
		),
	}
}

// DeleteScreenProject 根据大屏信息ID删除
func (l *DeleteScreenProjectLogic) DeleteScreenProject(in *pb.DelScreenProjectReq) (*pb.DelScreenProjectResp, error) {
	if _, err := query.ScreenProject.WithContext(l.ctx).
		Where(query.ScreenProject.ID.Eq(in.GetId())).
		Delete(); err != nil {
		return nil, err
	}

	return &pb.DelScreenProjectResp{}, nil
}
