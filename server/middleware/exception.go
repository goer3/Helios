package middleware

import (
	"Helios/common"
	"Helios/pkg/response"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// Exception 统一异常处理中间件。
// 正常情况下，handler 通过 panic(response.Response{}) 返回响应内容；
// 真正异常导致的 panic 会被捕获并返回 500，防止程序退出。
func Exception(ctx *gin.Context) {
	defer func() {
		err := recover()
		if err == nil {
			return
		}

		// 受控响应：handler 通过 panic(response.Response{}) 返回正常响应
		if resp, ok := err.(response.Response); ok {
			ctx.JSON(http.StatusOK, resp)
			ctx.Abort()
			return
		}

		// 真正异常：记录 stack trace 并返回 500
		stack := debug.Stack()
		if common.SystemLog != nil {
			common.SystemLog.Errorf("Panic: %v", err)
			common.SystemLog.Error(string(stack))
			_ = common.SystemLog.Sync()
		} else {
			fmt.Printf("Panic: %v\n%s", err, stack)
		}

		ctx.JSON(http.StatusOK, response.Response{
			Code:    response.InternalServerError,
			Message: response.ResponseMessage[response.InternalServerError],
			Data:    response.GetEmpty(),
		})
		ctx.Abort()
	}()
	ctx.Next()
}
