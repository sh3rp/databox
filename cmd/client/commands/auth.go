package commands

import (
	"context"
	"fmt"
	"os"
	"syscall"

	"github.com/sh3rp/databox/config"
	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/util"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var username string

var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate a user",
	Run: func(cmd *cobra.Command, args []string) {
		//reader := bufio.NewReader(os.Stdin)

		fmt.Print("Password: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		password := string(bytePassword)

		conn, err := Dial()

		if err != nil {
			fmt.Printf("Error connecting to server: %v\n", err)
			os.Exit(1)
		}

		defer conn.Close()

		client := msg.NewBoxServiceClient(conn)

		res, err := client.Authenticate(context.Background(), &msg.AuthRequest{
			Username: username,
			Password: util.GetPassHash(password),
		})

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		if res.Code == 0 {
			err = config.WriteToken(res.Token)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Printf("Wrote token\n")
			}
		} else {
			fmt.Printf("Error: %v\n", res.Message)
		}
	},
}
