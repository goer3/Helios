package api

import (
	"Helios/common"
	"Helios/dto"
	"Helios/model"
	"Helios/pkg/response"
	"Helios/service"

	"github.com/gin-gonic/gin"
)

// 系统API列表接口
func SystemApiListHandler(ctx *gin.Context) {
	var title string = "获取系统API列表"
	var list []model.SystemApi
	if err := common.DB.Find(&list).Error; err != nil {
		common.SystemLog.Error(title+"失败: ", err)
		response.FailureWithMessage(title + "失败:" + err.Error())
		return
	}
	response.SuccessWithData(map[string]any{
		"list": list,
	})
}

// 创建系统API接口
func SystemApiCreateHandler(ctx *gin.Context) {
	var title string = "创建系统API"
	var req dto.SystemApiCreateRequest
	if ok, msg := BindAndValidate(ctx, &req); !ok {
		response.FailureWithMessage(msg)
		return
	}
	if err := service.CreateFromDto[model.SystemApi](common.DB, &req); err != nil {
		response.FailureWithMessage(title + "失败：" + err.Error())
		return
	}
	response.SuccessWithMessage(title + "成功")
}
