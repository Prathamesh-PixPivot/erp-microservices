package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Server   ServerConfig   `yaml:"server"`
}

type DatabaseConfig struct {
	DSN string `yaml:"dsn"`
}

type ServerConfig struct {
	GRPCPort string `yaml:"grpc_port"`
	HTTPPort string `yaml:"http_port"`
}

// LoadConfig loads the configuration from config.yaml
func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
