package errcode

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/duke-git/lancet/v2/convertor"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 业务状态码与消息
//
//	业务错误码建议恒大于等于 0，与 grpc code 兼容
//	业务状态码等于 0 时，代表成功
//	业务状态码大于 0 时，代表错误
//	自定义业务状态码请大于等于 17，与 grpc 默认 code 区分开
//	99 代表意外错误，100 代表请求参数错误，均已被占用，建议之后业务状态码从 200 开始
const (
	// CodeOK 请求成功业务状态码
	CodeOK = 0
	// MsgOK 请求成功消息
	MsgOK = "ok"

	// CodeUnexpected 意外错误业务状态码
	CodeUnexpected = 99
	// MsgUnexpected 意外错误业务消息
	MsgUnexpected = "服务器繁忙，请稍后重试"

	// CodeInvalidParams 请求参数错误业务状态码
	CodeInvalidParams = 100
	// MsgInvalidParams 请求参数错误业务消息
	MsgInvalidParams = "请求参数错误"

	// GrpcMaxCode grpc 最大错误码
	GrpcMaxCode = 17
)

var descRegex = regexp.MustCompile(`code: (\d+), msg: (.+), http code: (\d+)`)

// 业务错误
var (
	// NoErr 无错误
	NoErr = New(CodeOK, MsgOK)
	// ErrUnexpected 意外错误
	ErrUnexpected = New(CodeUnexpected, MsgUnexpected)
	// ErrInvalidParams 请求参数错误
	ErrInvalidParams = New(CodeInvalidParams, MsgInvalidParams)
)

// Err 业务错误结构详情
type Err struct {
	Code     int
	HttpCode int
	Msg      string
}

// String Err 实现 String 方法
func (e *Err) String() string {
	return fmt.Sprintf("code: %d, msg: %s, http code: %d", e.Code, e.Msg, e.HttpCode)
}

// Error Err 实现 Error 方法
func (e *Err) Error() string {
	return e.Msg
}

// Format Err 实现 Format 方法
func (e *Err) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, e.String())
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.Msg)
	case 'q':
		fmt.Fprintf(s, "%q", e.Msg)
	}
}

// GRPCStatus Err 实现 GRPCStatus 方法
func (e *Err) GRPCStatus() *status.Status {
	return status.New(codes.Code(e.Code), e.String())
}

// New 新建业务错误
func New(code int, msg string, httpCode ...int) error {
	hc := http.StatusOK
	if len(httpCode) != 0 && httpCode[0] > 0 {
		hc = httpCode[0]
	}

	return &Err{Code: code, HttpCode: hc, Msg: msg}
}

// FromMessage 解析业务消息对应的业务错误
func FromMessage(msg string) (*Err, bool) {
	gs := descRegex.FindStringSubmatch(msg)
	if len(gs) == 4 {
		code, _ := convertor.ToInt(gs[1])
		httpCode, _ := convertor.ToInt(gs[3])
		return &Err{
			Code:     int(code),
			Msg:      gs[2],
			HttpCode: int(httpCode),
		}, true
	}

	var e *Err
	errors.As(ErrUnexpected, &e)

	return e, false
}

// FromError 解析错误对应的业务错误
func FromError(err error) (*Err, bool) {
	var e *Err
	if err == nil {
		errors.As(NoErr, &e)
		return e, true
	}

	if ok := errors.As(err, &e); ok {
		return e, true
	}

	if s, ok := status.FromError(err); ok {
		code, msg := s.Code(), s.Message()
		if code > GrpcMaxCode {
			if e, ok := FromMessage(msg); ok {
				return e, true
			}
		}
	}

	errors.As(ErrUnexpected, &e)

	return e, false
}

// IsErr 判断是否为业务错误
func IsErr(err error) bool {
	_, isErr := FromError(err)
	return isErr
}

// Is 判断业务错误是否相同
//
// target 须为业务错误，用于和 err 进行比对
func Is(err, target error) bool {
	te, ok := FromError(target)
	if !ok {
		return false
	}

	pe, ok := FromError(err)
	if !ok {
		return false
	}

	if pe.Code == te.Code && pe.Msg == te.Msg && pe.HttpCode == te.HttpCode {
		return true
	}

	return false
}
