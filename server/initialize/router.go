package initialize

import (
	"Helios/common"
	"Helios/middleware"
	"Helios/router"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func Router() *gin.Engine {
	// 设置运行模式
	gin.SetMode(common.Config.System.Mode)

	// 创建一个没有中间件的 Gin 路由引擎
	r := gin.New()

	// 全局中间件
	r.Use(middleware.AccessLogger)
	r.Use(middleware.Cors)
	r.Use(middleware.Exception)

	// JWT 中间件
	auth, err := middleware.JWTAuth()
	if err != nil {
		common.SystemLog.Fatal("JWT 中间件初始化异常：" + err.Error())
	}

	// OPEN API 开放路由，无需认证
	router.OpenApiRouter(r.Group(common.SYSTEM_OPEN_API_PREFIX))

	// 注册登录路由，无需认证
	router.LoginRouter(r.Group(common.SYSTEM_OPEN_API_PREFIX), auth)

	// 需要认证的路由
	authGroup := r.Group(common.SYSTEM_API_PREFIX)
	authGroup.Use(auth.MiddlewareFunc())

	{
		// 其他认证路由
		router.OtherAuthRouter(authGroup, auth)
		// 系统配置路由，需要认证
		router.SystemSettingWithAuthRouter(authGroup.Group("/system/setting"), auth)
		// 系统菜单路由，需要认证
		router.SystemMenuWithAuthRouter(authGroup.Group("/system/menu"), auth)
		// 系统角色路由，需要认证
		router.SystemRoleWithAuthRouter(authGroup.Group("/system/role"), auth)
		// 系统 API 分类路由，需要认证
		router.SystemApiCategoryWithAuthRouter(authGroup.Group("/system/api-category"), auth)
		// 系统 API 路由，需要认证
		router.SystemApiWithAuthRouter(authGroup.Group("/system/api"), auth)
		// 用户路由，需要认证
		router.UserWithAuthRouter(authGroup.Group("/user"), auth)
	}

	return r
}
