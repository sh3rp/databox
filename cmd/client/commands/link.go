package commands

import (
	"context"

	"github.com/sh3rp/databox/msg"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
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
		conn, err := grpc.Dial("127.0.0.1:5656", grpc.WithInsecure())

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := msg.NewBoxServiceClient(conn)
		client.GetLinkById(context.Background())

	},
}
