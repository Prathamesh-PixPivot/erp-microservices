package internal

import (
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// Server represents the ITOM HTTP server.
type Server struct {
	Addr        string
	ITOMService *ITOMService
	Config      *Config
}

type Config struct {
	ZabbixAPIURL string
	Port         string
	AuthToken    string
	UseTLS       bool
	TLSCertFile  string
	TLSKeyFile   string
}

// ITOMService contains core business logic for processing alerts.
type ITOMService struct {
	Logger       *zap.Logger   // Centralized logger (Zap)
	ZabbixClient *ZabbixClient // Zabbix API client
	// Additional dependencies (e.g., ITSM client) can be added here.
	Tracer trace.Tracer // OpenTelemetry tracer for distributed tracing
}
