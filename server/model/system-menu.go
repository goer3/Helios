package model

// 菜单模型
type SystemMenu struct {
	Id          uint         `gorm:"primaryKey;comment:自增编号" json:"id"`
	Name        string       `gorm:"type:varchar(64);not null;default:'';uniqueIndex:uk_system_menu_name_path,priority:1;comment:菜单名称" json:"name"`
	Path        string       `gorm:"type:varchar(128);not null;uniqueIndex:uk_system_menu_name_path,priority:2;comment:菜单路径" json:"path"`
	Icon        string       `gorm:"type:varchar(128);not null;default:'';comment:菜单图标" json:"icon"`
	ParentId    uint         `gorm:"not null;index:idx_system_menu_parent_id;default:0;comment:父id" json:"parent_id"`
	Children    []SystemMenu `gorm:"-" json:"children,omitempty"`
	SystemRoles []SystemRole `gorm:"many2many:system_role_menu;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"system_roles,omitempty"`
}

// 表名设置
func (SystemMenu) TableName() string {
	return "system_menu"
}
