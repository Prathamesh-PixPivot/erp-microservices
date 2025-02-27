package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config structure to hold all configuration values
type Config struct {
	Server      ServerConfig
	Database    DatabaseConfig
	Logging     LoggingConfig
	ExternalAPI ExternalAPIConfig
	Cache       CacheConfig
}

// ServerConfig holds gRPC server settings
type ServerConfig struct {
	GRPCPort int `mapstructure:"grpc_port"`
}

// DatabaseConfig holds database connection settings
type DatabaseConfig struct {
	Host               string `mapstructure:"host"`
	Port               int    `mapstructure:"port"`
	User               string `mapstructure:"user"`
	Password           string `mapstructure:"password"`
	DBName             string `mapstructure:"dbname"`
	SSLMode            string `mapstructure:"sslmode"`
	MaxConnections     int    `mapstructure:"max_connections"`
	MaxIdleConnections int    `mapstructure:"max_idle_connections"`
}

// LoggingConfig holds logging settings
type LoggingConfig struct {
	Level string `mapstructure:"level"`
	File  string `mapstructure:"file"`
}

// ExternalAPIConfig holds external API configurations (GST portal)
type ExternalAPIConfig struct {
	GSTPortalBaseURL string `mapstructure:"gst_portal_base_url"`
	APIKey           string `mapstructure:"api_key"`
}

// CacheConfig holds Redis cache settings
type CacheConfig struct {
	RedisHost     string `mapstructure:"redis_host"`
	RedisPort     int    `mapstructure:"redis_port"`
	RedisPassword string `mapstructure:"redis_password"`
	RedisDB       int    `mapstructure:"redis_db"`
}

// LoadConfig loads configuration from config.yaml using Viper
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("internal/infrastructure/config") // Path to look for config.yaml

	viper.AutomaticEnv() // Override with environment variables if available

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	return &config, nil
}
