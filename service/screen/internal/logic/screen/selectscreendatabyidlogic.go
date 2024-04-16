package screenlogic

import (
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/ch3nnn/goview-gozero/service/screen/internal/dal/query"
	"github.com/ch3nnn/goview-gozero/service/screen/internal/svc"
	"github.com/ch3nnn/goview-gozero/service/screen/pb"
)

// SelectScreenDataByIdLogic 业务逻辑上下文
type SelectScreenDataByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewSelectScreenDataByIdLogic 新建业务逻辑上下文
func NewSelectScreenDataByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectScreenDataByIdLogic {
	return &SelectScreenDataByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx).WithFields(
			logx.Field("service", svcCtx.Config.Name),
			logx.Field("method", "Screen.SelectScreenDataById"),
		),
	}
}

// SelectScreenDataById 根据大屏数据ID获取详情
func (l *SelectScreenDataByIdLogic) SelectScreenDataById(in *pb.SelectScreenDataByIdReq) (*pb.SelectScreenDataByIdResp, error) {
	screenDatum, err := query.ScreenDatum.WithContext(l.ctx).
		Where(query.ScreenDatum.ProjectID.Eq(in.GetProjectId())).
		Order(query.ScreenDatum.ID.Desc()).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &pb.SelectScreenDataByIdResp{ScreenData: &pb.ScreenData{
		Id:        screenDatum.ID,
		ProjectId: screenDatum.ProjectID,
		UserId:    0,
		Content:   screenDatum.Content,
		CreateAt:  screenDatum.CreateAt.Unix(),
	}}, nil
}
