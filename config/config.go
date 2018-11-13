package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config implements configuration entity
type Config struct {
	Address       string
	Username      string
	Password      string
	Database      string
	ServerAddress string
	FirstStart    bool
}

// Load returns config entity loaded from file
func Load() Config {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	config := Config{}
	json.Unmarshal(data, &config)
	return config
}
