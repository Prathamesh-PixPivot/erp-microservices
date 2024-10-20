package config

import (
	"os"
)

type Config struct {
	DB       string
	GRPCPort string
}

func Load() *Config {
	return &Config{
		DB:       getEnv("DB_CONNECTION_STRING", "host=localhost user=postgres password=root dbname=opportunity-service port=5432 sslmode=disable"),
		GRPCPort: getEnv("GRPC_PORT", ":50055"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
