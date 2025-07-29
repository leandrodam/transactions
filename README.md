# transactions
A REST API test practice case to manage transactions.
Developed in Go following Ports and Adapters (Hexagonal) architecture using MySQL as primary database.

## Technologies Used

- **Go** - Main programming language
- **MySQL** - Relational database
- **Goose** - Migrations tool
- **Docker** - Application containerization
- **Docker Compose** - Container orchestration

## Prerequisites

Before running the project, make sure you have installed:

- [Docker](https://docs.docker.com/get-docker/) (version 20.10 or higher)
- [Docker Compose](https://docs.docker.com/compose/install/) (version 2.0 or higher)
- [Make](https://www.gnu.org/software/make/) (optional, but recommended)

## Environment Setup

### 1. Environment Variables

The project uses environment variables for configuration. Follow these steps:
1. Copy the example configuration file:
   ```bash
   cp example.env .env
   ```

2. Edit the `.env` file with your configurations:
   ```bash
   nano .env
   ```

The `example.env` file contains all necessary variables. Make sure to adjust the values according to your environment:

- Application name, environment and port
- MySQL database configurations

### 2. Project Structure

```
.
├── cmd/
│   └── api/                   # Application entry point
├── internal/                  # Internal application code
│   ├── adapters/              # Adapters
│   │   ├── http/              # Primary adapters - HTTP handlers
│   │   └── repository/        # Secondary adapters - Database repositories
│   ├── domain/                # Core domain entities
│   ├── infrastructure/        # Infrastructure layer
│   │   ├── config/            # Configuration management
│   │   ├── exceptions/        # Custom exceptions and error handling
│   │   ├── server/            # HTTP server setup
│   │   └── validator/         # Input validation
│   └── usecases/              # Application services - Ports implementation
├── scripts/                   # Utility scripts
│   ├── db/                    # Database scripts
│   │   ├── init-db/           # Database initialization scripts
│   │   └── migrations/        # Database migrations
│   │       └── mysql/         # MySQL specific migrations
│   ├── docker/                # Docker related scripts
│   │   └── api/               # API container scripts
│   │       └── Dockerfile     # API image
│   └── docker-compose.yml     # Container orchestration configuration
├── Dockerfile                 # Application image
├── Makefile                   # Automated commands
├── example.env                # Environment variables example file
├── .env                       # Environment variables (create from example.env)
└── README.md                  # This file
```

## Running with Docker

### Using Makefile (Recommended)

The project includes a Makefile with useful commands for development:

```bash
# Start services with docker compose
make up

# Start database service image only
make up-db

# Start services rebuilding images
make up-build

# Stop docker compose services
make down

# Restart docker compose services
make restart

# Show docker compose services logs
make logs

# List docker compose services
make ps

# Run tests
make test

# Format Go code
make fmt

# Run linter
make lint

# Run go mod tidy and vendor
make vendor

# Show the help menu with all commands
make help
```

### Using Docker Compose Directly

If you prefer to use Docker Compose directly:

```bash
# Initialize containers
docker compose -f ./scripts/docker/docker-compose.yml --env-file .env up -d

# Stop containers
docker-compose -f ./scripts/docker/docker-compose.yml --env-file .env down

# Rebuild images
docker-compose -f ./scripts/docker/docker-compose.yml --env-file .env build
```

## Development

### Running Locally

For local development without Docker:

1. Configure environment variables
2. Make sure MySQL are running
3. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

## API Usage

### Main Endpoints

The API will be available at `http://localhost:8080` (or the port configured in `.env`).

For more details check the OpenAPI file definitions at `./openapi/openapi.yml`
