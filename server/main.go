package main

import (
	"Helios/cmd"
	"Helios/common"
	"embed"

	"github.com/dromara/carbon/v2"
)

//go:embed config/*
var fs embed.FS // Go 1.16 版本之后提供的将静态资源打包的方法，写法固定，可以将目录也打包

func main() {
	// 设置 carbon 默认配置
	carbon.SetDefault(carbon.Default{
		Layout:       carbon.DateTimeLayout,
		Timezone:     carbon.Local,
		Locale:       "en",
		WeekStartsAt: carbon.Monday,
		WeekendDays:  []carbon.Weekday{carbon.Saturday, carbon.Sunday},
	})
	carbon.SetTimezone(carbon.Local)
	common.FS = fs
	cmd.Execute()
}
