//go:build wireinject
// +build wireinject

package main

import (
	"api-gateway/infrastructure"
	"api-gateway/internal/delivery/graphql"
	"api-gateway/internal/repository"
	"api-gateway/internal/usecase"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Dependencies struct groups all dependencies
type Dependencies struct {
	Echo     *echo.Echo
	Logger   *zap.Logger
	Resolver *graphql.Resolver
}

// ProvideEcho initializes Echo
func ProvideEcho() *echo.Echo {
	e := echo.New()
	return e
}

// ProvideLogger initializes Zap Logger
func ProvideLogger() *zap.Logger {
	infrastructure.InitLogger("DEV") // Change "DEV" to "PROD" for production
	return infrastructure.GetLogger()
}

// ProvideUserRepository initializes UserRepository
func ProvideUserRepository() repository.UserRepository {
	return repository.NewUserRepository()
}

// ProvideWorkflowRepository initializes WorkflowRepository
func ProvideWorkflowRepository() repository.WorkflowRepository {
	return repository.NewWorkflowRepository()
}

// ProvideUserUseCase initializes UserUseCase with UserRepository
func ProvideUserUseCase(repo repository.UserRepository) usecase.UserUseCase {
	return usecase.NewUserUseCase(repo)
}

// ProvideWorkflowUseCase initializes WorkflowUseCase with WorkflowRepository
func ProvideWorkflowUseCase(repo repository.WorkflowRepository) usecase.WorkflowUseCase {
	return usecase.NewWorkflowUseCase(repo)
}

// ProvideGraphQLResolver initializes the main Resolver
func ProvideGraphQLResolver(
	userUseCase usecase.UserUseCase,
	workflowUseCase usecase.WorkflowUseCase,
) *graphql.Resolver {
	return &graphql.Resolver{
		UserUseCase:     userUseCase,
		WorkflowUseCase: workflowUseCase,
	}
}

// InitializeDependencies sets up all dependencies using Google Wire
func InitializeDependencies() (*Dependencies, error) {
	wire.Build(
		ProvideEcho,
		ProvideLogger,
		ProvideUserRepository,
		ProvideWorkflowRepository,
		ProvideUserUseCase,
		ProvideWorkflowUseCase,
		ProvideGraphQLResolver,
		wire.Struct(new(Dependencies), "Echo", "Logger", "Resolver"),
	)
	return &Dependencies{}, nil
}
