// config/config.go

package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all the configuration for the Contact Management service.
type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	GRPCPort int

	// Add more configuration fields as needed
}

// LoadConfig reads configuration from environment variables.
// It returns a Config struct populated with the values.
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
	cfg.DBPassword = getEnv("DB_PASSWORD", "root")
	cfg.DBName = getEnv("DB_NAME", "contacts-service")
	cfg.DBSSLMode = getEnv("DB_SSLMODE", "disable")

	// gRPC Server Configuration
	cfg.GRPCPort, err = getEnvAsInt("GRPC_PORT", 50056) // Default to 50052 for contact service
	if err != nil {
		return nil, fmt.Errorf("invalid GRPC_PORT: %v", err)
	}

	return &cfg, nil
}

// getEnv reads an environment variable and returns its value or a default value if not set.
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// getEnvAsInt reads an environment variable into integer or returns a default value.
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
