# Asbowo API Documentation

Base URL: `http://localhost:8080`

---

## Endpoints

### GET `/api/v1/quotes/daily`

Mengembalikan quote random berdasarkan tanggal hari ini. Quote yang sama akan dikembalikan untuk tanggal yang sama (deterministik).

**Response**

- **200 OK**
```json
{
  "id": 1,
  "text": "Ndhasmu etik!",
  "created_at": "0001-01-01T00:00:00Z"
}
```

- **500 Internal Server Error**
```json
{
  "error": "belum ada asbun yang tersimpan"
}
```

---

## Models

### Quote

| Field      | Type     | Description                    |
|------------|----------|--------------------------------|
| id         | int      | Unique identifier              |
| text       | string   | Isi quote                      |
| created_at | datetime | Waktu pembuatan quote (UTC)    |

---

## Notes

- Quote dipilih berdasarkan seed tanggal (`YYYYMMDD`), sehingga hasilnya konsisten dalam satu hari
- Saat ini data bersifat **in-memory** (akan hilang saat restart)
