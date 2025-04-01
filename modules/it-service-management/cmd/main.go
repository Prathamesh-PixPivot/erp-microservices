package main

import (
	"log"
)

func main() {
	deps, err := InitializeDependencies()
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}

	// Start ITSM gRPC server on port 9090 (for example).
	go func() {
		log.Println("Starting ITSM gRPC server on :9090 ...")
		if err := StartITSMServer(":9090", deps.ITSMServer); err != nil {
			log.Fatalf("ITSM gRPC server error: %v", err)
		}
	}()

	// Block forever (or implement graceful shutdown)
	select {}
}
