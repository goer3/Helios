package router

import (
	"Helios/api"

	jwt "github.com/appleboy/gin-jwt/v3"
	"github.com/gin-gonic/gin"
)

// 开放路由，无需认证
func OpenApiRouter(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/health", api.HealthHandler) // 健康检查接口
	return rg
}

// 登录路由，无需认证
func LoginRouter(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.POST("/login", auth.LoginHandler)          // 账号密码登录
	rg.POST("/login/dingtalk", auth.LoginHandler) // 钉钉扫码登录
	rg.POST("/login/feishu", auth.LoginHandler)   // 飞书扫码登录
	rg.POST("/login/wechat", auth.LoginHandler)   // 企微扫码登录
	return rg
}
