package screenlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ch3nnn/goview-gozero/service/screen/internal/svc"
	"github.com/ch3nnn/goview-gozero/service/screen/pb"
)

// SelectScreenProjectByIdLogic 业务逻辑上下文
type SelectScreenProjectByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewSelectScreenProjectByIdLogic 新建业务逻辑上下文
func NewSelectScreenProjectByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectScreenProjectByIdLogic {
	return &SelectScreenProjectByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "Screen.SelectScreenProjectById"),
		),
	}
}

// SelectScreenProjectById 根据大屏信息ID获取详情
func (l *SelectScreenProjectByIdLogic) SelectScreenProjectById(in *pb.SelectScreenProjectByIdReq) (*pb.SelectScreenProjectByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SelectScreenProjectByIdResp{}, nil
}
