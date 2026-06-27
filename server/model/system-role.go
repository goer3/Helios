package model

// 角色模型
type SystemRole struct {
	BaseModel
	Name        string       `gorm:"type:varchar(64);not null;uniqueIndex:uk_system_role_name;comment:角色名称" json:"name"`
	Description string       `gorm:"type:varchar(255);not null;default:'';comment:角色描述" json:"description"`
	Status      uint         `gorm:"type:tinyint unsigned;not null;default:1;comment:角色状态（0：禁用，1：启用）" json:"status"`
	Users       []User       `gorm:"foreignKey:SystemRoleId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"users,omitempty"`
	SystemMenus []SystemMenu `gorm:"many2many:system_role_menu;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"system_menus,omitempty"`
}

// 设置表名
func (SystemRole) TableName() string {
	return "system_role"
}
