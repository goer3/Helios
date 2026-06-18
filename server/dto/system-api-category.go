package dto

// 系统API分类创建请求
type SystemApiCategoryCreateRequest struct {
	Id       uint   `json:"id" binding:"required" msg:"ID不能为空"`
	Name     string `json:"name" binding:"required,min=2,max=50" msg:"名称不能为空，长度必须在2-50之间"`
	ParentId *uint  `json:"parent_id" binding:"required" msg:"父分类ID不能为空"`
}
