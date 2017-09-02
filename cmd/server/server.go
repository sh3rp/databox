package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/sh3rp/databox/db"
	"github.com/sh3rp/databox/search"
	"github.com/sh3rp/databox/web"
)

var V_MAJOR = 0
var V_MINOR = 1
var V_PATCH = 0

func main() {
	log.Info().Msgf("Loading databox %s", getVersion())
	err := os.Mkdir("db", 0600)

	if err != nil {
		panic(err)
	}

	database := db.NewBoltDB("db/box.db")
	search := search.NewBoltSearchEngine("db/search.db")
	server := web.NewServer(8080, 5656, database, search)
	server.Start()
}

func getVersion() string {
	return fmt.Sprintf("v%d.%d.%d", V_MAJOR, V_MINOR, V_PATCH)
}
