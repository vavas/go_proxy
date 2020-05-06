// Package config is gateway server configuration.
package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Server structure
type Server struct {
	Listen   string
}

// Config structure
type Config struct {
	Server   Server
}

// ReadConfig from a config file
func ReadConfig(configFile string) (config *Config) {

	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal("Unable to parse config file")
	}

	if len(config.Server.Listen) == 0 {
		config.Server.Listen = ":443"
	}

	return config
}
