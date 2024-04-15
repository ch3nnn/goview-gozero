package types

import (
	"github.com/ch3nnn/goview-gozero/pkg/gorms/xdatatype"
	"gorm.io/plugin/soft_delete"
)

type ScreenProjectDO struct {
	ID       int64                  `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	Name     string                 `gorm:"column:name;type:varchar(255);comment:大屏名称" json:"name"`                // 大屏名称
	State    int64                  `gorm:"column:state;type:int;comment:发布状态(-1 未发布  1 已发布)" json:"state"`        // 发布状态(-1 未发布  1 已发布)
	UserID   int64                  `gorm:"column:user_id;type:int;comment:创建用户ID" json:"user_id"`                 // 创建用户ID
	IndexImg string                 `gorm:"column:index_img;type:varchar(255);comment:缩略图" json:"index_img"`       // 缩略图
	Remark   string                 `gorm:"column:remark;type:text;comment:备注" json:"remark"`                      // 备注
	IsDel    soft_delete.DeletedAt  `gorm:"column:is_del;softDelete:flag" json:"is_del"`                           //  是否删除(0 未删除 1 已删除)
	CreateAt *xdatatype.TimeToInt64 `gorm:"column:create_at;type:int unsigned;autoCreateTime" json:"create_at"`    // 创建时间
}