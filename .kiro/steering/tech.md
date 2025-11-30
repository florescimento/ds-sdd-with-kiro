## Tech Stack

### Backend
- **Language**: Go 1.21+
- **HTTP Router**: gorilla/mux
- **WebSocket**: gorilla/websocket
- **JWT**: golang-jwt/jwt

### Infrastructure
- **Message Broker**: Apache Kafka (Confluent)
- **Database**: MongoDB 7.0
- **Cache**: Redis 7.2
- **Object Storage**: MinIO (S3-compatible)
- **Metadata Store**: etcd 3.5
- **API Gateway**: Nginx (to be configured)

### Observability
- **Metrics**: Prometheus + Grafana
- **Tracing**: Jaeger + OpenTelemetry
- **Logging**: Structured JSON logs

### Development Tools
- **Containerization**: Docker + Docker Compose
- **Build Tool**: Make
- **Linting**: golangci-lint

## Common Commands

### Setup and Build
```bash
make setup          # Initialize infrastructure and dependencies
make build          # Build all services
make build-frontend # Build specific service
make deps           # Download Go dependencies
make tidy           # Tidy go.mod
```

### Running Services
```bash
make run SERVICE=frontend    # Run specific service
make dev SERVICE=frontend    # Start infra + run service
```

### Testing
```bash
make test              # Run all tests
make test-coverage     # Run tests with coverage report
```

### Docker Management
```bash
make docker-up         # Start all Docker services
make docker-down       # Stop Docker services
make docker-clean      # Stop and remove volumes
make docker-logs       # View logs
```

### Infrastructure Initialization
```bash
make kafka-init        # Create Kafka topics
make kafka-topics      # List Kafka topics
make minio-init        # Create MinIO bucket
```

### Code Quality
```bash
make fmt               # Format code
make lint              # Run linter
make clean             # Clean build artifacts
```
