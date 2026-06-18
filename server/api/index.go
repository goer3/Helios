package api

import (
	"Helios/pkg/utils"

	"github.com/gin-gonic/gin"
)

// BindAndValidate 绑定 JSON 请求体并校验参数，校验失败时直接返回错误响应并终止请求。
// 用法：替代 handler 中重复的 ShouldBindJSON + GetValidateErrorMessage + response.FailureWithMessage 模式。
func BindAndValidate(ctx *gin.Context, req any) (ok bool, msg string) {
	if err := ctx.ShouldBindJSON(req); err != nil {
		msg := utils.GetValidateErrorMessage(err, req)
		return false, msg
	}
	return true, ""
}
