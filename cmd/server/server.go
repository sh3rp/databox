package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/sh3rp/databox/db"
	"github.com/sh3rp/databox/web"
)

var V_MAJOR = 0
var V_MINOR = 1
var V_PATCH = 0

func main() {
	log.Info().Msgf("Loading databox %s", getVersion())
	database := db.NewInMemoryDB()
	server := web.NewServer(8080, 5656, database)
	server.Start()
}

func getVersion() string {
	return fmt.Sprintf("v%d.%d.%d", V_MAJOR, V_MINOR, V_PATCH)
}
