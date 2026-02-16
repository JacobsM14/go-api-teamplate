.PHONY: run build test clean help

# Variables
BINARY_NAME=app
MAIN_PATH=main.go

# Default target
help:
	@echo "Available commands:"
	@echo "  make run     - Run the application"
	@echo "  make build   - Build the binary"
	@echo "  make test    - Run tests"
	@echo "  make clean   - Clean built files"

# Run the application
run:
	@go run $(MAIN_PATH)

# Build binary
build:
	@echo "Building..."
	@go build -o bin/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete: bin/$(BINARY_NAME)"

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean built files
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@go clean
	@echo "Clean complete"

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Run with hot reload (requires air)
dev:
	@air