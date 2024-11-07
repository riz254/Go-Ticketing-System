# Default target
start: build
	@echo "Starting the Go application..."
	@./bin/main || echo "Failed to start the application"

# Build target
build:
	@echo "Building the Go application..."
	@go build -o ./bin/main ./cmd/api/main.go || echo "Go build failed"
