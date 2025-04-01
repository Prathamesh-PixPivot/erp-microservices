// config.go
package config

import (
	"log"

	itom "itom/internal"

	"github.com/spf13/viper"
)

func LoadConfig() (*itom.Config, error) {
	// Set the file name of the configurations file (without extension)
	viper.SetConfigName("config")
	// Set the configuration file type (e.g., yaml, json, toml)
	viper.SetConfigType("yaml")
	// Set the path to look for the configurations file
	viper.AddConfigPath(".") // current directory, adjust if needed

	// Read in environment variables that match
	viper.AutomaticEnv()

	// Set default values (optional)
	viper.SetDefault("ZABBIX_API_URL", "http://172.27.5.246/api_jsonrpc.php")
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("USE_TLS", false)

	// Read configuration file (if available)
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: no config file found (%v), relying on environment variables", err)
	}

	config := &itom.Config{
		ZabbixAPIURL: viper.GetString("ZABBIX_API_URL"),
		Port:         viper.GetString("PORT"),
		AuthToken:    viper.GetString("AUTH_TOKEN"),
		UseTLS:       viper.GetBool("USE_TLS"),
		TLSCertFile:  viper.GetString("TLS_CERT_FILE"),
		TLSKeyFile:   viper.GetString("TLS_KEY_FILE"),
	}
	return config, nil
}
