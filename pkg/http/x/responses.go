package x

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// OkXml writes v into w with 200 OK.
func OkXml(w http.ResponseWriter, v any) {
	WriteXml(w, http.StatusOK, v)
}

// OkXmlCtx writes v into w with 200 OK.
func OkXmlCtx(ctx context.Context, w http.ResponseWriter, v any) {
	WriteXmlCtx(ctx, w, http.StatusOK, v)
}

// WriteXml writes v as xml string into w with code.
func WriteXml(w http.ResponseWriter, code int, v any) {
	if err := doWriteXml(w, code, v); err != nil {
		logx.Error(err)
	}
}

// WriteXmlCtx writes v as xml string into w with code.
func WriteXmlCtx(ctx context.Context, w http.ResponseWriter, code int, v any) {
	if err := doWriteXml(w, code, v); err != nil {
		logx.WithContext(ctx).Error(err)
	}
}

func doWriteXml(w http.ResponseWriter, code int, v any) error {
	bs, err := xml.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return fmt.Errorf("marshal xml failed, error: %w", err)
	}

	w.Header().Set(httpx.ContentType, XmlContentType)
	w.WriteHeader(code)

	if n, err := w.Write(bs); err != nil {
		// http.ErrHandlerTimeout has been handled by http.TimeoutHandler,
		// so it's ignored here.
		if !errors.Is(err, http.ErrHandlerTimeout) {
			return fmt.Errorf("write response failed, error: %w", err)
		}
	} else if n < len(bs) {
		return fmt.Errorf("actual bytes: %d, written bytes: %d", len(bs), n)
	}

	return nil
}
