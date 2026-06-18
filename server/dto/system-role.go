package dto

// 角色创建请求
type SystemRoleCreateRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=50" msg:"角色名称不能为空，长度必须在2-50之间"`
	Description string `json:"description" binding:"required,min=2,max=255" msg:"角色描述不能为空，长度必须在2-255之间"`
	Status      *uint  `json:"status" binding:"required,oneof=0 1" msg:"状态字段不能为空，且只能为 0（禁用）或 1（启用）"`
}
