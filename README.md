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
    data.go                 # Default quotes data
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

Terdapat 20 quote default yang sudah tersedia saat pertama kali server dijalankan.

## API Endpoints

| Method | Endpoint                | Description              |
|--------|-------------------------|--------------------------|
| GET    | /api/v1/quotes/daily    | Ambil quote hari ini     |
| POST   | /api/v1/admin/quotes    | Tambah quote baru        |
| GET    | /api/v1/admin/quotes    | Ambil semua quote        |

Lihat [API Documentation](docs/API.md) untuk detail lebih lanjut.

## License

MIT
