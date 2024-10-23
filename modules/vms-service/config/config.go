package config

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DBUrl      string `mapstructure:"DB_URL"`
	GRPCPort   string `mapstructure:"GRPC_PORT"`
	ServerPort string `mapstructure:"SERVER_PORT"`
}

var AppConfig *Config

// InitConfig loads configuration from config.yaml
func InitConfig() {
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	AppConfig = &Config{}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into config struct: %v", err)
	}

	log.Println("Configuration loaded successfully")
}

// InitDB initializes the database connection
func InitDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(AppConfig.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully")
	return db
}
