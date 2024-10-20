package config

import (
	"fmt"
	"log"
	"user-service/internal/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Declare the DB variable globally

// InitConfig initializes configuration using Viper
func InitConfig() {
	viper.SetConfigName("config") // Config file name (without extension)
	viper.SetConfigType("yaml")   // Config file type
	viper.AddConfigPath(".")      // Look for config in the current directory
	viper.AutomaticEnv()          // Override config with environment variables, if set

	// Read in the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}

// InitDB initializes the database connection using Viper for configuration
func InitDB() {
	var err error

	// Ensure all config values are loaded properly
	host := viper.GetString("DB_HOST")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	dbname := viper.GetString("DB_NAME")
	port := viper.GetString("DB_PORT")

	if host == "" || user == "" || password == "" || dbname == "" || port == "" {
		log.Fatalf("Database configuration values are missing")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbname, port,
	)

	// Open the connection using the DSN
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations for the User model
	DB.AutoMigrate(&models.User{})
	log.Println("Database connected and migrated successfully")
}
