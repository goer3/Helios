package initialize

import (
	"Helios/common"
	"Helios/model"
	"Helios/pkg/utils"
)

// 初始化分页配置
func SystemPageSetting() {
	common.SystemLog.Info("开始初始化分页配置")
	var m model.SystemSetting
	if err := common.DB.Where("name = ?", common.SYSTEM_PAGE_SETTING_KEY).First(&m).Error; err != nil {
		common.SystemLog.Fatal("获取分页配置失败：", err)
	}
	// 将 value 转换成结构体
	pageSetting := &common.SystemPageSetting{}
	if err := utils.JsonStringToStruct(m.Value, pageSetting); err != nil {
		common.SystemLog.Fatal("解析分页配置失败：", err)
	}
	// 重新赋值
	common.SystemPageSettingConfig.PageSize = pageSetting.PageSize
	common.SystemPageSettingConfig.MaxPageSize = pageSetting.MaxPageSize
}
