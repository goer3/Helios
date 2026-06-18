package model

// 系统设置
type SystemSetting struct {
	Name        string `gorm:"uniqueIndex:uk_system_setting_name;comment:设置名称" json:"name"`
	Description string `gorm:"comment:设置描述" json:"description"`
	Value       string `gorm:"comment:设置值" json:"value"`
}

func (s *SystemSetting) TableName() string {
	return "system_setting"
}
