package cmd

import (
	"Helios/common"

	"github.com/spf13/cobra"
)

// 命令入口
var rootCmd = &cobra.Command{
	Use:   "helios",
	Short: common.PROJECT_DESCRIPTION,
}

func Execute() {
	rootCmd.Execute()
}
