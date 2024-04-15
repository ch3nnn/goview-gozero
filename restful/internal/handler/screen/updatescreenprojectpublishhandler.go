package screen

import (
	"net/http"

	xhttp "github.com/ch3nnn/goview-gozero/pkg/http/x"

	"github.com/ch3nnn/goview-gozero/restful/internal/logic/screen"
	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// UpdateScreenProjectPublishHandler 更新大屏信息发布状态
func UpdateScreenProjectPublishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateScreenProjectPublishReq
		if err := xhttp.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := screen.NewUpdateScreenProjectPublishLogic(r.Context(), svcCtx)
		resp, err := l.UpdateScreenProjectPublish(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
