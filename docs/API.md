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

### POST `/api/v1/admin/quotes`

Menambahkan quote baru.

**Request Body**

| Field | Type   | Required | Description     |
|-------|--------|----------|-----------------|
| text  | string | Yes      | Isi quote baru  |

**Request Example**
```json
{
  "text": "Quote baru yang keren!"
}
```

**Response**

- **201 Created**
```json
{
  "id": 21,
  "text": "Quote baru yang keren!",
  "created_at": "2026-09-03T10:00:00Z"
}
```

- **400 Bad Request**
```json
{
  "error": "Format JSON salah atau text kosong"
}
```

- **500 Internal Server Error**
```json
{
  "error": "Terdapat kesalahan saat menyimpan quote baru. Silahkan coba lagi nanti"
}
```

---

### GET `/api/v1/admin/quotes`

Mengembalikan seluruh quote yang tersimpan.

**Response**

- **200 OK**
```json
[
  {
    "id": 1,
    "text": "Ndhasmu etik!",
    "created_at": "2026-09-03T00:00:00Z"
  },
  {
    "id": 2,
    "text": "Omon-omon saja.",
    "created_at": "2026-09-03T00:00:00Z"
  }
]
```

- **500 Internal Server Error**
```json
{
  "error": "Terdapat kesalahan saat mengambil quotes. Silahkan coba lagi nanti"
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
- Terdapat 20 quote default yang sudah tersedia saat pertama kali server dijalankan
- Semua pesan error menggunakan Bahasa Indonesia
