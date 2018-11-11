package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config ...
type Config struct {
	Address       string
	Username      string
	Password      string
	Database      string
	ServerAddress string
}

// Load ...
func Load() Config {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	config := Config{}
	json.Unmarshal(data, &config)
	return config
}
