// graphql-gateway/config/config.go

package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	GraphQLPort             int
	AuthServiceHost         string
	AuthServicePort         int
	UserServiceHost         string
	UserServicePort         int
	OrganizationServiceHost string
	OrganizationServicePort int
	LeadServiceHost         string
	LeadServicePort         int
	OpportunityServiceHost  string
	OpportunityServicePort  int
	ContactServiceHost      string
	ContactServicePort      int
	ActivityServiceHost     string
	ActivityServicePort     int
	VMSServiceHost                 string
	VMSServicePort                 int
}

func LoadConfig() (*Config, error) {
	var cfg Config
	var err error

	// GraphQL Server Configuration
	cfg.GraphQLPort, err = getEnvAsInt("GRAPHQL_PORT", 4000)
	if err != nil {
		return nil, fmt.Errorf("invalid GRAPHQL_PORT: %v", err)
	}

	// AuthService gRPC Configuration
	cfg.AuthServiceHost = getEnv("AUTH_SERVICE_HOST", "localhost")
	cfg.AuthServicePort, err = getEnvAsInt("AUTH_SERVICE_PORT", 50053)
	if err != nil {
		return nil, fmt.Errorf("invalid AUTH_SERVICE_PORT: %v", err)
	}

	// UserService gRPC Configuration
	cfg.UserServiceHost = getEnv("USER_SERVICE_HOST", "localhost")
	cfg.UserServicePort, err = getEnvAsInt("USER_SERVICE_PORT", 50051)
	if err != nil {
		return nil, fmt.Errorf("invalid USER_SERVICE_PORT: %v", err)
	}

	// OrganizationService gRPC Configuration
	cfg.OrganizationServiceHost = getEnv("ORGANIZATION_SERVICE_HOST", "localhost")
	cfg.OrganizationServicePort, err = getEnvAsInt("ORGANIZATION_SERVICE_PORT", 50052)
	if err != nil {
		return nil, fmt.Errorf("invalid ORGANIZATION_SERVICE_PORT: %v", err)
	}

	// LeadService gRPC Configuration
	cfg.LeadServiceHost = getEnv("LEAD_SERVICE_HOST", "localhost")
	cfg.LeadServicePort, err = getEnvAsInt("LEAD_SERVICE_PORT", 50054)
	if err != nil {
		return nil, fmt.Errorf("invalid LEAD_SERVICE_PORT: %v", err)
	}

	// OpportunityService gRPC Configuration
	cfg.OpportunityServiceHost = getEnv("OPPORTUNITY_SERVICE_HOST", "localhost")
	cfg.OpportunityServicePort, err = getEnvAsInt("OPPORTUNITY_SERVICE_PORT", 50055)
	if err != nil {
		return nil, fmt.Errorf("invalid OPPORTUNITY_SERVICE_PORT: %v", err)
	}

	// ContactService gRPC Configuration
	cfg.ContactServiceHost = getEnv("CONTACT_SERVICE_HOST", "localhost")
	cfg.ContactServicePort, err = getEnvAsInt("CONTACT_SERVICE_PORT", 50056)
	if err != nil {
		return nil, fmt.Errorf("invalid CONTACT_SERVICE_PORT: %v", err)
	}

	// ActivityService gRPC Configuration
	cfg.ActivityServiceHost = getEnv("ACTIVITY_SERVICE_HOST", "localhost")
	cfg.ActivityServicePort, err = getEnvAsInt("ACTIVITY_SERVICE_PORT", 50057)
	if err != nil {
		return nil, fmt.Errorf("invalid ACTIVITY_SERVICE_PORT: %v", err)
	}

	// VMS gRPC Configuration
	cfg.VMSServiceHost = getEnv("VMS_HOST", "localhost")
	cfg.VMSServicePort, err = getEnvAsInt("VMS_PORT", 50058)
	if err != nil {
		return nil, fmt.Errorf("invalid VMS_PORT: %v", err)
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
