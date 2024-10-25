package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	GRPCPort string `yaml:"grpc_port"`
}

type DatabaseConfig struct {
	DSN string `yaml:"dsn"`
}

// LoadConfig loads configuration from config.yaml
func LoadConfig() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
