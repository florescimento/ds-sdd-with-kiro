# Setup Guide

## Prerequisites

### Required Software

1. **Go 1.21+**
   ```bash
   # macOS
   brew install go
   
   # Verify installation
   go version
   ```

2. **Docker & Docker Compose**
   ```bash
   # macOS
   brew install docker docker-compose
   
   # Verify installation
   docker --version
   docker-compose --version
   ```

3. **Make** (usually pre-installed on macOS)
   ```bash
   make --version
   ```

### Optional Tools

1. **golangci-lint** (for code linting)
   ```bash
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   ```

## Initial Setup

### 1. Clone and Configure

```bash
# Clone the repository
git clone <repository-url>
cd ds-sdd-with-kiro

# Copy environment configuration
cp .env.example .env

# Edit .env if needed (defaults work for local development)
```

### 2. Install Go Dependencies

```bash
# Download all dependencies
make deps

# Tidy up go.mod and go.sum
make tidy
```

### 3. Start Infrastructure

```bash
# Start all infrastructure services and initialize them
make setup
```

This command will:
- Start Kafka, Zookeeper, MongoDB, Redis, MinIO, etcd
- Start Prometheus, Grafana, and Jaeger for observability
- Create necessary Kafka topics
- Create MinIO bucket for file storage

Wait for all services to be healthy (about 30 seconds).

### 4. Verify Infrastructure

Check that all services are running:

```bash
docker-compose ps
```

All services should show status "Up" or "Up (healthy)".

### 5. Build Services

```bash
# Build all services
make build
```

Binaries will be created in the `bin/` directory.

## Running Services

### Run Individual Service

```bash
# Run frontend service
make run SERVICE=frontend

# Run auth service
make run SERVICE=auth

# Run file service
make run SERVICE=file-service

# Run router worker
make run SERVICE=router-worker

# Run presence service
make run SERVICE=presence
```

### Development Mode

Start infrastructure and run a service in one command:

```bash
make dev SERVICE=frontend
```

## Accessing Services

### Application Services
- Frontend API: http://localhost:8080
- Auth Service: http://localhost:8081
- File Service: http://localhost:8082
- Router Worker: http://localhost:8083
- Presence Service: http://localhost:8084

### Infrastructure Services
- Kafka: localhost:9092
- MongoDB: localhost:27017
- Redis: localhost:6379
- MinIO API: http://localhost:9000
- MinIO Console: http://localhost:9001 (minioadmin/minioadmin)
- etcd: localhost:2379

### Observability
- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000 (admin/admin)
- Jaeger UI: http://localhost:16686

## Testing

### Run All Tests

```bash
make test
```

### Run Tests with Coverage

```bash
make test-coverage
```

This generates `coverage.html` that you can open in a browser.

### Run Specific Package Tests

```bash
go test -v ./internal/models/...
go test -v ./internal/shared/utils/...
```

## Troubleshooting

### Port Conflicts

If you get port binding errors, check if ports are already in use:

```bash
# Check specific port
lsof -i :9092  # Kafka
lsof -i :27017 # MongoDB
lsof -i :6379  # Redis
```

Kill the process or change the port in `docker-compose.yml`.

### Docker Issues

```bash
# Clean up everything and start fresh
make docker-clean
make setup
```

### Kafka Topics Not Created

```bash
# Manually create topics
make kafka-init

# List topics to verify
make kafka-topics
```

### MinIO Bucket Not Created

```bash
# Manually create bucket
make minio-init
```

### Go Module Issues

```bash
# Clean and re-download dependencies
go clean -modcache
make deps
```

## Next Steps

1. Review the API documentation in `docs/API.md`
2. Check the architecture documentation in `.kiro/specs/distributed-chat-api/design.md`
3. Start implementing tasks from `.kiro/specs/distributed-chat-api/tasks.md`

## Stopping Services

### Stop Infrastructure Only

```bash
make docker-down
```

### Stop and Clean Everything

```bash
make docker-clean
```

This removes all containers, networks, and volumes.
