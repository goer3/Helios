package initialize

import (
	"Helios/common"
	"Helios/middleware"
	"Helios/router"
	"path"

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

	// 系统路由组
	{
		// 系统路由前缀
		systemRouterPrefix := path.Join(common.SYSTEM_API_PREFIX, "system")

		// 注册系统配置路由，需要认证
		systemSettingWithAuthRouterGroup := r.Group(path.Join(systemRouterPrefix, "setting"))
		systemSettingWithAuthRouterGroup.Use(auth.MiddlewareFunc())
		router.SystemSettingWithAuthRouter(systemSettingWithAuthRouterGroup, auth)

		// 注册菜单路由，需要认证
		systemMenuWithAuthRouterGroup := r.Group(path.Join(systemRouterPrefix, "menu"))
		systemMenuWithAuthRouterGroup.Use(auth.MiddlewareFunc())
		router.SystemMenuWithAuthRouter(systemMenuWithAuthRouterGroup, auth)

		// 注册角色路由，需要认证
		systemRoleWithAuthRouterGroup := r.Group(path.Join(systemRouterPrefix, "role"))
		systemRoleWithAuthRouterGroup.Use(auth.MiddlewareFunc())
		router.SystemRoleWithAuthRouter(systemRoleWithAuthRouterGroup, auth)

		// 注册用户路由，需要认证
		systemUserWithAuthRouterGroup := r.Group(path.Join(systemRouterPrefix, "user"))
		systemUserWithAuthRouterGroup.Use(auth.MiddlewareFunc())
		router.SystemUserWithAuthRouter(systemUserWithAuthRouterGroup, auth)

		// 注册API分类路由，需要认证
		systemApiCategoryWithAuthRouterGroup := r.Group(path.Join(systemRouterPrefix, "api-category"))
		systemApiCategoryWithAuthRouterGroup.Use(auth.MiddlewareFunc())
		router.SystemApiCategoryWithAuthRouter(systemApiCategoryWithAuthRouterGroup, auth)

		// 注册API路由，需要认证
		systemApiWithAuthRouterGroup := r.Group(path.Join(systemRouterPrefix, "api"))
		systemApiWithAuthRouterGroup.Use(auth.MiddlewareFunc())
		router.SystemApiWithAuthRouter(systemApiWithAuthRouterGroup, auth)
	}

	return r
}
