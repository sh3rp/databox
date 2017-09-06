package commands

import (
	"context"
	"fmt"
	"syscall"

	"github.com/sh3rp/databox/config"
	"github.com/sh3rp/databox/msg"
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
		if err == nil {
			fmt.Println("\nPassword typed: " + string(bytePassword))
		}
		password := string(bytePassword)

		conn, err := Dial()

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := msg.NewBoxServiceClient(conn)

		res, err := client.Authenticate(context.Background(), &msg.AuthRequest{
			Username: username,
			Password: password,
		})

		if res.Code == 0 {
			config.WriteToken(res.Token)
		} else {
			fmt.Printf("Error: %v\n", res.Message)
		}
	},
}
