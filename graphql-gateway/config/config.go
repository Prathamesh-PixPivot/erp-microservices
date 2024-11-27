// graphql-gateway/config/config.go

package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	GraphQLPort                int
	AuthServiceHost            string
	AuthServicePort            int
	UserServiceHost            string
	UserServicePort            int
	OrganizationServiceHost    string
	OrganizationServicePort    int
	LeadServiceHost            string
	LeadServicePort            int
	OpportunityServiceHost     string
	OpportunityServicePort     int
	ContactServiceHost         string
	ContactServicePort         int
	ActivityServiceHost        string
	ActivityServicePort        int
	VendorServiceHost          string
	VendorServicePort          int
	PaymentServiceHost         string
	PaymentServicePort         int
	PerformanceServiceHost     string
	PerformanceServicePort     int
	PurchaseOrderServiceHost   string
	PurchaseOrderServicePort   int
	InvoiceServiceHost         string
	InvoiceServicePort         int
	CreditDebitNoteServiceHost string
	CreditDebitNoteServicePort int
	PaymentDueServiceHost      string
	PaymentDueServicePort      int
	LedgerServiceHost          string
	LedgerServicePort          int
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

	// Vendor gRPC Configuration
	cfg.VendorServiceHost = getEnv("Vendor_HOST", "localhost")
	cfg.VendorServicePort, err = getEnvAsInt("Vendor_PORT", 50058)
	if err != nil {
		return nil, fmt.Errorf("invalid Vendor_PORT: %v", err)
	}

	// Invoice gRPC Configuration
	cfg.InvoiceServiceHost = getEnv("Invoice_HOST", "localhost")
	cfg.InvoiceServicePort, err = getEnvAsInt("Invoice_PORT", 50002)
	if err != nil {
		return nil, fmt.Errorf("invalid Invoice_PORT: %v", err)
	}

	// Payment gRPC Configuration
	cfg.PaymentServiceHost = getEnv("Payment_HOST", "localhost")
	cfg.PaymentServicePort, err = getEnvAsInt("Payment_PORT", 50058)
	if err != nil {
		return nil, fmt.Errorf("invalid Payment_PORT: %v", err)
	}

	// Performance gRPC Configuration
	cfg.PerformanceServiceHost = getEnv("Performance_HOST", "localhost")
	cfg.PerformanceServicePort, err = getEnvAsInt("Performance_PORT", 50058)
	if err != nil {
		return nil, fmt.Errorf("invalid Performance_PORT: %v", err)
	}

	// PurchaseOrder gRPC Configuration
	cfg.PurchaseOrderServiceHost = getEnv("PurchaseOrder_HOST", "localhost")
	cfg.PurchaseOrderServicePort, err = getEnvAsInt("PurchaseOrder_PORT", 50058)
	if err != nil {
		return nil, fmt.Errorf("invalid PurchaseOrder_PORT: %v", err)
	}

	// CreditNote gRPC Configuration
	cfg.CreditDebitNoteServiceHost = getEnv("CreditNote_HOST", "localhost")
	cfg.CreditDebitNoteServicePort, err = getEnvAsInt("CreditNote_PORT", 50002)
	if err != nil {
		return nil, fmt.Errorf("invalid CreditNote_PORT: %v", err)
	}

	// PaymentDue gRPC Configuration
	cfg.PaymentDueServiceHost = getEnv("PaymentDue_HOST", "localhost")
	cfg.PaymentDueServicePort, err = getEnvAsInt("PaymentDue_PORT", 50002)
	if err != nil {
		return nil, fmt.Errorf("invalid PaymentDue_PORT: %v", err)
	}

	// Ledger gRPC Configuration
	cfg.LedgerServiceHost = getEnv("Ledger_HOST", "localhost")
	cfg.LedgerServicePort, err = getEnvAsInt("Ledger_PORT", 50002)
	if err != nil {
		return nil, fmt.Errorf("invalid Ledger_PORT: %v", err)
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
