// Code generated by goctl. DO NOT EDIT.
package types

type UserLogoutRep struct{}

type UserLogoutResp struct{}

type OssInfoRep struct{}

type OssInfoResp struct {
	BucketURL string `json:"bucketURL"`
}

type UserLoginRep struct {
	Username string `json:"username" validate:"required,gte=6,lte=16"` // 用户名
	Password string `json:"password" validate:"required"`              // 密码
}

type UserInfo struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	DepId    int64  `json:"depId"`
	PosId    string `json:"posId"`
	DepName  string `json:"depName"`
	PosName  string `json:"posName"`
}

type Token struct {
	TokenName            string            `json:"tokenName"`
	TokenValue           string            `json:"tokenValue"`
	IsLogin              bool              `json:"isLogin"`
	LoginId              string            `json:"loginId"`
	LoginType            string            `json:"loginType"`
	TokenTimeout         int               `json:"tokenTimeout"`
	SessionTimeout       int               `json:"sessionTimeout"`
	TokenSessionTimeout  int               `json:"tokenSessionTimeout"`
	TokenActivityTimeout int               `json:"tokenActivityTimeout"`
	LoginDevice          string            `json:"loginDevice"`
	Tag                  map[string]string `json:"tag"`
}

type UserLoginResp struct {
	UserInfo UserInfo `json:"userinfo"`
	Token    Token    `json:"token"`
}

type InsertScreenProjectDataReq struct {
	ProjectId int64  `form:"projectId"`
	Content   string `form:"content"`
}

type InsertScreenProjectDataResp struct{}

type ScreenProjectDataByIDReq struct {
	ProjectId int64 `form:"projectId"` // 大屏信息ID
}

type ScreenProjectDataByIDResp struct {
	Content      string `json:"content"`
	Id           string `json:"id"`
	Name         string `json:"projectName"`
	State        int64  `json:"state"`
	CreateAt     string `json:"createTime"`
	CreateUserId int64  `json:"createUserId"`
	IsDel        int64  `json:"isDelete"`
	IndexImage   string `json:"indexImage"`
	Remark       string `json:"remarks"`
}

type ScreenProject struct {
	Id       int64  `json:"id"`           //  ID
	Name     string `json:"projectName"`  //  大屏名称
	State    int64  `json:"state"`        //  发布状态(-1 未发布  1 已发布)
	UserId   int64  `json:"createUserId"` //  创建用户ID
	IndexImg string `json:"indexImage"`   //  缩略图
	Remark   string `json:"remarks"`      //  备注
	IsDel    bool   `json:"isDelete"`     //  是否删除(0 未删除 1 已删除)
	CreateAt string `json:"createTime"`   //  创建时间
}

type AddScreenProjectReq struct {
	Name     string `json:"projectName"`         //  大屏名称
	IndexImg string `json:"indexImage,optional"` //  缩略图
	Remark   string `json:"remarks,optional"`    //  备注
}

type AddScreenProjectResp struct {
	Id       int64  `json:"id"`           //  ID
	Name     string `json:"projectName"`  //  大屏名称
	State    int64  `json:"state"`        //  发布状态(-1 未发布  1 已发布)
	UserId   int64  `json:"createUserId"` //  创建用户ID
	IndexImg string `json:"indexImage"`   //  缩略图
	Remark   string `json:"remarks"`      //  备注
	IsDel    bool   `json:"isDelete"`     //  是否删除(0 未删除 1 已删除)
	CreateAt string `json:"createTime"`   //  创建时间
}

type UpdateScreenProjectReq struct {
	Id       string  `json:"id"`                   //  ID
	Name     *string `json:"projectName,optional"` //  大屏名称
	IndexImg *string `json:"indexImage,optional"`  //  缩略图
	Remark   *string `json:"remarks,optional"`     //  备注
}

type UpdateScreenProjectResp struct{}

type UpdateScreenProjectPublishReq struct {
	Id    int64 `json:"id"`    //  ID
	State int64 `json:"state"` //  发布状态(-1 未发布  1 已发布)
}

type UpdateScreenProjectPublishResp struct{}

type DelScreenProjectReq struct {
	Ids int64 `form:"ids"` //  ID
}

type DelScreenProjectResp struct{}

type SelectScreenProjectListReq struct {
	Page  int64 `form:"page"`  //  页码
	Limit int64 `form:"limit"` //  每页数量
}

type SelectScreenProjectListResp struct {
	Count   int64           `json:"count"`   //  总数
	Results []ScreenProject `json:"results"` //  screen_project
}

type UploadScreenProjectFileReq struct {
	Filename string `form:"filename,optional"` // 文件名
	Object   string `form:"object,optional"`   // 文件
}

type UploadScreenProjectFileResp struct {
	Filename string `json:"fileName"` // 文件名
}
