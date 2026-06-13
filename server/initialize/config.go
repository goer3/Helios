package initialize

import (
	"Helios/common"
	"bytes"
	"log"
	"os"

	"github.com/spf13/viper"
)

// 配置初始化
func Config() {
	var bs []byte
	var err error
	var configType string = "yaml"

	// 初始化 viper，设置配置文件类型
	v := viper.New()
	v.SetConfigType(configType)

	// 读取配置文件内容
	if common.ParamSystemConfigFile != "" {
		// 如果启动参数有指定配置文件，则使用指定的外部配置文件进行初始化
		log.Printf("使用外部配置文件进行初始化: %s", common.ParamSystemConfigFile)
		bs, err = os.ReadFile(common.ParamSystemConfigFile)
	} else {
		// 否则使用默认配置文件，读取方式不一样
		log.Printf("使用默认配置文件进行初始化: config/default.yaml")
		bs, err = common.FS.ReadFile("config/default.yaml")
	}

	// 读取配置文件失败
	if err != nil {
		log.Fatalln("读取配置文件失败：" + err.Error())
	}

	// viper 解析配置文件
	if err = v.ReadConfig(bytes.NewReader(bs)); err != nil {
		log.Fatalln("解析配置文件失败：" + err.Error())
	}

	// 反序列化配置，使用 UnmarshalExact 进行严格类型检查
	if err = v.UnmarshalExact(&common.Config); err != nil {
		log.Fatalln("反序列化配置失败（类型不匹配）：" + err.Error())
	}
}
