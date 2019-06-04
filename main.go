package main

import (
	"github.com/pubgo/assert"
	"github.com/pubgo/knotifier/cmds"
	"os"
)

func main() {
	rootCmd := cmds.RootCmd
	rootCmd.AddCommand(
		cmds.VersionCmd,
		cmds.ProxyCmd,
		cmds.ServiceCmd,
		cmds.WebCmd,
	)

	if err := assert.KTry(func() {
		assert.Throw(cmds.PrepareBaseCmd(rootCmd, "KN", os.ExpandEnv("$PWD/kata")).Execute())
	}); err != nil {
		assert.P(err)
		os.Exit(-1)
	}
}
