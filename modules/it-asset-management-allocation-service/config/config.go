package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// Config structure to map YAML and ENV values
type Config struct {
	Server struct {
		Mode         string `mapstructure:"mode"`
		GRPCPort     int    `mapstructure:"grpc_port"`
		ReadTimeout  string `mapstructure:"read_timeout"`
		WriteTimeout string `mapstructure:"write_timeout"`
	} `mapstructure:"server"`

	Database struct {
		Driver   string `mapstructure:"driver"`
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
	} `mapstructure:"database"`

	// Redis struct {
	// 	Host     string
	// 	Port     int
	// 	Password string
	// } `mapstructure:"redis"`

	// APIGateway struct {
	// 	URL string
	// } `mapstructure:"api_gateway"`

	Logging struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
		Output string `mapstructure:"output"`
	} `mapstructure:"logging"`

	Tracing struct {
		Enabled      bool   `mapstructure:"enabled"`
		Exporter     string `mapstructure:"exporter"`
		OTLPEndpoint string
	} `mapstructure:"tracing"`

	Kafka struct {
		Enabled bool `mapstructure:"enabled"`
		Brokers []string
	} `mapstructure:"kafka"`
}

// LoadConfig reads configuration from file and environment variables
func LoadConfig() (*Config, error) {
	// ✅ Load .env file first
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found. Using system environment variables.")
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	// Replace `.` in env variables with `_`
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read config.yaml
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config.yaml file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}

	// ✅ Load required credentials from .env
	config.Database.Host = mustGetEnv("DB_HOST")
	config.Database.User = mustGetEnv("DB_USER")
	config.Database.Password = mustGetEnv("DB_PASSWORD")
	config.Database.DBName = mustGetEnv("DB_NAME")
	config.Database.SSLMode = getEnv("DB_SSL_MODE", "disable") // Default to "disable" for local

	// config.AuthService.URL = mustGetEnv("AUTH_SERVICE_URL")
	// config.APIGateway.URL = getEnv("API_GATEWAY_URL", "https://api.edgeflowtech.com")
	config.Tracing.OTLPEndpoint = getEnv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://otel-collector:4317")

	// ✅ Load Redis credentials
	// config.Redis.Host = getEnv("REDIS_HOST", "127.0.0.1")
	// config.Redis.Port = getIntFromEnv("REDIS_PORT", 6379)
	// config.Redis.Password = getEnv("REDIS_PASSWORD", "")

	// ✅ Convert integer environment variables
	config.Database.Port = getIntFromEnv("DB_PORT", 5432)
	config.Server.GRPCPort = getIntFromEnv("GRPC_PORT", 50051)

	// ✅ Convert comma-separated Kafka brokers to an array
	kafkaBrokers := getEnv("KAFKA_BROKER", "kafka1:9092,kafka2:9092,kafka3:9092")
	config.Kafka.Brokers = strings.Split(kafkaBrokers, ",")

	// ✅ Load Logging Configurations (from config.yaml)
	config.Logging.Level = viper.GetString("logging.level")
	config.Logging.Format = viper.GetString("logging.format")
	config.Logging.Output = viper.GetString("logging.output")

	fmt.Println("✅ Config successfully loaded")
	return &config, nil
}

// ✅ Helper function to enforce required variables
func mustGetEnv(envVar string) string {
	value, exists := os.LookupEnv(envVar)
	if !exists || value == "" {
		log.Fatalf("🚨 Missing required environment variable: %s", envVar)
	}
	return value
}

// ✅ getIntFromEnv safely converts an environment variable to an integer
func getIntFromEnv(envVar string, defaultValue int) int {
	if value, exists := os.LookupEnv(envVar); exists {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("🚨 Invalid value for %s: %v", envVar, err)
		}
		return intValue
	}
	return defaultValue
}

// ✅ getEnv safely retrieves environment variables with a default fallback
func getEnv(envVar string, defaultValue string) string {
	if value, exists := os.LookupEnv(envVar); exists {
		return value
	}
	return defaultValue
}
