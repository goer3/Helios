package api

import (
	"Helios/common"
	"Helios/model"
	"Helios/pkg/response"

	"github.com/gin-gonic/gin"
)

// 获取角色列表
func SystemRoleListHandler(ctx *gin.Context) {
	var roleList []model.SystemRole
	if err := common.DB.Find(&roleList).Error; err != nil {
		common.SystemLog.Error("获取系统角色列表失败: ", err)
		response.FailureWithMessage("获取系统角色列表失败")
		return
	}
	response.SuccessWithData(map[string]any{
		"list": roleList,
	})
}
