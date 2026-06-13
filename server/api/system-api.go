package api

import (
	"Helios/common"
	"Helios/model"
	"Helios/pkg/response"

	"github.com/gin-gonic/gin"
)

// 系统API列表接口
func SystemApiListHandler(ctx *gin.Context) {
	var apiList []model.SystemApi
	if err := common.DB.Find(&apiList).Error; err != nil {
		common.SystemLog.Error("查询系统API列表失败：", err.Error())
		response.FailureWithMessage("查询系统API列表失败")
		return
	}
	response.SuccessWithData(map[string]any{
		"list": apiList,
	})
}
