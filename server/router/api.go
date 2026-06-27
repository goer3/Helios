package router

import (
	"Helios/api"

	jwt "github.com/appleboy/gin-jwt/v3"
	"github.com/gin-gonic/gin"
)

// 其他路由，需要认证
func OtherAuthRouter(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/token-verify", api.TokenVerifyHandler) // 验证 token
	return rg
}

// 系统配置路由，需要认证
func SystemSettingWithAuthRouter(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/list", api.SystemSettingListHandler) // 获取系统配置列表
	return rg
}

// 角色路由，需要认证
func SystemRoleWithAuthRouter(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/list", api.SystemRoleListHandler)      // 获取角色列表
	rg.POST("/create", api.SystemRoleCreateHandler) // 创建角色
	return rg
}

// 菜单路由，需要认证
func SystemMenuWithAuthRouter(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/list", api.SystemMenuListHandler)      // 获取菜单列表
	rg.POST("/create", api.SystemMenuCreateHandler) // 创建菜单
	return rg
}

// API 分类路由，需要认证
func SystemApiCategoryWithAuthRouter(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/list", api.SystemApiCategoryListHandler)      // 获取API分类列表
	rg.POST("/create", api.SystemApiCategoryCreateHandler) // 创建API分类
	return rg
}

// API 路由，需要认证
func SystemApiWithAuthRouter(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/list", api.SystemApiListHandler)      // 获取API列表
	rg.POST("/create", api.SystemApiCreateHandler) // 创建API
	return rg
}

// 用户路由，需要认证
func UserWithAuthRouter(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/list", api.UserListHandler)      // 获取用户列表
	rg.POST("/create", api.UserCreateHandler) // 创建用户
	return rg
}
