package sys

import (
	"net/http"

	xhttp "github.com/ch3nnn/goview-gozero/pkg/http/x"

	"github.com/ch3nnn/goview-gozero/restful/internal/logic/sys"
	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// OssInfoHandler 获取oss信息
func OssInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OssInfoRep
		if err := xhttp.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := sys.NewOssInfoLogic(r.Context(), svcCtx)
		resp, err := l.OssInfo(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
