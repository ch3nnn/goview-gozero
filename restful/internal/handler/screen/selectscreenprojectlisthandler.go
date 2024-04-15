package screen

import (
	"net/http"

	xhttp "github.com/ch3nnn/goview-gozero/pkg/http/x"

	"github.com/ch3nnn/goview-gozero/restful/internal/logic/screen"
	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// SelectScreenProjectListHandler 大屏信息列表
func SelectScreenProjectListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SelectScreenProjectListReq
		if err := xhttp.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := screen.NewSelectScreenProjectListLogic(r.Context(), svcCtx)
		resp, err := l.SelectScreenProjectList(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
