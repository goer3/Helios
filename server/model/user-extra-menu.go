package model

// 用户附加授权菜单模型
type UserExtraMenu struct {
	BaseModel
	GrantType    uint        `gorm:"type:tinyint unsigned;not null;default:1;index:idx_user_extra_menu_grant_type;comment:授权类型（1：额外授权，2：取消授权）" json:"grant_type"`
	UserId       uint        `gorm:"not null;uniqueIndex:uk_user_extra_menu_user_menu,priority:1;comment:用户ID" json:"user_id"`
	User         *User       `gorm:"foreignKey:UserId;references:Id" json:"user,omitempty"`
	SystemMenuId uint        `gorm:"not null;uniqueIndex:uk_user_extra_menu_user_menu,priority:2;comment:菜单ID" json:"system_menu_id"`
	SystemMenu   *SystemMenu `gorm:"foreignKey:SystemMenuId;references:Id" json:"system_menu,omitempty"`
}

// 设置表名
func (UserExtraMenu) TableName() string {
	return "user_extra_menu"
}
