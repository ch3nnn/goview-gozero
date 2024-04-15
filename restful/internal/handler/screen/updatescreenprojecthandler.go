package screen

import (
	"net/http"

	xhttp "github.com/ch3nnn/goview-gozero/pkg/http/x"

	"github.com/ch3nnn/goview-gozero/restful/internal/logic/screen"
	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// UpdateScreenProjectHandler 更新大屏信息
func UpdateScreenProjectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateScreenProjectReq
		if err := xhttp.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := screen.NewUpdateScreenProjectLogic(r.Context(), svcCtx)
		resp, err := l.UpdateScreenProject(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
