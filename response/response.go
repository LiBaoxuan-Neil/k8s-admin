package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS         = 200 // 成功
	ERROR_COMMON    = 500 // 通用错误
	ERROR_PARAM     = 400 // 参数错误
	ERROR_AUTH      = 401 // 认证错误
	ERROR_FORBID    = 403 // 权限不足
	ERROR_NOT_FOUND = 404 // 资源不存在
)

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  "success",
	})
}
func SuccessWithMessage(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  msg,
	})
}
func SuccessWithData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  "success",
		"data": data,
	})
}
func SuccessWithDetailed(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  msg,
		"data": data,
	})
}

func Fail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": ERROR_COMMON,
		"msg":  "fail",
	})
}
func FailWithMessage(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": ERROR_COMMON,
		"msg":  msg,
	})
}
func FailWithData(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": ERROR_COMMON,
		"msg":  msg,
		"data": data,
	})
}
func FailWithDetailed(c *gin.Context, code int, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// FailWithCode 使用指定错误码返回失败响应
func FailWithCode(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
