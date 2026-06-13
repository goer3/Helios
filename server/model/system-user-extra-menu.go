package model

// 用户附加授权菜单模型
type SystemUserExtraMenu struct {
	BaseModel
	GrantType    uint        `gorm:"type:tinyint unsigned;not null;default:1;index:idx_system_user_extra_menu_user_grant_menu;comment:授权类型（1：额外授权，2：取消授权）" json:"grant_type"`
	SystemUserId uint        `gorm:"not null;uniqueIndex:uk_system_user_extra_menu_user_menu,priority:1;comment:用户ID" json:"system_user_id"`
	SystemUser   *SystemUser `gorm:"foreignKey:SystemUserId;references:Id" json:"system_user,omitempty"`
	SystemMenuId uint        `gorm:"not null;uniqueIndex:uk_system_user_extra_menu_user_menu,priority:2;comment:菜单ID" json:"system_menu_id"`
	SystemMenu   *SystemMenu `gorm:"foreignKey:SystemMenuId;references:Id" json:"system_menu,omitempty"`
}

// 设置表名
func (SystemUserExtraMenu) TableName() string {
	return "system_user_extra_menu"
}
