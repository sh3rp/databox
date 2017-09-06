package config

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"

	"github.com/sh3rp/databox/msg"
)

var CONFIG_DIR = ".bawx"
var CONFIG_FILE = "bawx.config"

type ClientConfig struct {
	DefaultBoxId string `json:"default_box"`
	Server       string `json:"server"`
	Port         int    `json:"port"`
}

func (cfg *ClientConfig) Read() *ClientConfig {
	usr, err := user.Current()

	if err != nil {
		panic(err)
	}

	_, err = os.Open(usr.HomeDir + "/" + CONFIG_DIR)

	if err != nil {
		fmt.Printf("Error opening config directory, attempting to create\n")
		err = os.Mkdir(usr.HomeDir+"/"+CONFIG_DIR, 0700)

		if err != nil {
			panic(err)
		}
	}

	filename := usr.HomeDir + "/" + CONFIG_DIR + "/" + CONFIG_FILE

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

	return cfg
}

func (cfg *ClientConfig) Write() {
	usr, err := user.Current()

	if err != nil {
		panic(err)
	}

	_, err = os.Open(usr.HomeDir + "/" + CONFIG_DIR)

	if err != nil {
		fmt.Printf("Error opening config directory, attempting to create\n")
		err = os.Mkdir(usr.HomeDir+"/"+CONFIG_DIR, 0700)

		if err != nil {
			panic(err)
		}
	}

	configFile, err := os.OpenFile(usr.HomeDir+"/"+CONFIG_DIR+"/"+CONFIG_FILE, os.O_RDWR|os.O_TRUNC, 0)

	err = json.NewEncoder(configFile).Encode(cfg)

	if err != nil {
		fmt.Printf("Config write error: %v\n", err)
	}

	configFile.Close()
}

func (cfg *ClientConfig) getConfigFile() (*os.File, error) {
	usr, err := user.Current()

	if err != nil {
		return nil, err
	}

	_, err = os.Open(usr.HomeDir + "/" + CONFIG_DIR)

	if err != nil {
		fmt.Printf("Error opening config directory, attempting to create\n")
		err = os.Mkdir(usr.HomeDir+"/"+CONFIG_DIR, 0700)

		if err != nil {
			return nil, err
		}
	}

	configFile, err := os.OpenFile(usr.HomeDir+"/"+CONFIG_DIR+"/"+CONFIG_FILE, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0)

	return configFile, nil
}

func WriteToken(token *msg.Token) error {
	usr, err := user.Current()

	if err != nil {
		return err
	}

	file, err := os.Open(usr.HomeDir + "/token.json")

	err = json.NewEncoder(file).Encode(token)

	return err
}

func ReadToken() (*msg.Token, error) {
	usr, err := user.Current()

	if err != nil {
		return nil, err
	}

	file, err := os.Open(usr.HomeDir + "/token.json")

	var token *msg.Token

	json.NewDecoder(file).Decode(token)

	return token, nil
}
