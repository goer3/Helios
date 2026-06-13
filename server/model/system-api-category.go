package model

// 接口分类模型
type SystemApiCategory struct {
	Id         uint                `gorm:"primaryKey;comment:自增编号" json:"id"`
	Name       string              `gorm:"type:varchar(64);not null;default:'';uniqueIndex:uk_system_api_category_name;comment:接口类型名称" json:"name"`
	ParentId   uint                `gorm:"not null;index:idx_system_api_category_parent_id;default:0;comment:父id" json:"parent_id"`
	Children   []SystemApiCategory `gorm:"-" json:"children,omitempty"`
	SystemApis []SystemApi         `gorm:"foreignKey:SystemApiCategoryId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"system_apis,omitempty"`
}

// 表名设置
func (SystemApiCategory) TableName() string {
	return "system_api_category"
}
