package commands

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"syscall"

	"github.com/sh3rp/databox/config"
	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/util"
	"golang.org/x/crypto/ssh/terminal"

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

		token, err := config.ReadToken()

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Print("Box password: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		hasher := sha256.New()
		hasher.Write(bytePassword)
		password := string(hex.EncodeToString(hasher.Sum(nil)))

		req := &msg.UnlockRequest{
			Token:       token,
			Box:         &msg.Box{Name: boxName, Description: boxDescription},
			BoxPassword: password,
		}

		box, err := client.NewBox(context.Background(), req)

		if err != nil {
			fmt.Printf("Error creating new box: %v\n", err)
		} else {
			if setBoxEnv {
				cfg := &config.ClientConfig{}
				cfg.Read()
				cfg.DefaultBoxId = box.Id.Id
				cfg.Write()
			}
			util.PrettyPrint(box)
		}
	},
}

var BoxUnlockCmd = &cobra.Command{
	Use:   "unlock",
	Short: "Unlock the specified box for writing into",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := Dial()

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := msg.NewBoxServiceClient(conn)

		if boxId == "" {
			fmt.Println("You must specify a boxId when unlocking a box.")
			os.Exit(1)
		}

		token, err := config.ReadToken()

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Print("Box password: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		hasher := sha256.New()
		hasher.Write(bytePassword)
		password := string(hex.EncodeToString(hasher.Sum(nil)))

		req := &msg.UnlockRequest{
			Token:       token,
			Box:         &msg.Box{Id: &msg.Key{Id: boxId}},
			BoxPassword: password,
		}

		box, err := client.UnlockBox(context.Background(), req)

		if err != nil {
			fmt.Printf("Error creating new box: %v\n", err)
		} else {
			if setBoxEnv {
				cfg := &config.ClientConfig{}
				cfg.Read()
				cfg.DefaultBoxId = box.Id.Id
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

		token, err := config.ReadToken()

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		req := &msg.Request{
			Token: token,
			Objects: &msg.Request_Box{
				Box: &msg.Box{Name: boxName, Description: boxDescription},
			},
		}

		boxes, err := client.GetBoxes(context.Background(), req)

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

		token, err := config.ReadToken()

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		req := &msg.Request{
			Token: token,
			Objects: &msg.Request_Box{
				Box: &msg.Box{Id: &msg.Key{Id: boxId, Type: msg.Key_BOX}},
			},
		}

		box, err := client.GetBoxById(context.Background(), req)

		if err != nil {
			fmt.Printf("Error getting box: %v\n", err)
			os.Exit(1)
		}

		cfg := &config.ClientConfig{}
		cfg.Read()
		cfg.DefaultBoxId = box.Id.Id
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
