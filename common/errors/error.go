/*
 * 统一错误处理
 */
package errors

import (
	"github.com/NothingXiang/go-study/common/util"
)

type APIError struct {
	Code    string
	Message string
	Key     string
}

func newAPIError(code string, key string) *APIError {
	return &APIError{Code: code, Key: key}
}

func (A APIError) Error() string {
	if util.IsEmpty(A.Message) {
		return A.Message
	}
	return A.Key + ":" + A.Message
}

func (A *APIError) SetMsg(err error) *APIError {
	A.Message = err.Error()
	return A
}

var (
	Unknown      = newAPIError("10001", "unknown errors")
	NotExist     = newAPIError("10002", "not exist")
	InvalidParam = newAPIError("10003", "invalid param")
	ParamFmt     = newAPIError("10004", "param format errors")
	// another Error
)
