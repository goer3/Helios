package api

import (
	"Helios/common"
	"Helios/dto"
	"Helios/model"
	"Helios/pkg/response"
	"Helios/service"

	"github.com/gin-gonic/gin"
)

// 获取用户列表
func UserListHandler(ctx *gin.Context) {
	var title string = "获取用户列表"
	var list []model.User

	// TODO: 实现分页和查询条件
	if err := common.DB.Preload("SystemRole").Find(&list).Error; err != nil {
		common.SystemLog.Error(title+"失败: ", err)
		response.FailureWithMessage(title + "失败:" + err.Error())
		return
	}
	response.SuccessWithData(map[string]any{
		"list": list,
	})
}

// 创建用户
func UserCreateHandler(ctx *gin.Context) {
	var title string = "创建用户"
	var req dto.UserCreateRequest
	if ok, msg := BindAndValidate(ctx, &req); !ok {
		response.FailureWithMessage(msg)
		return
	}
	if err := service.UserCreate(&req); err != nil {
		response.FailureWithMessage(err.Error())
		return
	}
	response.SuccessWithMessage(title + "成功")
}
