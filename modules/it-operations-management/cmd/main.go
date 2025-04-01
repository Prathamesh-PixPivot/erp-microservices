package main

import (
	"log"
)

func main() {
	deps, err := InitializeDependencies()
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}
	// You can now use deps.Server to start your ITOM service.
	log.Println("Starting ITOM server...")
	if err := deps.Server.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
