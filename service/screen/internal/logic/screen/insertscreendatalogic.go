package screenlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ch3nnn/goview-gozero/service/screen/internal/dal/model"
	"github.com/ch3nnn/goview-gozero/service/screen/internal/dal/query"
	"github.com/ch3nnn/goview-gozero/service/screen/internal/svc"
	"github.com/ch3nnn/goview-gozero/service/screen/pb"
)

// InsertScreenDataLogic 业务逻辑上下文
type InsertScreenDataLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewInsertScreenDataLogic 新建业务逻辑上下文
func NewInsertScreenDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertScreenDataLogic {
	return &InsertScreenDataLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "Screen.InsertScreenData"),
		),
	}
}

// InsertScreenData 创建大屏数据
func (l *InsertScreenDataLogic) InsertScreenData(in *pb.AddScreenDataReq) (*pb.AddScreenDataResp, error) {
	if err := query.ScreenDatum.WithContext(l.ctx).Create(&model.ScreenDatum{
		ProjectID: in.GetProjectId(),
		Content:   in.GetContent(),
	}); err != nil {
		return nil, err
	}

	return &pb.AddScreenDataResp{}, nil
}
