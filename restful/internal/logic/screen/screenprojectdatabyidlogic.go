package screen

import (
	"context"
	"strconv"

	screenClient "github.com/ch3nnn/goview-gozero/service/screen/client/screen"
	"github.com/ch3nnn/goview-gozero/service/screen/pb"
	"github.com/golang-module/carbon/v2"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/fx"
	"github.com/zeromicro/go-zero/core/logx"

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
		batchError errorx.BatchError
		datumRpc   *pb.SelectScreenDataByIdResp
		projectRpc *pb.SelectScreenProjectByIdResp
	)

	fx.Parallel(
		func() {
			datumRpc, err = l.svcCtx.ScreenRpc.SelectScreenDataById(l.ctx, &screenClient.SelectScreenDataByIdReq{ProjectId: req.ProjectId})
			if err != nil {
				batchError.Add(err)
			}
		},
		func() {
			projectRpc, err = l.svcCtx.ScreenRpc.SelectScreenProjectById(l.ctx, &screenClient.SelectScreenProjectByIdReq{Id: req.ProjectId})
			if err != nil {
				batchError.Add(err)
			}
		},
	)

	if batchError.NotNil() {
		return nil, batchError.Err()
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
		IsDel:        1,
		CreateUserId: projectRpc.GetScreenProject().GetUserId(),
		CreateAt:     carbon.CreateFromTimestamp(datumRpc.GetScreenData().GetCreateAt()).ToString(),
	}, nil
}
