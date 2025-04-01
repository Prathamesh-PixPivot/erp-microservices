package main

import (
	"log"
)

func main() {
	// Initialize dependencies using Wire
	deps, err := InitializeDependencies()
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}

	// Start API Gateway Server
	StartServer(deps.Echo, deps.Logger, deps.Resolver)
}
