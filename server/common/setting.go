package common

// 需要从数据库初始化的配置项
// 分页 Key
const SYSTEM_PAGE_SETTING_KEY = "page_setting"

// 分页配置
type SystemPageSetting struct {
	PageSize    uint `json:"page_size"`     // 默认每页数据量
	MaxPageSize uint `json:"max_page_size"` // 最大每页数据量
}

var SystemPageSettingConfig = SystemPageSetting{
	PageSize:    10,  // 默认每页数据量
	MaxPageSize: 100, // 默认最大每页数据量
}
