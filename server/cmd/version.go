package cmd

import (
	"Helios/common"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

// 版本命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示 Helios 服务版本信息",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Helios " + common.PROJECT_VERSION)
	},
}
