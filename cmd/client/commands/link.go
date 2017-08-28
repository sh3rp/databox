package commands

import (
	"github.com/spf13/cobra"
)

var linkName string
var linkUrl string
var linkBoxId string
var linkTags string

var LinkCmd = &cobra.Command{
	Use:   "link",
	Short: "Manage your links",
}

var LinkAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a link",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var LinkGetLinksCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a link(s)",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var LinkLoadCmd = &cobra.Command{
	Use:   "get",
	Short: "Load a link(s)",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
