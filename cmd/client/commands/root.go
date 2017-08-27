package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "bawx",
	Short: "Client to access your digital box",
}

func init() {
	BoxNewCmd.Flags().StringVarP(&boxName, "name", "n", "", "Box name")
	BoxCmd.AddCommand(BoxNewCmd, BoxListCmd, BoxNewCmd)

	LinkCmd.Flags().StringVarP(&linkName, "name", "n", "", "Link name")
	LinkCmd.Flags().StringVarP(&linkUrl, "url", "u", "", "Link URL")
	LinkCmd.Flags().StringVarP(&linkBoxId, "box", "b", "", "Box to store this link in")
	LinkCmd.Flags().StringVarP(&linkTags, "tag", "t", "", "Tags to apply to link")

	LinkCmd.AddCommand(LinkAddCmd, LinkGetLinksCmd, LinkLoadCmd)
	RootCmd.AddCommand(BoxCmd, LinkCmd)
}
