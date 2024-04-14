package x

import (
	"context"
	"encoding/xml"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ch3nnn/goview-gozero/pkg/errcode"
)

// BaseResponse is the base response struct.
type BaseResponse[T any] struct {
	// TraceId represents the request trace id.
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty" example:"a1b2c3d4e5f6g7h8"` // 链路追踪id
	// Code represents the business code, not the http status code.
	Code int `json:"code" xml:"code" example:"0"`
	// Msg represents the business message, if Code = BusinessCodeOK,
	// and Msg is empty, then the Msg will be set to BusinessMsgOK.
	Msg string `json:"msg" xml:"msg" example:"ok"`
	// Data represents the business data.
	Data T `json:"data,omitempty" xml:"data,omitempty"`
}

type baseXmlResponse[T any] struct {
	XMLName  xml.Name `xml:"xml"`
	Version  string   `xml:"version,attr"`
	Encoding string   `xml:"encoding,attr"`
	BaseResponse[T]
}

// JsonBaseResponse writes v into w with http.StatusOK.
func JsonBaseResponse(w http.ResponseWriter, v any) {
	JsonBaseResponseCtx(context.Background(), w, v)
}

// JsonBaseResponseCtx writes v into w with http.StatusOK.
func JsonBaseResponseCtx(ctx context.Context, w http.ResponseWriter, v any) {
	code, resp := handleBaseResponseCtx(ctx, v)
	httpx.WriteJsonCtx(ctx, w, code, resp)
}

// XmlBaseResponse writes v into w with http.StatusOK.
func XmlBaseResponse(w http.ResponseWriter, v any) {
	XmlBaseResponseCtx(context.Background(), w, v)
}

// XmlBaseResponseCtx writes v into w with http.StatusOK.
func XmlBaseResponseCtx(ctx context.Context, w http.ResponseWriter, v any) {
	code, resp := handleXmlBaseResponseCtx(ctx, v)
	WriteXmlCtx(ctx, w, code, resp)
}

func handleXmlBaseResponseCtx(ctx context.Context, v any) (int, baseXmlResponse[any]) {
	code, base := handleBaseResponseCtx(ctx, v)
	return code, baseXmlResponse[any]{
		Version:      xmlVersion,
		Encoding:     xmlEncoding,
		BaseResponse: base,
	}
}

func handleBaseResponseCtx(ctx context.Context, v any) (int, BaseResponse[any]) {
	httpCode := http.StatusOK
	resp := BaseResponse[any]{TraceId: GetTraceId(ctx)}
	switch data := v.(type) {
	case *errors.CodeMsg:
		resp.Code = data.Code
		resp.Msg = data.Msg
	case errors.CodeMsg:
		resp.Code = data.Code
		resp.Msg = data.Msg
	case *errcode.Err:
		httpCode = data.HttpCode
		resp.Code = data.Code
		resp.Msg = data.Msg
	case errcode.Err:
		httpCode = data.HttpCode
		resp.Code = data.Code
		resp.Msg = data.Msg
	case *status.Status:
		// 非业务错误统一拦截返回意外错误，由日志记录此次实际错误
		resp.Code = int(data.Code())
		resp.Msg = errcode.MsgUnexpected
		if resp.Code == int(codes.OK) {
			resp.Msg = data.Message()
		} else if resp.Code < errcode.GrpcMaxCode {
			httpCode = errcode.CodeFromGrpcCode(data.Code())
		} else if e, ok := errcode.FromMessage(data.Message()); ok {
			// 尝试根据 grpc message 解析业务错误码
			httpCode = e.HttpCode
			resp.Code = e.Code
			resp.Msg = e.Msg
		}
	case interface{ GRPCStatus() *status.Status }:
		// 非业务错误统一拦截返回意外错误，由日志记录此次实际错误
		resp.Code = int(data.GRPCStatus().Code())
		resp.Msg = errcode.MsgUnexpected
		if resp.Code == int(codes.OK) {
			resp.Msg = data.GRPCStatus().Message()
		} else if resp.Code < errcode.GrpcMaxCode {
			httpCode = errcode.CodeFromGrpcCode(data.GRPCStatus().Code())
		} else if e, ok := errcode.FromMessage(data.GRPCStatus().Message()); ok {
			// 尝试根据 grpc message 解析业务错误码
			httpCode = e.HttpCode
			resp.Code = e.Code
			resp.Msg = e.Msg
		}
	case error:
		// resp.Code = BusinessCodeError
		// resp.Msg = data.Error()
		// 非业务错误统一拦截返回意外错误，由日志记录此次实际错误
		resp.Code = errcode.CodeUnexpected
		resp.Msg = errcode.MsgUnexpected
	default:
		resp.Code = BusinessCodeOK
		resp.Msg = BusinessMsgOK
		resp.Data = v
	}

	if resp.Code > BusinessCodeOK {
		logx.WithContext(ctx).Errorf("an error occurred in this request, err: %+v", v)
	}

	return httpCode, resp
}
