# Mentorship CRUD App

A robust Document Management System built in Go, featuring a RESTful CRUD API for managing documents. Data is persisted in PostgreSQL running inside Docker, with a modern, production-ready project structure and developer tooling.

## Features

- **CRUD API**: Create, Read, Update, Delete documents
- **Document Fields**: ID, Title, Author, Content, Created_at, Updated_at
- **PostgreSQL**: Persistent storage in Dockerized PostgreSQL
- **GIN**: Fast, idiomatic HTTP web framework for Go
- **GORM**: ORM for elegant database interactions
- **Cobra CLI**: Run server, migrations, and seeder via command-line
- **Migrations**: Managed with golang-migrate
- **Logging**: Structured logging with logrus
- **Dockerized**: Easy local and live environments
- **Testing**: Ready for unit, integration, and load testing (k6)

## Tech Stack

- **Go** (Golang)
- **GIN** (HTTP API)
- **GORM** (ORM)
- **Cobra** (CLI)
- **golang-migrate** (DB migrations)
- **logrus** (Logging)
- **PostgreSQL** (Database)
- **Docker & Docker Compose**
- **k6** (Load testing)

## Getting Started

### Prerequisites
- Go 1.20+
- Docker & Docker Compose

### Local Development

```bash
# Start local environment (PostgreSQL)
make run-local

# Run the server locally
make run-server

# Run database migrations
make migrate

# Seed the database
make seed

# Tear down local environment
make down-local
```

### Live Environment (Docker Compose)

```bash
# Start live environment
make run-live

# Tear down live environment
make down-live
```

## API Endpoints

- `POST   /documents`      — Create a new document
- `PUT    /documents/:id`  — Update a document
- `DELETE /documents/:id`  — Delete a document
- `GET    /documents/:id`  — Get a document by ID
- `GET    /documents`      — List all documents

## Project Structure

```
mentorship-crud-app/
├── cmd/                 # CLI entrypoints (server, migrate, seed)
├── internal/pkg/        # Internal packages (server, handlers, models)
├── data/seeders/        # Seeder logic
├── Dockerfile           # App Dockerfile
├── docker-compose.yaml  # Local environment
├── docker-compose-live.yaml # Live/production environment
├── makefile             # Dev workflow commands
└── README.md            # Project documentation
```

## Development Workflow

- Use `make` commands for all development tasks
- Use Cobra CLI for running server, migrations, and seeding
- All configs via environment variables (`.env`)
- Logging output is JSON-formatted for easy parsing

## Testing

- Unit & integration tests: `go test ./...`
- Load test: [k6](https://k6.io/) scripts (to be added)

## Contributing

1. Fork the repo and create your feature branch (`git checkout -b feature/your-feature`)
2. Commit your changes (`git commit -am 'Add new feature'`)
3. Push to the branch (`git push origin feature/your-feature`)
4. Open a Pull Request

## License

MIT License © 2025 Ahmed Mohamed
