package screen

import (
	"context"
	"encoding/json"

	carbon "github.com/golang-module/carbon/v2"
	"github.com/zeromicro/go-zero/core/logx"

	screenClient "github.com/ch3nnn/goview-gozero/service/screen/client/screen"

	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// InsertScreenProjectLogic 创建大屏信息上下文
type InsertScreenProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewInsertScreenProjectLogic 新建创建大屏信息上下文
func NewInsertScreenProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertScreenProjectLogic {
	return &InsertScreenProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "screen.InsertScreenProject"),
		),
	}
}

// InsertScreenProject 创建大屏信息
func (l *InsertScreenProjectLogic) InsertScreenProject(req *types.AddScreenProjectReq) (resp *types.AddScreenProjectResp, err error) {
	userId, _ := l.ctx.Value("id").(json.Number).Int64()
	rpcResp, err := l.svcCtx.ScreenRpc.InsertScreenProject(l.ctx, &screenClient.AddScreenProjectReq{
		Name:     req.Name,
		UserId:   userId,
		IndexImg: req.IndexImg,
		Remark:   req.Remark,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddScreenProjectResp{
		Id:       rpcResp.GetProject().GetId(),
		Name:     rpcResp.GetProject().GetName(),
		State:    rpcResp.GetProject().GetState(),
		UserId:   rpcResp.GetProject().GetUserId(),
		IndexImg: rpcResp.GetProject().GetIndexImg(),
		Remark:   rpcResp.GetProject().GetRemark(),
		IsDel:    rpcResp.GetProject().GetIsDel(),
		CreateAt: carbon.CreateFromTimestamp(rpcResp.GetProject().GetCreateAt()).ToString(),
	}, nil
}
