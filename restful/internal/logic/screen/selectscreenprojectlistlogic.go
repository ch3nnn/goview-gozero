package screen

import (
	"context"

	carbon "github.com/golang-module/carbon/v2"
	"github.com/zeromicro/go-zero/core/logx"

	screenClient "github.com/ch3nnn/goview-gozero/service/screen/client/screen"

	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// SelectScreenProjectListLogic 大屏信息列表上下文
type SelectScreenProjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewSelectScreenProjectListLogic 新建大屏信息列表上下文
func NewSelectScreenProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectScreenProjectListLogic {
	return &SelectScreenProjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "screen.SelectScreenProjectList"),
		),
	}
}

// SelectScreenProjectList 大屏信息列表
func (l *SelectScreenProjectListLogic) SelectScreenProjectList(req *types.SelectScreenProjectListReq) (resp *types.SelectScreenProjectListResp, err error) {
	rpcResp, err := l.svcCtx.ScreenRpc.SelectScreenProjectList(l.ctx, &screenClient.SelectScreenProjectListReq{
		Page:     req.Page,
		PageSize: req.Limit,
	})
	if err != nil {
		return nil, err
	}

	projects := make([]types.ScreenProject, len(rpcResp.GetResults()))
	for i, project := range rpcResp.GetResults() {
		projects[i] = types.ScreenProject{
			Id:       project.GetId(),
			Name:     project.GetName(),
			State:    project.GetState(),
			UserId:   project.GetUserId(),
			IndexImg: project.GetIndexImg(),
			Remark:   project.GetRemark(),
			IsDel:    project.GetIsDel(),
			CreateAt: carbon.CreateFromTimestamp(project.GetCreateAt()).ToString(),
		}
	}

	return &types.SelectScreenProjectListResp{
		Count:   rpcResp.GetCount(),
		Results: projects,
	}, nil
}
