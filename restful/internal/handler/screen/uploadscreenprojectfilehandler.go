package screen

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/zeromicro/go-zero/core/utils"

	xhttp "github.com/ch3nnn/goview-gozero/pkg/http/x"

	"github.com/ch3nnn/goview-gozero/restful/internal/logic/screen"
	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
	"github.com/ch3nnn/goview-gozero/restful/internal/types"
)

// UploadScreenProjectFileHandler 文件上传
func UploadScreenProjectFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadScreenProjectFileReq
		if err := xhttp.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		// 获取上传的文件（FormData）
		file, _, err := r.FormFile("object")
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}
		defer file.Close()

		filename := utils.NewUuid() + ".png"
		out, err := os.Create(path.Join("static", filename))
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}
		defer out.Close()

		if _, err = io.Copy(out, file); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		req.Filename = filename

		l := screen.NewUploadScreenProjectFileLogic(r.Context(), svcCtx)
		resp, err := l.UploadScreenProjectFile(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
