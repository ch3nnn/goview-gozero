package x

import (
	"github.com/ch3nnn/goview-gozero/pkg/errcode"
)

const (
	xmlVersion  = "1.0"
	xmlEncoding = "UTF-8"

	// BusinessCodeOK represents the business code for success.
	BusinessCodeOK = errcode.CodeOK
	// BusinessMsgOK represents the business message for success.
	BusinessMsgOK = errcode.MsgOK

	// BusinessCodeError represents the business code for error.
	// Any code greater than 0 is a business error.
	BusinessCodeError = 1

	// XmlContentType represents the content type for xml.
	XmlContentType = "application/xml"
)
