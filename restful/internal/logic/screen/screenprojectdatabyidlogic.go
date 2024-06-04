package screen

import (
	"context"
	"strconv"

	carbon "github.com/golang-module/carbon/v2"
	"github.com/zeromicro/go-zero/core/fx"
	"github.com/zeromicro/go-zero/core/logx"

	screenClient "github.com/ch3nnn/goview-gozero/service/screen/client/screen"
	"github.com/ch3nnn/goview-gozero/service/screen/pb"

	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// ScreenProjectDataByIDLogic 根据大屏信息ID获取大屏数据上下文
type ScreenProjectDataByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewScreenProjectDataByIDLogic 新建根据大屏信息ID获取大屏数据上下文
func NewScreenProjectDataByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScreenProjectDataByIDLogic {
	return &ScreenProjectDataByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "screen.ScreenProjectDataByID"),
		),
	}
}

// ScreenProjectDataByID 根据大屏信息ID获取大屏数据
func (l *ScreenProjectDataByIDLogic) ScreenProjectDataByID(req *types.ScreenProjectDataByIDReq) (*types.ScreenProjectDataByIDResp, error) {
	var (
		err        error
		datumRpc   *pb.SelectScreenDataByIdResp
		projectRpc *pb.SelectScreenProjectByIdResp
	)

	err = fx.ParallelErr(
		func() error {
			datumRpc, err = l.svcCtx.ScreenRpc.SelectScreenDataById(l.ctx, &screenClient.SelectScreenDataByIdReq{
				ProjectId: req.ProjectId,
			})

			return err
		},
		func() error {
			projectRpc, err = l.svcCtx.ScreenRpc.SelectScreenProjectById(l.ctx, &screenClient.SelectScreenProjectByIdReq{
				Id: req.ProjectId,
			})

			return err
		},
	)
	if err != nil {
		return nil, err
	}

	if datumRpc.GetScreenData() == nil {
		return nil, nil
	}

	return &types.ScreenProjectDataByIDResp{
		Id:           strconv.FormatInt(projectRpc.GetScreenProject().GetId(), 10),
		Name:         projectRpc.GetScreenProject().GetName(),
		State:        projectRpc.GetScreenProject().GetState(),
		Content:      datumRpc.GetScreenData().GetContent(),
		IndexImage:   projectRpc.GetScreenProject().GetIndexImg(),
		Remark:       projectRpc.GetScreenProject().GetRemark(),
		CreateUserId: projectRpc.GetScreenProject().GetUserId(),
		CreateAt:     carbon.CreateFromTimestamp(datumRpc.GetScreenData().GetCreateAt()).ToString(),
		IsDel:        map[bool]int64{true: 1, false: -1}[projectRpc.GetScreenProject().GetIsDel()],
	}, nil
}
