# go-employee-analytics

`go-employee-analytics` adalah aplikasi backend berbasis **Golang** dan **PostgreSQL** yang digunakan untuk melakukan analisis data karyawan, pengolahan data berbasis query, serta pemrosesan logic dan algoritma dalam bentuk **HTTP REST API**.

Aplikasi ini menghasilkan **response JSON**, mendukung **ekspor data ke file**, serta menyediakan berbagai endpoint untuk kebutuhan analisis dan manipulasi data.

---

## Fitur Utama

- Analisis data karyawan aktif
- Filtering dan pengurutan data karyawan
- Agregasi data (selisih tanggal, total ulasan, dll)
- Proyeksi kenaikan gaji tahunan
- Ekspor hasil analisis ke file `.txt`
- Membaca kembali data dari file menjadi JSON
- Operasi logic berbasis array dan string
- Generator data acak (huruf & angka) beserta statistiknya

---

## Tech Stack

- Golang
- PostgreSQL
- GORM
- HTTP REST API
- Postman

---

## Struktur Project

```
go-employee-analytics/
├── cmd/
│   └── web/
│       └── main.go     # entry point aplikasi
├── internal/
│   ├── handler/        # HTTP handler
│   ├── service/        # business logic
│   ├── repository/     # akses database
│   └── model/          # struktur data
├── sql/
│   ├── schema.sql      # definisi tabel database
│   └── seed.sql        # data awal
├── output/             # file hasil ekspor
├── README.md
└── RUN_RESULT.pdf      # dokumentasi hasil eksekusi
```

---

## Menjalankan Aplikasi

Project ini dapat dijalankan dalam beberapa mode **CLI** sesuai kebutuhan development.

### 1️⃣ Clone Repository

```bash
git clone https://github.com/hackim18/go-employee-analytics.git
cd go-employee-analytics
```

### 2️⃣ Mode Eksekusi

#### 🟢 Full Mode (Drop → Migrate → Seed → Run)

```bash
go run cmd/web/main.go --drop-table --migrate --seed --run
```

#### 🟡 Migrate → Seed → Run

```bash
go run cmd/web/main.go --migrate --seed --run
```

#### 🔵 Migrate + Run

```bash
go run cmd/web/main.go --migrate --run
```

#### 🟣 Hanya Menjalankan Server

```bash
go run cmd/web/main.go
```

#### 🧪 Hot Reload (Air)

```bash
air
```

Pastikan file `.air.toml` tersedia.

Server akan berjalan di:

```
http://localhost:8080
```

---

## Implementasi Fitur & Endpoint API

Bagian ini menjelaskan implementasi fitur utama aplikasi dalam bentuk HTTP API dan response JSON.

### Analisis Karyawan Aktif Berdasarkan Kriteria Nama

```
GET /q2
```

### Karyawan Tanpa Data Ulasan

```
GET /q3
```

### Selisih Tanggal Perekrutan

```
GET /q4
```

### Proyeksi Kenaikan Gaji dan Total Ulasan

```
GET /q5
```

### Ekspor Data ke File

```
POST /q6/save
```

### Membaca Data dari File

```
GET /q7/file
```

### Validasi dan Pencarian Data Kota

```
GET /q8
```

### Operasi Data Array

```
/q9/*
```

### Generator Data Acak

```
GET /q10
```

---

## Output File

Hasil ekspor data disimpan pada folder:

```
/output
```

Setiap file berisi data dalam format JSON yang dapat dibaca kembali melalui API.

---

## Dokumentasi Hasil

Seluruh proses eksekusi aplikasi terdokumentasi dalam file:

```
RUN_RESULT.pdf
```

---

## Catatan

- Seluruh response API menggunakan format JSON.
- Aplikasi difokuskan pada backend logic dan data processing.
- Struktur project dirancang agar mudah dikembangkan lebih lanjut.

---

## Author

**Khakim**  
Backend Developer
