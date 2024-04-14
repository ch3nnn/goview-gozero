/**
 * @Author: chentong
 * @Date: 2023/10/30 15:30
 */

package http

import (
	"net/http"

	"google.golang.org/grpc/status"
)

// CustomErr 自定义错误码
const CustomErr = 499

type ErrorMsg struct {
	Detail string `json:"detail"` // 错误详情
}

func NewErrorMsg(detail string) *ErrorMsg {
	return &ErrorMsg{Detail: detail}
}

func NewCustomErrHandler(v any) (int, any) {
	switch data := v.(type) {
	case error:
		return CustomErr, NewErrorMsg(data.Error())
	case *status.Status:
		return CustomErr, NewErrorMsg(data.Message())
	default:
		return http.StatusOK, data
	}
}
