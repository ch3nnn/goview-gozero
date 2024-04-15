package screenlogic

import (
	"context"
	"errors"

	"github.com/ch3nnn/goview-gozero/common/bizerr"
	"github.com/ch3nnn/goview-gozero/service/screen/internal/dal/query"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

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
	project, err := query.ScreenProject.WithContext(l.ctx).
		Where(query.ScreenProject.ID.Eq(in.GetId())).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, bizerr.ScreenRpcProjectNotFound
		}
		return nil, err
	}

	return &pb.SelectScreenProjectByIdResp{ScreenProject: &pb.ScreenProject{
		Id:       project.ID,
		Name:     project.Name,
		State:    project.State,
		UserId:   project.UserID,
		IndexImg: project.IndexImg,
		Remark:   project.Remark,
		IsDel:    project.IsDel == 1,
		CreateAt: project.CreateAt.Unix(),
	}}, nil
}
