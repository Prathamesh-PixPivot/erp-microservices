package main

import (
	"log"
	"hrms/infrastructure/migrations"
)

func main() {
	// Initialize dependencies using Wire
	deps, err := InitializeDependencies()
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}

	// Migrations
	migrations.Migrate(deps.Database.DB, deps.Logger)

	// Start the gRPC server
	StartServer(deps)
}
