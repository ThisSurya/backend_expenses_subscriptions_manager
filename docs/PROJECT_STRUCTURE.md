# Struktur Proyek Backend Terbaru

Dokumen ini merangkum struktur folder dan file utama pada backend **Smart Expenses & Subscriptions Manager** berdasarkan kondisi workspace saat ini.

## Gambaran Umum

Backend ini menggunakan pola berlapis:

- `routes` untuk definisi endpoint
- `controllers` untuk HTTP handler
- `services` untuk business logic
- `repository` untuk akses data ke database
- `models` untuk entitas dan konfigurasi database
- `requests` untuk DTO payload request
- `utils` untuk helper umum

## Struktur Folder Saat Ini

```text
backend/
├── main.go
├── go.mod
├── go.sum
├── Makefile
├── .env
├── config/
│   └── config.go
├── controllers/
│   ├── auth_controller.go
│   ├── category_controller.go
│   ├── expense_controller.go
│   └── subscription_controller.go
├── docs/
│   ├── ARCHITECTURE.md
│   └── PROJECT_STRUCTURE.md
├── models/
│   ├── category.go
│   ├── expense.go
│   ├── notification.go
│   ├── refresh_token.go
│   ├── subscription.go
│   ├── user.go
│   ├── config/
│   │   ├── database.go
│   │   └── postgres.go
│   └── migrations/
│       ├── 000001_create_users_table.down.sql
│       ├── 000001_create_users_table.up.sql
│       ├── 000002_create_categories_table.down.sql
│       ├── 000002_create_categories_table.up.sql
│       ├── 000003_create_expenses_table.down.sql
│       ├── 000003_create_expenses_table.up.sql
│       ├── 000004_create_subscriptions_table.down.sql
│       ├── 000004_create_subscriptions_table.up.sql
│       ├── 000005_create_notifications_table.down.sql
│       ├── 000005_create_notifications_table.up.sql
│       ├── 000006_create_refresh_tokens_table.down.sql
│       └── 000006_create_refresh_tokens_table.up.sql
├── repository/
│   ├── category_repository.go
│   ├── errors.go
│   ├── expense_repository.go
│   ├── subscription_repository.go
│   └── user_repository.go
├── requests/
│   ├── category_request.go
│   ├── expense_request.go
│   ├── subscription_request.go
│   └── user_request.go
├── routes/
│   ├── routes.go
│   └── middleware/
│       └── auth.go
├── services/
│   ├── category_service.go
│   ├── errors.go
│   ├── expense_service.go
│   ├── expense_service_test.go
│   ├── subscription_service.go
│   ├── subscription_service_test.go
│   ├── user_service.go
│   └── user_service_test.go
└── utils/
    ├── err.go
    ├── hash.go
    ├── response.go
    ├── session.go
    └── validator.go
```

## Fungsi Tiap Folder

### `main.go`
Entry point aplikasi. Di sini konfigurasi dimuat, koneksi database dibuat, middleware dipasang, lalu route publik dan protected didaftarkan.

### `config/`
Berisi konfigurasi aplikasi umum, terutama pembacaan environment variable dan setup awal runtime.

### `controllers/`
Menangani request HTTP, validasi input, mengambil data dari context, lalu memanggil service yang sesuai.

### `services/`
Menyimpan business logic utama. Layer ini menjadi penghubung antara controller dan repository.

### `repository/`
Berisi akses data ke database melalui GORM. Layer ini fokus pada query dan operasi persistence.

### `models/`
Berisi struktur entitas domain dan konfigurasi database.

- File model seperti `user.go`, `expense.go`, dan `subscription.go` mewakili tabel/domain utama.
- `models/config/` menyimpan inisialisasi koneksi database.
- `models/migrations/` menyimpan SQL migration `up` dan `down`.

### `requests/`
Berisi struktur payload request untuk binding dan validasi input dari client.

### `routes/`
Menyimpan definisi endpoint API dan middleware yang digunakan pada routing.

### `utils/`
Helper umum seperti response formatter, hashing password, session helper, dan validator error formatter.

### `docs/`
Berisi dokumentasi arsitektur dan struktur proyek.

## Alur Dependency Utama

Urutan dependensi yang digunakan di backend ini adalah:

```text
Route -> Controller -> Service -> Repository -> Model/Database
```

## Ringkasan Endpoint Utama

Berdasarkan routing yang ada saat ini, API utama dikelompokkan menjadi:

- Auth: `/v1/api/register`, `/v1/api/login`
- Expenses: `/v1/api/expenses`
- Categories: `/v1/api/categories`
- Subscriptions: `/v1/api/subscriptions`

## Catatan Pemeliharaan

Jika ada folder, file model, repository, atau service baru, dokumentasi ini sebaiknya diperbarui agar tetap sesuai dengan struktur terbaru.
