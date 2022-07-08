package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Data      interface{} `json:"data,omitempty"`     // 返回结果
	ErrMsg    string      `json:"err_msg,omitempty"`  // 错误信息
	ErrCode   errCode     `json:"err_code,omitempty"` // 错误状态码
	IsSuccess bool        `json:"is_success"`         // 请求结果
}

const (
	Fail    = false
	SUCCESS = true
)

// Result 通用返回格式
func Result(httpCode int, data interface{}, errCode errCode, ErrMsg string, isSuccess bool, c *gin.Context) {
	c.JSON(httpCode, Response{
		data,
		ErrMsg,
		errCode,
		isSuccess,
	})
}

// ResultWithErrCode 通用返回格式  携带errCode
func ResultWithErrCode(httpCode int, data interface{}, errCode errCode, isSuccess bool, c *gin.Context) {
	Result(httpCode, data, errCode, errCode.String(), isSuccess, c)
}

// ParamBindErrorResult 参数绑定异常返回
func ParamBindErrorResult(ErrMsg string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, Response{
		ErrMsg:  ErrMsg,
		ErrCode: ParamBindError,
	})
}

func ResultOk(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Data:      data,
		IsSuccess: SUCCESS,
	})
}

func ResultFail(httpCode int, errCode errCode, c *gin.Context) {
	c.JSON(httpCode, Response{
		ErrMsg:    errCode.String(),
		ErrCode:   errCode,
		IsSuccess: Fail,
	})
}
