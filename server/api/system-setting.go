package api

import (
	"Helios/common"
	"Helios/model"
	"Helios/pkg/response"

	"github.com/gin-gonic/gin"
)

// 获取系统配置列表
func SystemSettingListHandler(ctx *gin.Context) {
	var title string = "获取系统配置列表"
	var list []model.SystemSetting
	if err := common.DB.Find(&list).Error; err != nil {
		common.SystemLog.Error(title+"失败: ", err)
		response.FailureWithMessage(title + "失败:" + err.Error())
		return
	}
	response.SuccessWithData(map[string]any{
		"list": list,
	})
}
