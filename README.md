# Asbowo API

API sederhana untuk quote harian asbun dalam Prabowo Subianto.

## Tech Stack

- Go 1.26.5
- Gin (HTTP Framework)

## Project Structure

```
cmd/
  api/
    main.go                 # Entry point
internal/
  model/
    quote.go                # Data model
  repository/
    quote_repository.go     # Repository interface
    memory_quote_repository.go  # In-memory implementation
  service/
    quote_service.go        # Business logic
  handler/
    quote_handler.go        # HTTP handlers
```

## Getting Started

### Prerequisites

- Go 1.26.5 atau lebih tinggi

### Installation

```bash
git clone https://github.com/username/asbowo-api.git
cd asbowo-api
go mod download
```

### Run

```bash
go run cmd/api/main.go
```

Server akan berjalan di `http://localhost:8080`.

## API Endpoints

| Method | Endpoint                | Description              |
|--------|-------------------------|--------------------------|
| GET    | /api/v1/quotes/daily    | Ambil quote hari ini     |

Lihat [API Documentation](docs/API.md) untuk detail lebih lanjut.

## License

MIT
