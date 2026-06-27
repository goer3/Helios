package model

// 用户附加授权接口模型
type UserExtraApi struct {
	BaseModel
	GrantType   uint       `gorm:"type:tinyint unsigned;not null;default:1;index:idx_user_extra_api_grant_type;comment:授权类型（1：额外授权，2：取消授权）" json:"grant_type"`
	UserId      uint       `gorm:"not null;uniqueIndex:uk_user_extra_api_user_api,priority:1;comment:用户ID" json:"user_id"`
	User        *User      `gorm:"foreignKey:UserId;references:Id" json:"user,omitempty"`
	SystemApiId uint       `gorm:"not null;uniqueIndex:uk_user_extra_api_user_api,priority:2;comment:接口ID" json:"system_api_id"`
	SystemApi   *SystemApi `gorm:"foreignKey:SystemApiId;references:Id" json:"system_api,omitempty"`
}

// 设置表名
func (UserExtraApi) TableName() string {
	return "user_extra_api"
}
