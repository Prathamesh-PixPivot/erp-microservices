package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config struct holds all configuration values
type Config struct {
	Server struct {
		Port string
	}
}

// Cfg holds the loaded configuration
var Cfg Config

// LoadConfig reads the configuration from the config.yaml file
func LoadConfig() {
	viper.SetConfigFile("config/config.yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}
}
