package cmds

import (
	"github.com/pubgo/knotifier/version"
	"github.com/spf13/cobra"
	"log"
)

// VersionCmd ...
var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Show version info",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("version", version.Version)
		log.Println("commit version", version.CommitV)
		log.Println("build version", version.BuildV)
	},
}
