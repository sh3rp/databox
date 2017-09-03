package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ServerConfig struct {
	GrpcPort      int    `json:"grpc_port"`
	HttpPort      int    `json:"http_port"`
	DataDirectory string `json:"data_directory"`
}

func (cfg *ServerConfig) Read(filename string) {
	configFile, err := os.Open(filename)

	if err != nil {
		newFile, newFileErr := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

		if newFileErr != nil {
			panic(newFileErr)
		}

		json.NewEncoder(newFile).Encode(cfg)
		newFile.Close()

		configFile, err = os.Open(filename)
	}

	err = json.NewDecoder(configFile).Decode(cfg)

	if err != nil {
		fmt.Printf("Config read error: %v\n", err)
	}

	configFile.Close()
}
