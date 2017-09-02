package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "bawx",
	Short: "Client to access your digital box",
}

func init() {
	BoxNewCmd.Flags().StringVarP(&boxName, "name", "n", "", "Box name")
	BoxNewCmd.Flags().StringVarP(&boxDescription, "description", "d", "", "Box description")
	BoxNewCmd.Flags().BoolVarP(&setBoxEnv, "setEnv", "e", false, "Set newly created box to current working box")

	BoxSetCmd.Flags().StringVarP(&boxId, "id", "i", "", "Box id")

	BoxCmd.AddCommand(BoxNewCmd, BoxListCmd, BoxGetCmd, BoxSetCmd)

	LinkLoadCmd.Flags().StringVarP(&linkId, "id", "i", "", "Link id")
	LinkLoadCmd.Flags().StringVarP(&linkBoxId, "box", "b", "", "Box contain links to load")

	LinkAddCmd.Flags().StringVarP(&linkName, "name", "n", "", "Link name")
	LinkAddCmd.Flags().StringVarP(&linkUrl, "url", "u", "", "Link URL")
	LinkAddCmd.Flags().StringVarP(&linkBoxId, "box", "b", "", "Box to store this link in")
	LinkAddCmd.Flags().StringVarP(&linkTags, "tag", "t", "", "Tags to apply to link; comma delimited")

	LinkGetLinksCmd.Flags().StringVarP(&linkBoxId, "box", "b", "", "Box id of the box containing links")
	LinkGetLinksCmd.Flags().StringVarP(&linkTags, "tag", "t", "", "Tag to search for")

	LinkTagCmd.Flags().StringVarP(&linkId, "id", "i", "", "Link id")
	LinkTagCmd.Flags().StringVarP(&linkTags, "tag", "t", "", "Tags to apply to link; comma delimited")

	LinkSearchCmd.Flags().StringVarP(&searchTerm, "term", "t", "", "Tag term to search for")
	LinkSearchCmd.Flags().IntVarP(&searchCount, "count", "c", 10, "Number of results to return back")
	LinkSearchCmd.Flags().IntVarP(&searchPage, "page", "p", 0, "Result page to return")
	LinkSearchCmd.Flags().BoolVarP(&searchLoadLinks, "load", "l", false, "Load the links returned in a browser")

	LinkCmd.AddCommand(LinkAddCmd, LinkGetLinksCmd, LinkLoadCmd, LinkTagCmd, LinkSearchCmd)
	RootCmd.AddCommand(BoxCmd, LinkCmd)
}
