package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/util"

	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

var boxName string

var BoxCmd = &cobra.Command{
	Use:   "box",
	Short: "Manage your boxes",
}

var BoxNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new box",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("127.0.0.1:5656", grpc.WithInsecure())

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := msg.NewBoxServiceClient(conn)

		if boxName == "" {
			fmt.Println("You must specify a name when creating a box.")
			os.Exit(1)
		}

		box, err := client.NewBox(context.Background(), &msg.Box{Name: boxName})

		util.PrettyPrint(box)
	},
}

var BoxListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all boxes",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("127.0.0.1:5656", grpc.WithInsecure())

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := msg.NewBoxServiceClient(conn)

		boxes, err := client.GetBoxes(context.Background(), &msg.None{})

		util.PrettyPrint(boxes)
	},
}

var BoxGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a box config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Get BOX")
	},
}
