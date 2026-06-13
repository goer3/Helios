package api

import (
	"Helios/common"
	"Helios/model"
	"Helios/pkg/response"

	"github.com/gin-gonic/gin"
)

// 获取菜单列表
func SystemMenuListHandler(ctx *gin.Context) {
	var systemMenuList []model.SystemMenu
	err := common.DB.Find(&systemMenuList).Error
	if err != nil {
		common.SystemLog.Error("获取菜单列表失败: ", err)
		response.FailureWithMessage("获取菜单列表失败")
		return
	}
	response.SuccessWithData(map[string]any{
		"list": systemMenuList,
	})
}
