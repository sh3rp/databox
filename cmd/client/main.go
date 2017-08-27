package main

import (
	"fmt"
	"os"

	"github.com/sh3rp/databox/cmd/client/commands"
)

func main() {
	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
