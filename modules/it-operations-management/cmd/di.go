//go:build wireinject
// +build wireinject

package main

import (
	"time"

	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"

	"itom/config"
	configT "itom/internal"
	service "itom/internal"
	zabbix "itom/internal"
)

type Dependencies struct {
	Config       *configT.Config
	Logger       *zap.Logger
	ZabbixClient *zabbix.ZabbixClient
	ITOMService  *service.ITOMService
	Server       *Server
}

func NewZapLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}

func NewZabbixClient(cfg *configT.Config) *zabbix.ZabbixClient {
	return &zabbix.ZabbixClient{
		APIURL:    cfg.ZabbixAPIURL,
		AuthToken: cfg.AuthToken,
		Timeout:   10 * time.Second,
	}
}

func NewITOMService(logger *zap.Logger, client *zabbix.ZabbixClient) *service.ITOMService {
	tracer := otel.Tracer("itom-service")
	return &service.ITOMService{
		Logger:       logger,
		ZabbixClient: client,
		Tracer:       tracer,
	}
}

func NewServer(cfg *configT.Config, srv *service.ITOMService) *Server {
	return NewFiberServer(cfg, srv)
}

var wireSet = wire.NewSet(
	config.LoadConfig, // load config from file/env
	NewZapLogger,
	NewZabbixClient,
	NewITOMService,
	NewServer,
)

func InitializeDependencies() (*Dependencies, error) {
	wire.Build(
		wireSet,
		wire.Struct(new(Dependencies), "*"),
	)
	return &Dependencies{}, nil
}
