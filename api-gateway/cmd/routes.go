package main

import (
	"api-gateway/config"
	"api-gateway/graph"
	"api-gateway/internal/delivery/graphql"
	v1 "api-gateway/internal/delivery/rest/v1"
	v2 "api-gateway/internal/delivery/rest/v2"
	"api-gateway/internal/middleware"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

// GraphQLHandler Echo-Compatible GraphQL Handler
func GraphQLHandler(gqlResolver *graphql.Resolver) echo.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: gqlResolver}))
	return echo.WrapHandler(h)
}

// RegisterRoutes sets up API and GraphQL routes dynamically
func RegisterRoutes(e *echo.Echo, gqlResolver *graphql.Resolver) {
	// Versioned REST API routes
	for _, version := range config.APIVersions {
		apiPrefix := fmt.Sprintf("/api/%s", version)
		vGroup := e.Group(apiPrefix, middleware.DeprecationMiddleware()) // ✅ Adds versioning

		if version == "v1" {
			vGroup.GET("/health", v1.HealthCheckHandler)
			vGroup.GET("/greet", v1.GreetHandler)
		}

		if version == "v2" {
			vGroup.GET("/health", v1.HealthCheckHandler) // Shared from v1
			vGroup.GET("/greet", v2.GreetHandler)        // v2-specific greeting
			vGroup.GET("/info", v2.InfoHandler)          // New v2-only feature
		}
	}

	// ✅ Register GraphQL Routes (Now Compatible with Echo)
	e.POST("/graphql", GraphQLHandler(gqlResolver))
	e.GET("/graphql/playground", echo.WrapHandler(playground.Handler("GraphQL Playground", "/graphql")))
}
