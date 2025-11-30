## Project Structure

```
.
├── cmd/                          # Service entry points
│   ├── frontend/                 # Frontend API service
│   ├── auth/                     # Authentication service
│   ├── file-service/             # File upload service
│   ├── router-worker/            # Message routing worker
│   └── presence/                 # Presence tracking service
├── internal/                     # Internal packages
│   ├── models/                   # Domain models
│   │   ├── user.go
│   │   ├── conversation.go
│   │   ├── message.go
│   │   └── file.go
│   └── shared/                   # Shared utilities
│       ├── config/               # Configuration management
│       ├── errors/               # Error handling
│       ├── utils/                # Utility functions
│       ├── health/               # Health check handlers
│       └── logger/               # Structured logging
├── config/                       # Configuration files
│   ├── prometheus.yml
│   └── grafana/
├── scripts/                      # Utility scripts
│   └── mongo-init.js
├── docs/                         # Documentation
│   └── SETUP.md
├── .kiro/                        # Kiro configuration
│   ├── specs/                    # Feature specifications
│   └── steering/                 # Development guidelines
├── docker-compose.yml            # Docker services
├── Makefile                      # Build commands
├── go.mod                        # Go dependencies
├── .env.example                  # Environment template
└── README.md                     # Project overview
```

### Code Organization

- **cmd/**: Each service has its own directory with a `main.go` entry point
- **internal/models/**: Shared data models used across services
- **internal/shared/**: Common utilities, configuration, and error handling
- Services will have their own packages under `internal/` as they are implemented

### Development Approach

This project follows specification-driven development:
1. Define requirements and design in spec files
2. Iterate on specifications before implementation
3. Implement features based on finalized specs
4. Each task in `tasks.md` corresponds to specific implementation work
