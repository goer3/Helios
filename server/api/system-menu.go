package api

import (
	"Helios/common"
	"Helios/dto"
	"Helios/model"
	"Helios/pkg/response"
	"Helios/service"

	"github.com/gin-gonic/gin"
)

// 获取菜单列表
func SystemMenuListHandler(ctx *gin.Context) {
	var title string = "获取菜单列表"
	var list []model.SystemMenu
	err := common.DB.Find(&list).Error
	if err != nil {
		common.SystemLog.Error(title+"失败: ", err)
		response.FailureWithMessage(title + "失败:" + err.Error())
		return
	}
	response.SuccessWithData(map[string]any{
		"list": list,
	})
}

// 创建系统菜单
func SystemMenuCreateHandler(ctx *gin.Context) {
	var title string = "创建系统菜单"
	var req dto.SystemMenuCreateRequest
	if ok, msg := BindAndValidate(ctx, &req); !ok {
		response.FailureWithMessage(msg)
		return
	}
	if err := service.CreateFromDto[model.SystemMenu](common.DB, &req); err != nil {
		response.FailureWithMessage(title + "失败：" + err.Error())
		return
	}
	response.SuccessWithMessage(title + "成功")
}
