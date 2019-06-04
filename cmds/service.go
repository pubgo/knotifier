package cmds

import (
	"github.com/spf13/cobra"
)

// 初始化 redis
// 任务相关的API

// ServiceCmd ...
var ServiceCmd = &cobra.Command{
	Use:     "s",
	Aliases: []string{"s"},
	Short:   "kake service",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
