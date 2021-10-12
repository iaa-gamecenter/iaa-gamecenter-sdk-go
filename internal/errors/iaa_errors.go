package errors

import (
	"fmt"
)

type AppError struct {
	Code    int32
	Message string
}

//对外输出Message
func (e *AppError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

//对外输出code
func (e *AppError) OutputCode() int32 {
	if e.Code > 0 {
		return -1
	}
	return e.Code
}

/**
新增message后续的详情
*/
func (e AppError) AppendData(data string) *AppError {
	e.Message = fmt.Sprintf(e.Message+"(%s)", data)
	return &e
}

var (
	ErrOverMaxRetryCount = &AppError{1000001, "客户端重试超过最大次数"}
)

func NewError(code int32, msg string) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
	}
}

//是否可以重试
func RetryAble(err error) bool {
	return false
}
