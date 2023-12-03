package config

import (
	"encoding/json"
	"os"
)

// configFilePath is path of config file
const configFilePath = "configs/config.json"

// Config is structure of a valid config for URL shortener service
type Config struct {
	ServerConfig struct {
		Hostname string `json:"hostname"`
		Port     int    `json:"port"`
	} `json:"server_config"`
	DatabaseConfig struct {
		Address  string `json:"address"`
		Password string `json:"password"`
	} `json:"database_config"`
}

var configContent Config

// GetConfig returns loaded config.
//
// Use 'LoadConfig' function if you never loaded config file.
func GetConfig() Config { return configContent }

// LoadConfig tries to load config from json config file.
//
// Use 'GetConfig' to read loaded config.
func LoadConfig() error {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &configContent)
}
