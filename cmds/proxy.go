package cmds

import "github.com/spf13/cobra"

// 配置任务服务的API接口
// 配置任务类型和回调地址

// ProxyCmd ...
var ProxyCmd = &cobra.Command{
	Use:     "proxy",
	Short:   "kake proxy",
	Run: func(cmd *cobra.Command, args []string) {
	},
}