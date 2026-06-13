package model

// 接口模型
type SystemApi struct {
	Id                  uint               `gorm:"primaryKey;comment:自增编号" json:"id"`
	Name                string             `gorm:"type:varchar(64);not null;default:'';index:idx_system_api_name;comment:接口名称" json:"name"`
	Method              string             `gorm:"type:varchar(16);not null;default:'';uniqueIndex:uk_system_api_method_api,priority:1;comment:请求方法" json:"method"`
	Api                 string             `gorm:"type:varchar(255);not null;default:'';uniqueIndex:uk_system_api_method_api,priority:2;comment:接口URI" json:"api"`
	IsAuthApi           uint               `gorm:"type:tinyint unsigned;not null;default:0;comment:是否需要授权接口（0：否，1：是）" json:"is_auth_api"`
	SystemApiCategoryId uint               `gorm:"not null;index:idx_system_api_category_id;comment:接口分类ID" json:"system_api_category_id"`
	SystemApiCategory   *SystemApiCategory `gorm:"foreignKey:SystemApiCategoryId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"system_api_category,omitempty"`
}

// 表名设置
func (SystemApi) TableName() string {
	return "system_api"
}
