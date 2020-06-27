package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Success       = 0
	ParamError    = 1
	BusinessError = 2
	MySqlError    = 3
	SystemError   = 4
	AuthError	  = 1001
)

// JsonResponse 有错误发生时，发送错误JSON
func JsonResponse(code int, msg string, args ...interface{}) {
	if len(args) == 0 {
		panic("缺少 *gin.Context")
	}
	var c *gin.Context
	var errNo = code
	if len(args) == 1 {
		theCtx, ok := args[0].(*gin.Context)
		if !ok {
			// Todo: 内部调用实际上不会出现这种情况的
			panic("缺少 *gin.Context")
		}
		c = theCtx
	}
	c.JSON(http.StatusOK, gin.H{
		"code": errNo,
		"msg":  msg,
		"data": nil,
	})
	// 终止请求链
	c.Abort()
}
