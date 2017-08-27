package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "bawx",
	Short: "Client to access your digital box",
}

func init() {
	BoxNewCmd.Flags().StringVarP(&boxName, "name", "n", "", "Box name")
	BoxCmd.AddCommand(BoxNewCmd, BoxListCmd, BoxNewCmd)
	RootCmd.AddCommand(BoxCmd, LinkCmd)
}
