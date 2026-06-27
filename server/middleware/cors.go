package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 跨域配置
func Cors(ctx *gin.Context) {
	method := ctx.Request.Method
	// 允许的源（生产环境建议将 "*" 改为前端确切的域名，例如 "https://example.com"）
	ctx.Header("Access-Control-Allow-Origin", "*")
	// 允许的请求头字段
	ctx.Header("Access-Control-Allow-Headers", "Content-Type, Access-Token, Authorization, Token, X-HELIOS-NAME, X-HELIOS-REQUEST-ID")
	// 允许的请求方法
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	// 允许前端读取跨域响应中的其他头部（非必须）
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	// 凭证共享（若前端 axios 请求需要携带 cookie，则此项必须为 true）
	ctx.Header("Access-Control-Allow-Credentials", "true")
	// 放行所有 OPTIONS 预检请求
	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
	ctx.Next()
}
