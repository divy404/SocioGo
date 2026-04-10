<div align="center">

# SocioGo

**A blazing-fast, production-grade social backend engineered in Go.**

[![Go Version](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go&logoColor=white)](https://golang.org/)
[![Chi Router](https://img.shields.io/badge/Router-chi%20v5-blueviolet)](https://github.com/go-chi/chi)
[![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL%2016-336791?logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?logo=docker&logoColor=white)](https://docs.docker.com/compose/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

</div>

---

SocioGo is a clean, modular REST API backend for social networking applications. Built with idiomatic Go, it follows a layered architecture with strict separation between routing, business logic, and data access — designed to scale without getting in the way.

---

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Environment Setup](#environment-setup)
  - [Running with Docker](#running-with-docker)
  - [Running Locally](#running-locally)
- [Database Migrations](#database-migrations)
- [Development Workflow](#development-workflow)
- [API Overview](#api-overview)
- [Configuration Reference](#configuration-reference)
- [Contributing](#contributing)

---

## Features

- **RESTful API** built with the lightweight, idiomatic [`chi`](https://github.com/go-chi/chi) router
- **PostgreSQL** as the primary datastore, accessed via the native `lib/pq` driver
- **Structured logging** using `go.uber.org/zap` — zero allocations, production-ready
- **Request validation** powered by `go-playground/validator` with custom rules
- **UUID-based identifiers** using `google/uuid` for all primary keys
- **Password hashing** with `bcrypt` via `golang.org/x/crypto`
- **SQL migrations** managed with `golang-migrate` and a clean Makefile interface
- **Database seeding** for rapid local development (`make seed`)
- **Live reload** in development via [Air](https://github.com/air-verse/air)
- **Docker Compose** for one-command infrastructure setup
- **Clean architecture** with `internal/` isolation — no business logic leaking into routing or storage layers

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.24 |
| Router | `go-chi/chi` v5 |
| Database | PostgreSQL 16 |
| DB Driver | `lib/pq` |
| Validation | `go-playground/validator` v10 |
| Logging | `go.uber.org/zap` |
| UUIDs | `google/uuid` |
| Auth (passwords) | `golang.org/x/crypto` (bcrypt) |
| Config | `joho/godotenv` |
| Migrations | `golang-migrate` |
| Dev Hot-Reload | Air |
| Containerization | Docker Compose |

---

## Project Structure

```
SocioGo/
├── cmd/
│   ├── api/                  # Application entrypoint (main.go)
│   └── migrate/
│       ├── migrations/       # SQL migration files (up/down)
│       └── seed/             # Database seed script
├── internal/                 # Private application code (not importable externally)
│   ├── app/                  # Server wiring, router setup, middleware
│   ├── handler/              # HTTP handlers — decode request, call service, encode response
│   ├── service/              # Business logic layer
│   ├── store/                # Data access layer (SQL queries)
│   └── model/                # Domain types and DTOs
├── scripts/                  # Utility scripts
├── bin/                      # Compiled binaries (git-ignored)
├── .air.toml                 # Air live-reload configuration
├── .env.sample               # Environment variable template
├── docker-compose.yml        # PostgreSQL service definition
├── go.mod
├── go.sum
└── Makefile                  # Developer task runner
```

The `internal/` package enforces Go's visibility rules — nothing inside leaks outside the module. Handlers only talk to services; services only talk to the store. Dependencies flow inward.

---

## Getting Started

### Prerequisites

- [Go 1.24+](https://golang.org/dl/)
- [Docker & Docker Compose](https://docs.docker.com/get-docker/)
- [`golang-migrate` CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [Air](https://github.com/air-verse/air) *(optional, for live reload)*

```bash
# Install golang-migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Install Air
go install github.com/air-verse/air@latest
```

### Environment Setup

Copy the sample environment file and fill in your values:

```bash
cp .env.sample .env
```

Edit `.env`:

```env
ADDR=:8080
DB_ADDR=postgres://admin:adminpassword@localhost:5433/socio?sslmode=disable
```

> The `DB_ADDR` above matches the Docker Compose defaults. Change credentials for any non-local environment.

### Running with Docker

Start PostgreSQL:

```bash
docker-compose up -d
```

This spins up a PostgreSQL 16 container (`socio` database, `admin` user) exposed on port **5433**.

### Running Locally

```bash
# Install dependencies
go mod download

# Run migrations
make migrate-up

# Seed the database (optional)
make seed

# Start the server
go run ./cmd/api
```

The API will be available at `http://localhost:8080`.

---

## Database Migrations

Migrations live in `cmd/migrate/migrations/` as sequential `.up.sql` / `.down.sql` file pairs, managed via the Makefile.

```bash
# Apply all pending migrations
make migrate-up

# Roll back the last migration batch
make migrate-down 1

# Create a new migration (replace <name> with a descriptive name)
make migrate-create <name>

# Force a specific migration version (use with caution)
make migrate-force <version>

# Seed the database with initial data
make seed
```

---

## Development Workflow

For live reload during development, use Air:

```bash
air
```

Air watches for `.go`, `.html`, `.tpl`, and `.tmpl` file changes, recompiles to `./bin/main`, and restarts the server automatically. Configuration is in `.air.toml`.

To do a manual build:

```bash
go build -o ./bin/main ./cmd/api
./bin/main
```

---

## API Overview

> Full API documentation coming soon. The base route structure follows RESTful conventions.

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/v1/users` | Register a new user |
| `POST` | `/v1/auth/login` | Authenticate and receive a token |
| `GET` | `/v1/users/:id` | Get user profile |
| `PUT` | `/v1/users/:id` | Update user profile |
| `POST` | `/v1/posts` | Create a new post |
| `GET` | `/v1/posts/:id` | Get a post by ID |
| `DELETE` | `/v1/posts/:id` | Delete a post |
| `POST` | `/v1/posts/:id/comments` | Add a comment to a post |
| `POST` | `/v1/users/:id/follow` | Follow a user |
| `DELETE` | `/v1/users/:id/follow` | Unfollow a user |

---

## Configuration Reference

All configuration is loaded from the `.env` file via `godotenv` at startup.

| Variable | Description | Example |
|----------|-------------|---------|
| `ADDR` | Server listen address | `:8080` |
| `DB_ADDR` | Full PostgreSQL connection string | `postgres://admin:pass@localhost:5433/socio?sslmode=disable` |

---

## Contributing

Contributions are welcome. Please open an issue first to discuss significant changes.

```bash
# Fork the repo, then clone your fork
git clone https://github.com/<your-username>/SocioGo.git
cd SocioGo

# Create a feature branch
git checkout -b feature/your-feature-name

# Make changes, then push
git push origin feature/your-feature-name

# Open a Pull Request on GitHub
```

Please keep PRs focused — one feature or fix per PR. All new endpoints should have corresponding migration files.

---

<div align="center">
  Built with Go · Made by <a href="https://github.com/divy404">Divyansh Jain</a>
</div>
