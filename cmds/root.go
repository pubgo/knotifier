package cmds

import (
	"github.com/pubgo/assert"
	"github.com/pubgo/knotifier/internal/config"
	"github.com/pubgo/knotifier/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:     "kake",
	Short:   "kake app",
	Version: version.Version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return assert.KTry(func() {
			cfg := config.DefaultConfig()
			assert.ErrWrap(viper.Unmarshal(cfg.Cfg), "config init error")
		})
	},
}
