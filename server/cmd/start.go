package cmd

import (
	"Helios/common"
	"Helios/initialize"
	"Helios/task"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

// 初始化命令
func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVar(&common.ParamSystemConfigFile, "config", "", "指定服务启动配置文件")
	startCmd.Flags().StringVar(&common.ParamSystemListenHost, "host", common.ParamSystemListenHost, "指定服务启动监听地址")
	startCmd.Flags().IntVar(&common.ParamSystemListenPort, "port", common.ParamSystemListenPort, "指定服务启动监听端口")
	startCmd.Flags().IntVar(&common.ParamSystemRoleLeader, "role-leader", common.ParamSystemRoleLeader, "指定服务是否参与领导者选举, 0: 不参与, 1: 参与")
	startCmd.Flags().IntVar(&common.ParamSystemRoleWorker, "role-worker", common.ParamSystemRoleWorker, "指定服务是否开启工作节点角色, 0: 不开启, 1: 开启")
	startCmd.Flags().IntVar(&common.ParamSystemRoleWeb, "role-web", common.ParamSystemRoleWeb, "指定服务是否开启 Web 后端服务角色, 0: 不开启, 1: 开启")
}

// 启动命令
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 Helios 服务，更多参数请使用 --help 查看",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(common.LOGO)

		// 1. 初始化配置（从 YAML 读取）
		initialize.Config()

		// 参数校验
		{
			// 监听地址校验并覆盖
			if cmd.Flags().Changed("host") {
				ip := net.ParseIP(common.ParamSystemListenHost)
				if ip == nil || !ip.IsGlobalUnicast() {
					log.Fatalln("监听地址参数 --host 设置错误，仅支持 IPV4 地址")
				}
				common.Config.System.Listen.Host = common.ParamSystemListenHost
			}

			// 监听端口校验并覆盖
			if cmd.Flags().Changed("port") {
				if common.ParamSystemListenPort < 80 || common.ParamSystemListenPort > 65535 {
					log.Fatalln("监听端口参数 --port 设置错误，仅支持 80-65535 之间的整数")
				}
				common.Config.System.Listen.Port = common.ParamSystemListenPort
			}

			// 角色参数校验并覆盖
			if cmd.Flags().Changed("role-leader") {
				v := common.ParamSystemRoleLeader
				if v != 0 && v != 1 {
					log.Fatalln("角色参数 --role-leader 设置错误，仅支持 0（不参与） 或 1（参与）")
				}
				common.Config.System.Role.Leader = v == 1
			}
			if cmd.Flags().Changed("role-worker") {
				v := common.ParamSystemRoleWorker
				if v != 0 && v != 1 {
					log.Fatalln("角色参数 --role-worker 设置错误，仅支持 0（不开启） 或 1（开启）")
				}
				common.Config.System.Role.Worker = v == 1
			}
			if cmd.Flags().Changed("role-web") {
				v := common.ParamSystemRoleWeb
				if v != 0 && v != 1 {
					log.Fatalln("角色参数 --role-web 设置错误，仅支持 0（不开启） 或 1（开启）")
				}
				common.Config.System.Role.Web = v == 1
			}

			// 最终校验配置参数
			if !common.Config.System.Role.Leader && !common.Config.System.Role.Worker && !common.Config.System.Role.Web {
				log.Fatalln("所有角色都未启用，无法启动服务，服务即将退出...")
			}
		}

		// 2. 日志初始化
		initialize.SystemLogger()
		initialize.AccessLogger()

		// 3. 数据库初始化
		initialize.MySQL()

		// 4. Redis 初始化
		initialize.Redis()

		// 5. 验证器初始化
		initialize.Validator()

		// 判断是否开启 Web 服务
		if common.Config.System.Role.Web {
			task.StartWebServer()
		} else {
			// 保持进程运行
			select {}
		}
	},
}
