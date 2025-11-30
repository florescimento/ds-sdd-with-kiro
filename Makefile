.PHONY: help build test run docker-up docker-down clean lint fmt

# Default target
help:
	@echo "Available commands:"
	@echo "  make build        - Build all services"
	@echo "  make test         - Run tests"
	@echo "  make run          - Run a specific service (use SERVICE=name)"
	@echo "  make docker-up    - Start all Docker services"
	@echo "  make docker-down  - Stop all Docker services"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make lint         - Run linter"
	@echo "  make fmt          - Format code"

# Build all services
build:
	@echo "Building all services..."
	@go build -o bin/frontend ./cmd/frontend
	@go build -o bin/auth ./cmd/auth
	@go build -o bin/file-service ./cmd/file-service
	@go build -o bin/router-worker ./cmd/router-worker
	@go build -o bin/presence ./cmd/presence
	@echo "Build complete!"

# Build a specific service
build-%:
	@echo "Building $*..."
	@go build -o bin/$* ./cmd/$*

# Run tests
test:
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
	@echo "Tests complete!"

# Run tests with coverage report
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
	@go tool cover -html=coverage.txt -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run a specific service
run:
	@if [ -z "$(SERVICE)" ]; then \
		echo "Error: SERVICE not specified. Usage: make run SERVICE=frontend"; \
		exit 1; \
	fi
	@echo "Running $(SERVICE)..."
	@go run ./cmd/$(SERVICE)

# Start all Docker services
docker-up:
	@echo "Starting Docker services..."
	@docker-compose up -d
	@echo "Waiting for services to be healthy..."
	@sleep 10
	@docker-compose ps
	@echo "Docker services started!"

# Stop all Docker services
docker-down:
	@echo "Stopping Docker services..."
	@docker-compose down
	@echo "Docker services stopped!"

# Stop and remove all Docker services including volumes
docker-clean:
	@echo "Cleaning Docker services and volumes..."
	@docker-compose down -v
	@echo "Docker cleanup complete!"

# View Docker logs
docker-logs:
	@docker-compose logs -f

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -f coverage.txt coverage.html
	@echo "Clean complete!"

# Run linter
lint:
	@echo "Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Format complete!"

# Tidy dependencies
tidy:
	@echo "Tidying dependencies..."
	@go mod tidy
	@echo "Tidy complete!"

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@echo "Dependencies downloaded!"

# Initialize Kafka topics
kafka-init:
	@echo "Creating Kafka topics..."
	@docker exec -it kafka kafka-topics --create --topic message.events --bootstrap-server localhost:9092 --partitions 50 --replication-factor 1 --if-not-exists
	@docker exec -it kafka kafka-topics --create --topic message.status --bootstrap-server localhost:9092 --partitions 50 --replication-factor 1 --if-not-exists
	@docker exec -it kafka kafka-topics --create --topic file.events --bootstrap-server localhost:9092 --partitions 10 --replication-factor 1 --if-not-exists
	@docker exec -it kafka kafka-topics --create --topic presence.events --bootstrap-server localhost:9092 --partitions 10 --replication-factor 1 --if-not-exists
	@docker exec -it kafka kafka-topics --create --topic message.dlq --bootstrap-server localhost:9092 --partitions 10 --replication-factor 1 --if-not-exists
	@echo "Kafka topics created!"

# List Kafka topics
kafka-topics:
	@docker exec -it kafka kafka-topics --list --bootstrap-server localhost:9092

# Create MinIO bucket
minio-init:
	@echo "Creating MinIO bucket..."
	@docker exec -it minio mc alias set local http://localhost:9000 minioadmin minioadmin
	@docker exec -it minio mc mb local/chat-files --ignore-existing
	@echo "MinIO bucket created!"

# Full setup (Docker + Kafka + MinIO)
setup: docker-up
	@sleep 15
	@make kafka-init
	@make minio-init
	@echo "Setup complete! All services are ready."

# Development mode - start infrastructure and run a service
dev:
	@if [ -z "$(SERVICE)" ]; then \
		echo "Error: SERVICE not specified. Usage: make dev SERVICE=frontend"; \
		exit 1; \
	fi
	@make docker-up
	@sleep 10
	@make run SERVICE=$(SERVICE)
