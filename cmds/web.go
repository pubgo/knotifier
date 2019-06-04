package cmds

import "github.com/spf13/cobra"

// 初始化 redis
// 初始化 mysql
// 提供 Prometheus API

// WebCmd ...
var WebCmd = &cobra.Command{
	Use:     "web",
	Short:   "kake web",
	Run: func(cmd *cobra.Command, args []string) {
	},
}