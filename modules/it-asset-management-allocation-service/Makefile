wire:
	@echo "Generating Wire dependencies..."
	@wire ./cmd

build:
	go build -o asset-management-service ./cmd/main.go

run: build
	./asset-management-service
