// activity-task-service/config/config.go

package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	GRPCPort   int
}

func LoadConfig() (*Config, error) {
	var cfg Config
	var err error

	// Database Configuration
	cfg.DBHost = getEnv("DB_HOST", "localhost")
	cfg.DBPort, err = getEnvAsInt("DB_PORT", 5432)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %v", err)
	}
	cfg.DBUser = getEnv("DB_USER", "postgres")
	cfg.DBPassword = getEnv("DB_PASSWORD", "Ehsaas@2718")
	cfg.DBName = getEnv("DB_NAME", "pixerp")
	cfg.DBSSLMode = getEnv("DB_SSLMODE", "disable")

	// gRPC Server Configuration
	cfg.GRPCPort, err = getEnvAsInt("GRPC_PORT", 50057) // Default port
	if err != nil {
		return nil, fmt.Errorf("invalid GRPC_PORT: %v", err)
	}

	return &cfg, nil
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) (int, error) {
	valueStr := getEnv(name, "")
	if valueStr == "" {
		return defaultVal, nil
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, err
	}
	return value, nil
}
