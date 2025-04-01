package main

import (
	"api-gateway/config"
	"api-gateway/internal/delivery/graphql"
	custommiddleware "api-gateway/internal/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// StartServer initializes Echo and starts the API Gateway
func StartServer(e *echo.Echo, logger *zap.Logger, graphQLResolver *graphql.Resolver) {
	// Load Configuration
	config.LoadConfig()

	// Initialize OpenTelemetry Tracing
	_, err := custommiddleware.InitTracer(logger)
	if err != nil {
		logger.Fatal("Failed to initialize OpenTelemetry Tracing", zap.Error(err))
	}
	_, err = custommiddleware.InitMetrics(logger)
	if err != nil {
		logger.Fatal("Failed to initialize OpenTelemetry Metrics", zap.Error(err))
	}

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(custommiddleware.RequestIDMiddleware(logger))
	e.Use(custommiddleware.RequestLoggerMiddleware(logger))

	custommiddleware.TracingMiddleware(e) // ✅ Add OpenTelemetry Tracing Middleware

	// Register Routes
	RegisterRoutes(e, graphQLResolver)

	// Start Server
	logger.Info("Starting API Gateway...", zap.String("port", config.Cfg.Server.Port))
	e.Logger.Fatal(e.Start(":" + config.Cfg.Server.Port))
}
