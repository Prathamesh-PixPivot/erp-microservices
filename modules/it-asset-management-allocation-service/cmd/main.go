package main

import (
	"amaa/infrastructure/migrations"
	"log"
)

func main() {
	// Initialize dependencies using Wire
	deps, err := InitializeDI()
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}

	log.Println("Dependencies initialized successfully", deps)
	// Migrations
	migrations.Migrate(deps.Database.DB, deps.Logger)

	// Start the gRPC server
	StartServer(deps)
}
