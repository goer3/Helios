package model

import "github.com/dromara/carbon/v2"

// 用户模型
// OnDelete 的选项：
// CASCADE：删角色时，关联用户也被删（风险高）
// SET NULL：删角色时，用户 system_role_id 置空（需允许 NULL）
// RESTRICT：有引用就不让删（最安全，常用于权限/组织等核心主数据）
type SystemUser struct {
	BaseModel
	Nickname     string         `gorm:"type:varchar(64);not null;index:idx_system_user_nickname;default:'';comment:昵称" json:"nickname"`
	Username     string         `gorm:"type:varchar(64);not null;uniqueIndex:uk_system_user_username;comment:用户名" json:"username"`
	Mobile       string         `gorm:"type:varchar(20);not null;uniqueIndex:uk_system_user_mobile;comment:手机号" json:"mobile"`
	HideMobile   uint           `gorm:"type:tinyint unsigned;not null;default:0;comment:是否隐藏手机号（0：否，1：是）" json:"hide_mobile"`
	Email        string         `gorm:"type:varchar(100);not null;uniqueIndex:uk_system_user_email;comment:邮箱" json:"email"`
	Password     string         `gorm:"type:varchar(255);not null;default:'';comment:密码" json:"-" copier:"-"` // JSON 返回和转换的时候忽略该字段
	Gender       uint           `gorm:"type:tinyint unsigned;not null;default:0;comment:性别（0：未知，1：男，2：女）" json:"gender"`
	AvatarUrl    string         `gorm:"type:varchar(255);not null;default:'';comment:头像" json:"avatar_url"`
	LastLoginIP  string         `gorm:"type:varchar(45);not null;default:'';comment:最后登录IP" json:"last_login_ip"`
	LastLoginAt  *carbon.Carbon `gorm:"type:datetime;index:idx_system_user_last_login_at;comment:最后登录时间" json:"last_login_at"`
	Status       uint           `gorm:"type:tinyint unsigned;not null;default:0;comment:用户状态（0：禁用，1：启用）" json:"status"`
	ExpireAt     *carbon.Carbon `gorm:"type:datetime;index:idx_system_user_expire_at;comment:过期时间" json:"expire_at" copier:"-"`
	SystemRoleId uint           `gorm:"not null;index:idx_system_user_role_id;comment:角色ID" json:"system_role_id"`
	SystemRole   *SystemRole    `gorm:"foreignKey:SystemRoleId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"system_role,omitempty"` // 指针类型解决 omitempty 为空问题
}

// 设置表名
func (SystemUser) TableName() string {
	return "system_user"
}
