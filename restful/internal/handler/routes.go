// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"

	screen "github.com/ch3nnn/goview-gozero/restful/internal/handler/screen"
	sys "github.com/ch3nnn/goview-gozero/restful/internal/handler/sys"
	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 用户登录
				Method:  http.MethodPost,
				Path:    "/sys/login",
				Handler: sys.UserLoginHandler(serverCtx),
			},
			{
				// 用户登出
				Method:  http.MethodGet,
				Path:    "/sys/logout",
				Handler: sys.UserLogoutHandler(serverCtx),
			},
			{
				// 获取oss信息
				Method:  http.MethodGet,
				Path:    "/sys/getOssInfo",
				Handler: sys.OssInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/goview"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 创建大屏信息
				Method:  http.MethodPost,
				Path:    "/project/create",
				Handler: screen.InsertScreenProjectHandler(serverCtx),
			},
			{
				// 更新大屏信息
				Method:  http.MethodPost,
				Path:    "/project/edit",
				Handler: screen.UpdateScreenProjectHandler(serverCtx),
			},
			{
				// 更新大屏信息发布状态
				Method:  http.MethodPut,
				Path:    "/project/publish",
				Handler: screen.UpdateScreenProjectPublishHandler(serverCtx),
			},
			{
				// 根据大屏信息ID删除
				Method:  http.MethodDelete,
				Path:    "/project/delete",
				Handler: screen.DeleteScreenProjectHandler(serverCtx),
			},
			{
				// 大屏信息列表
				Method:  http.MethodGet,
				Path:    "/project/list",
				Handler: screen.SelectScreenProjectListHandler(serverCtx),
			},
			{
				// 文件上传
				Method:  http.MethodPost,
				Path:    "/project/upload",
				Handler: screen.UploadScreenProjectFileHandler(serverCtx),
			},
			{
				// 根据大屏信息ID获取大屏数据
				Method:  http.MethodGet,
				Path:    "/project/getData",
				Handler: screen.ScreenProjectDataByIDHandler(serverCtx),
			},
			{
				// 保存大屏数据
				Method:  http.MethodPost,
				Path:    "/project/save/data",
				Handler: screen.InsertScreenProjectDataHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/goview"),
	)
}
