package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"

	"github.com/sh3rp/databox/auth"
	"github.com/sh3rp/databox/util"
	"golang.org/x/crypto/ssh/terminal"
)

var user string
var action string
var dbpath string

func main() {
	flag.StringVar(&user, "u", "", "Username to add/delete")
	flag.StringVar(&action, "a", "", "Action to perform [add|del]")
	flag.StringVar(&dbpath, "p", "/tmp/bawx/user.db", "Path to the user database")
	flag.Parse()

	if action == "" {
		fmt.Printf("You must specify an action.\n")
		os.Exit(1)
	}

	if user == "" {
		fmt.Printf("You must specify a user.\n")
		os.Exit(1)
	}

	switch action {
	case "add":
		addUser(user, dbpath)
	case "del":
		delUser(user, dbpath)
	}

}

func addUser(user, dbpath string) {
	auth := auth.NewBoltDBAuth(dbpath)

	fmt.Print("Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	err = auth.AddUser(user, util.GetPassHash(password))

	if err != nil {
		fmt.Printf("\nError adding user: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n")
}

func delUser(user, dbpath string) {
	auth := auth.NewBoltDBAuth(dbpath)

	err := auth.DeleteUser(user)

	if err != nil {
		fmt.Printf("\nError deleting user: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n")
}
