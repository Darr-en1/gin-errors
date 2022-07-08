package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

type customError struct {
	Code errCode `json:"code"` // 业务码
	Err  error
}

func (e *customError) Error() string {
	errMsg := ""
	if e.Err != nil {
		errMsg = e.Err.Error()
	}
	return fmt.Sprintf("code: %d, message: %s, error: %s ", e.Code, e.Code.String(), errMsg)
}

// 错误码
type errCode int64

// 定义errorCo
// 执行 go generate 生成 String 方法
//go:generate stringer -type errCode -linecomment
const (
	ServerError     errCode = 10101 // 服务器内部错误
	TooManyRequests errCode = 10102 // 请求太多
	ParamBindError  errCode = 10103 // 参数信息有误

	MySQLNoQueryError errCode = 20505 // 查询数据库无结果
	MySQLNoFieldError errCode = 20506 // 查询数据库字段不存在
)

func (i errCode) WrapWithMessage(err error, message string) error { // 为 customError 添加 stack 并附加信息
	return errors.Wrap(&customError{
		Code: i, Err: err,
	}, message)
}

func (i errCode) Wrap(err error) error { // 为customError 添加 stack
	return errors.Wrap(&customError{
		Code: i, Err: err,
	}, "")
}
