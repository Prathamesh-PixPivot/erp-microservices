package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	GRPCPort string
	WSAddr   string
	DB       string
	Keycloak string
}

func Load() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	cfg := &Config{
		GRPCPort: viper.GetString("grpc.port"),
		WSAddr:   viper.GetString("websocket.addr"),
		DB:       viper.GetString("database.url"),
		Keycloak: viper.GetString("keycloak.url"),
	}

	return cfg
}
