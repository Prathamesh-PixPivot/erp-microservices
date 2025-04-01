package config

import (
	"log"

	"github.com/spf13/viper"
)

// DatabaseConfig holds database connection settings.
type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
	SSLMode  string
}

// Config holds configuration settings for the ITSM service.
type Config struct {
	// ITOM/General Settings
	ZabbixAPIURL string
	Port         string
	AuthToken    string
	UseTLS       bool
	TLSCertFile  string
	TLSKeyFile   string

	// Database Settings for ITSM
	Database DatabaseConfig

	// Logging Settings
	Logging struct {
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"`
		Output string `mapstructure:"output"`
	} `mapstructure:"logging"`
}

// LoadConfig loads configuration from a file (if available) and environment variables.
func LoadConfig() (*Config, error) {
	// Set configuration file details.
	viper.SetConfigName("config") // e.g., config.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // current directory; adjust if needed
	viper.AutomaticEnv()     // override with environment variables if present

	// Set default values for ITOM/General settings.
	viper.SetDefault("ZABBIX_API_URL", "http://172.27.5.246/api_jsonrpc.php")
	viper.SetDefault("PORT", "9090")
	viper.SetDefault("USE_TLS", false)
	viper.SetDefault("TLS_CERT_FILE", "")
	viper.SetDefault("TLS_KEY_FILE", "")

	// Set default values for the database configuration.
	viper.SetDefault("DATABASE.HOST", "localhost")
	viper.SetDefault("DATABASE.USER", "postgres")
	viper.SetDefault("DATABASE.PASSWORD", "password")
	viper.SetDefault("DATABASE.DBNAME", "itsm")
	viper.SetDefault("DATABASE.PORT", 5432)
	viper.SetDefault("DATABASE.SSLMODE", "disable")

	// Attempt to read the configuration file.
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: no config file found (%v), relying on environment variables", err)
	}

	// Construct the configuration object.
	cfg := &Config{
		ZabbixAPIURL: viper.GetString("ZABBIX_API_URL"),
		Port:         viper.GetString("PORT"),
		AuthToken:    viper.GetString("AUTH_TOKEN"),
		UseTLS:       viper.GetBool("USE_TLS"),
		TLSCertFile:  viper.GetString("TLS_CERT_FILE"),
		TLSKeyFile:   viper.GetString("TLS_KEY_FILE"),
		Database: DatabaseConfig{
			Host:     viper.GetString("DATABASE.HOST"),
			User:     viper.GetString("DATABASE.USER"),
			Password: viper.GetString("DATABASE.PASSWORD"),
			DBName:   viper.GetString("DATABASE.DBNAME"),
			Port:     viper.GetInt("DATABASE.PORT"),
			SSLMode:  viper.GetString("DATABASE.SSLMODE"),
		},
		Logging: struct {
			Level  string `mapstructure:"level"`
			Format string `mapstructure:"format"`
			Output string `mapstructure:"output"`
		}{
			Level:  viper.GetString("LOGGING.LEVEL"),
			Format: viper.GetString("LOGGING.FORMAT"),
			Output: viper.GetString("LOGGING.OUTPUT"),
		},
	}
	return cfg, nil
}
