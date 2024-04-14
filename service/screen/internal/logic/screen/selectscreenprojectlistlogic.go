package screenlogic

import (
	"context"

	"github.com/ch3nnn/goview-gozero/pkg/math"
	"github.com/ch3nnn/goview-gozero/service/screen/internal/dal/query"
	"github.com/ch3nnn/goview-gozero/service/screen/internal/types"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ch3nnn/goview-gozero/service/screen/internal/svc"
	"github.com/ch3nnn/goview-gozero/service/screen/pb"
)

// SelectScreenProjectListLogic 业务逻辑上下文
type SelectScreenProjectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewSelectScreenProjectListLogic 新建业务逻辑上下文
func NewSelectScreenProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectScreenProjectListLogic {
	return &SelectScreenProjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "Screen.SelectScreenProjectList"),
		),
	}
}

// SelectScreenProjectList 大屏信息列表
func (l *SelectScreenProjectListLogic) SelectScreenProjectList(in *pb.SelectScreenProjectListReq) (*pb.SelectScreenProjectListResp, error) {
	var projectDos []*types.ScreenProjectDO
	count, err := query.ScreenProject.WithContext(l.ctx).ScanByPage(&projectDos, int((in.Page-1)*in.PageSize), int(in.PageSize))
	if err != nil {
		return nil, err
	}

	var projects []*pb.ScreenProject
	if err := copier.Copy(&projects, &projectDos); err != nil {
		return nil, err
	}

	return &pb.SelectScreenProjectListResp{
		Count:     count,
		PageCount: math.Ceil(count, in.PageSize),
		Results:   projects,
	}, nil
}
