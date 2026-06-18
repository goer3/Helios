package dto

// 系统API创建请求
type SystemApiCreateRequest struct {
	Name                string `json:"name" binding:"required,min=2,max=50" msg:"名称不能为空，长度必须在2-50之间"`
	Method              string `json:"method" binding:"required,min=2,max=50" msg:"方法不能为空，长度必须在2-50之间"`
	Api                 string `json:"api" binding:"required,min=2,max=100" msg:"路径不能为空，长度必须在2-100之间"`
	IsAuthApi           *uint  `json:"is_auth_api" binding:"required,oneof=0 1" msg:"是否需要认证字段不能为空，且只能为 0（不需要认证）或 1（需要认证）"`
	SystemApiCategoryId *uint  `json:"system_api_category_id" binding:"required" msg:"系统分类ID不能为空"`
}
