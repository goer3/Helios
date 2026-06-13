package api

import (
	"Helios/common"
	"Helios/dto"
	"Helios/model"
	"Helios/pkg/response"
	"Helios/pkg/utils"
	"Helios/service"

	"github.com/gin-gonic/gin"
)

// 获取用户列表
func SystemUserListHandler(ctx *gin.Context) {
	var userList []model.SystemUser

	// TODO: 实现分页和查询条件
	if err := common.DB.Preload("SystemRole").Find(&userList).Error; err != nil {
		common.SystemLog.Error("获取用户列表失败：", err.Error())
		response.FailureWithMessage("获取用户列表失败：" + err.Error())
		return
	}

	response.SuccessWithData(map[string]any{
		"list": userList,
	})
}

// 创建用户
func SystemUserCreateHandler(ctx *gin.Context) {
	var req dto.SystemUserCreateRequest
	// 验证请求参数，使用 binding 标签进行验证
	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := utils.GetValidateErrorMessage(err, &req)
		response.FailureWithMessage(msg)
		return
	}

	// 调用服务层创建用户
	if err := service.SystemUserCreate(&req); err != nil {
		response.FailureWithMessage(err.Error())
		return
	}

	response.SuccessWithMessage("用户创建成功")
}
