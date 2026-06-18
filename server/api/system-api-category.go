package api

import (
	"Helios/common"
	"Helios/dto"
	"Helios/model"
	"Helios/pkg/response"
	"Helios/service"

	"github.com/gin-gonic/gin"
)

// 系统API分类列表接口
func SystemApiCategoryListHandler(ctx *gin.Context) {
	var title string = "获取系统API分类列表"
	var list []model.SystemApiCategory
	if err := common.DB.Preload("SystemApis").Find(&list).Error; err != nil {
		common.SystemLog.Error(title+"失败: ", err)
		response.FailureWithMessage(title + "失败:" + err.Error())
		return
	}
	response.SuccessWithData(map[string]any{
		"list": list,
	})
}

// 创建系统API分类接口
func SystemApiCategoryCreateHandler(ctx *gin.Context) {
	var title string = "创建系统API分类"
	var req dto.SystemApiCategoryCreateRequest
	if ok, msg := BindAndValidate(ctx, &req); !ok {
		response.FailureWithMessage(msg)
		return
	}
	if err := service.CreateFromDto[model.SystemApiCategory](common.DB, &req); err != nil {
		response.FailureWithMessage(title + "失败：" + err.Error())
		return
	}
	response.SuccessWithMessage(title + "成功")
}
