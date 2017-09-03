package commands

import (
	"github.com/sh3rp/databox/config"
	"github.com/sh3rp/databox/util"
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "conf",
	Short: "Shows the current client configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cc := &config.ClientConfig{}
		cc.Read()
		util.PrettyPrint(cc)
	},
}
