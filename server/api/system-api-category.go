package api

import (
	"Helios/common"
	"Helios/model"
	"Helios/pkg/response"

	"github.com/gin-gonic/gin"
)

// 系统API分类列表接口
func SystemApiCategoryListHandler(ctx *gin.Context) {
	var apiCategoryList []model.SystemApiCategory
	if err := common.DB.Find(&apiCategoryList).Error; err != nil {
		common.SystemLog.Error("查询系统API分类列表失败：", err.Error())
		response.FailureWithMessage("查询系统API分类列表失败")
		return
	}
	response.SuccessWithData(map[string]any{
		"list": apiCategoryList,
	})
}
