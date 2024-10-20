package config

import (
	"log"

	"github.com/spf13/viper"
)

// InitConfig initializes the configuration using Viper
func InitConfig() {
	viper.SetConfigFile("./config.yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
}
