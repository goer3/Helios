package dto

// 系统菜单创建请求体
type SystemMenuCreateRequest struct {
	Id       uint   `json:"id" binding:"required" msg:"ID不能为空"`
	Name     string `json:"name" binding:"required,min=2,max=50" msg:"名称不能为空，长度必须在2-50之间"`
	Path     string `json:"path" binding:"required,min=2,max=50" msg:"路径不能为空，长度必须在2-50之间"`
	Icon     string `json:"icon" binding:"required,max=50" msg:"图标长度必须在0-50之间"`
	ParentId *uint  `json:"parent_id" binding:"required" msg:"父菜单ID不能为空"`
}
