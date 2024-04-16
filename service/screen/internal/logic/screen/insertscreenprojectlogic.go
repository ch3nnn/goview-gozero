package screenlogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ch3nnn/goview-gozero/service/screen/internal/dal/model"
	"github.com/ch3nnn/goview-gozero/service/screen/internal/dal/query"
	"github.com/ch3nnn/goview-gozero/service/screen/internal/svc"
	"github.com/ch3nnn/goview-gozero/service/screen/pb"
)

// InsertScreenProjectLogic 业务逻辑上下文
type InsertScreenProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewInsertScreenProjectLogic 新建业务逻辑上下文
func NewInsertScreenProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertScreenProjectLogic {
	return &InsertScreenProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "Screen.InsertScreenProject"),
		),
	}
}

// InsertScreenProject 创建大屏信息
func (l *InsertScreenProjectLogic) InsertScreenProject(in *pb.AddScreenProjectReq) (*pb.AddScreenProjectResp, error) {
	sp := model.ScreenProject{
		Name:     in.GetName(),
		State:    1,
		UserID:   in.GetUserId(),
		IndexImg: in.GetIndexImg(),
		Remark:   in.GetRemark(),
	}

	if err := query.ScreenProject.WithContext(l.ctx).Create(&sp); err != nil {
		return nil, err
	}

	return &pb.AddScreenProjectResp{
		Project: &pb.ScreenProject{
			Id:       sp.ID,
			Name:     sp.Name,
			State:    sp.State,
			UserId:   sp.UserID,
			IndexImg: sp.IndexImg,
			Remark:   sp.IndexImg,
			IsDel:    false,
			CreateAt: sp.CreateAt.Unix(),
		},
	}, nil
}
