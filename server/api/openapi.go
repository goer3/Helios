package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health 健康检查接口
func HealthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}
