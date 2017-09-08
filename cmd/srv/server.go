package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sh3rp/databox/auth"
	"github.com/sh3rp/databox/config"
	"github.com/sh3rp/databox/db"
	"github.com/sh3rp/databox/search"
	"github.com/sh3rp/databox/server"
	"github.com/sh3rp/databox/util"
)

var configFile string

func main() {
	log.Info().Msgf("Loading databox %s", getVersion())

	flag.StringVar(&configFile, "c", "server.json", "Config file")
	flag.Parse()

	serverConfig := &config.ServerConfig{}
	serverConfig.Read(configFile)

	if _, err := os.Open(serverConfig.DataDirectory); err != nil {
		newMkdirErr := os.Mkdir(serverConfig.DataDirectory, 0755)

		if newMkdirErr != nil {
			panic(newMkdirErr)
		}
	}

	database := db.NewBoltDB(serverConfig.DataDirectory + "/box.db")
	search := search.NewBoltSearchEngine(serverConfig.DataDirectory + "/search.db")
	authenticator := auth.NewBoltDBAuth(serverConfig.DataDirectory + "/user.db")
	s := server.NewServer(serverConfig.HttpPort, serverConfig.GrpcPort, database, search, authenticator)
	s.Start()

	for {
		time.Sleep(1 * time.Second)
	}
}

func getVersion() string {
	return fmt.Sprintf("v%d.%d.%d", util.V_MAJOR, util.V_MINOR, util.V_PATCH)
}
