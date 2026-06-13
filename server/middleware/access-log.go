package middleware

import (
	"Helios/common"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 访问日志数据
type RequestLog struct {
	Method           string `json:"method"`
	RequestURI       string `json:"request_uri"`
	Status           int    `json:"status"`
	ExecMilliseconds int64  `json:"exec_milliseconds"`
	ClientIP         string `json:"client_ip"`
}

// 访问日志中间件
func AccessLogger(ctx *gin.Context) {
	startTime := time.Now()
	ctx.Next()

	// 收集日志数据
	d := RequestLog{
		Method:           ctx.Request.Method,
		RequestURI:       ctx.Request.RequestURI,
		Status:           ctx.Writer.Status(),
		ExecMilliseconds: time.Since(startTime).Milliseconds(),
		ClientIP:         ctx.ClientIP(),
	}

	// 确定日志级别（OPTIONS 请求使用 DEBUG）
	isDebug := d.Method == "OPTIONS"

	// 根据格式选择日志输出方式
	if common.Config.Log.Access.Formatter == "json" {
		if isDebug {
			common.AccessLog.Debugw("request", "data", d)
		} else {
			common.AccessLog.Infow("request", "data", d)
		}
	} else {
		msg := fmt.Sprintf("%s\t%s\t%d\t%dms\t%s",
			d.Method,
			d.RequestURI,
			d.Status,
			d.ExecMilliseconds,
			d.ClientIP,
		)
		if isDebug {
			common.AccessLog.Debug(msg)
		} else {
			common.AccessLog.Info(msg)
		}
	}
}
