package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/sh3rp/databox/config"
	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/util"

	"github.com/spf13/cobra"
)

var boxId string
var boxName string
var boxDescription string

var setBoxEnv bool

var BoxCmd = &cobra.Command{
	Use:   "box",
	Short: "Manage your boxes",
}

var BoxNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new box",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := Dial()

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := msg.NewBoxServiceClient(conn)

		if boxName == "" {
			fmt.Println("You must specify a name when creating a box.")
			os.Exit(1)
		}

		box, err := client.NewBox(context.Background(), &msg.Box{Name: boxName, Description: boxDescription})

		if err != nil {
			fmt.Printf("Error creating new box: %v\n", err)
		} else {
			if setBoxEnv {
				cfg := &config.ClientConfig{}
				cfg.Read()
				cfg.DefaultBoxId = box.Id
				cfg.Write()
			}
			util.PrettyPrint(box)
		}
	},
}

var BoxListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all boxes",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := Dial()

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := msg.NewBoxServiceClient(conn)

		boxes, err := client.GetBoxes(context.Background(), &msg.None{})

		if err != nil {
			fmt.Printf("Error getting boxes: %v\n", err)
		} else {
			util.PrettyPrint(boxes)
		}
	},
}

var BoxSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the current box",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := Dial()

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := msg.NewBoxServiceClient(conn)

		box, err := client.GetBoxById(context.Background(), &msg.Box{Id: boxId})

		if err != nil {
			fmt.Printf("Error getting box: %v\n", err)
			os.Exit(1)
		}

		cfg := &config.ClientConfig{}
		cfg.Read()
		cfg.DefaultBoxId = box.Id
		cfg.Write()

		fmt.Printf("Set default box to %s\n", box.Name)
	},
}

var BoxGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a box config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Get BOX")
	},
}
