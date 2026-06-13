package model

// 用户附加授权接口模型
type SystemUserExtraApi struct {
	BaseModel
	GrantType    uint        `gorm:"type:tinyint unsigned;not null;default:1;index:idx_system_user_extra_api_user_grant_api;comment:授权类型（1：额外授权，2：取消授权）" json:"grant_type"`
	SystemUserId uint        `gorm:"not null;uniqueIndex:uk_system_user_extra_api_user_api,priority:1;comment:用户ID" json:"system_user_id"`
	SystemUser   *SystemUser `gorm:"foreignKey:SystemUserId;references:Id" json:"system_user,omitempty"`
	SystemApiId  uint        `gorm:"not null;uniqueIndex:uk_system_user_extra_api_user_api,priority:2;comment:接口ID" json:"system_api_id"`
	SystemApi    *SystemApi  `gorm:"foreignKey:SystemApiId;references:Id" json:"system_api,omitempty"`
}

// 设置表名
func (SystemUserExtraApi) TableName() string {
	return "system_user_extra_api"
}
