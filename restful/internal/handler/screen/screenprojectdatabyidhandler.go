package screen

import (
	"net/http"

	xhttp "github.com/ch3nnn/goview-gozero/pkg/http/x"

	"github.com/ch3nnn/goview-gozero/restful/internal/logic/screen"
	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// ScreenProjectDataByIDHandler 根据大屏信息ID获取大屏数据
func ScreenProjectDataByIDHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ScreenProjectDataByIDReq
		if err := xhttp.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := screen.NewScreenProjectDataByIDLogic(r.Context(), svcCtx)
		resp, err := l.ScreenProjectDataByID(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
