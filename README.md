# Backend Sistem Monitoring ISPU

## Deskripsi Proyek
Backend ini adalah layanan inti untuk sistem pemantauan Indeks Standar Pencemar Udara (ISPU). Sistem ini dirancang untuk mengumpulkan, mengelola, dan menyajikan data kualitas udara dari berbagai stasiun pemantauan secara *real-time*. Proyek ini dibangun menggunakan bahasa pemrograman Go (Golang) dengan arsitektur yang bersih dan skalabel.

## Fitur Utama
-   **Manajemen Stasiun Pemantauan**: CRUD (Create, Read, Update, Delete) untuk data stasiun pemantauan udara.
-   **Pencatatan Data Kualitas Udara**: Merekam parameter ISPU seperti PM10, PM2.5, SO2, CO, O3, dan NO2.
-   **Dashboard Analitik**: Menyediakan ringkasan statistik, distribusi kategori kualitas udara, dan tren data terkini.
-   **API RESTful**: Menyediakan antarmuka standar untuk komunikasi dengan frontend atau perangkat IoT.
-   **Kategorisasi Otomatis**: Penentuan kategori ISPU (Baik, Sedang, Tidak Sehat, dll.) berdasarkan nilai pengukuran.

## Teknologi yang Digunakan
-   **Bahasa Pemrograman**: [Go (Golang)](https://go.dev/) - Versi 1.21+
-   **Web Framework**: [Gin Web Framework](https://github.com/gin-gonic/gin) - Untuk routing dan handling HTTP request.
-   **Database**: [PostgreSQL](https://www.postgresql.org/) - Penyimpanan data relasional utama.
-   **ORM**: [GORM](https://gorm.io/) - Untuk interaksi database yang mudah dan aman.
-   **Caching**: [Redis](https://redis.io/) - Untuk caching data dan peningkatan performa.
-   **Konfigurasi**: [Godotenv](https://github.com/joho/godotenv) - Manajemen variabel lingkungan (.env).

## Struktur Proyek
Proyek ini mengikuti prinsip *Clean Architecture* untuk memudahkan pemeliharaan dan pengujian:
-   `cmd/`: Entry point aplikasi.
-   `internal/`: Kode privat aplikasi.
    -   `handler/`: Menangani HTTP request dan response.
    -   `service/`: Berisi logika bisnis utama.
    -   `repository/`: Berinteraksi langsung dengan database.
    -   `model/`: Definisi struktur data (structs).
    -   `middleware/`: Middleware HTTP (misal: CORS, Auth).
    -   `config/`: Konfigurasi aplikasi dan koneksi database.

## Cara Menjalankan Proyek

### Prasyarat
Pastikan Anda telah menginstal:
-   Go (versi 1.21 atau lebih baru)
-   PostgreSQL
-   Redis (opsional, tergantung konfigurasi)

### Instalasi
1.  **Clone repositori ini:**
    ```bash
    git clone <repository-url>
    cd backend
    ```

2.  **Setup Variabel Lingkungan:**
    Salin file `.env.example` menjadi `.env` dan sesuaikan konfigurasinya dengan environment lokal Anda.
    ```bash
    cp .env.example .env
    ```
    Pastikan pengaturan database (`DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_PORT`) sudah benar.

3.  **Install Dependencies:**
    ```bash
    go mod download
    ```

4.  **Jalankan Aplikasi:**
    ```bash
    go run cmd/main.go
    ```

## Endpoint API Utama
Berikut adalah beberapa endpoint utama yang tersedia:

-   `GET /api/v1/dashboard`: Mendapatkan ringkasan data dashboard.
-   `GET /api/v1/stations`: Mendapatkan daftar stasiun pemantauan.
-   `POST /api/v1/stations`: Menambahkan stasiun baru.
-   `GET /api/v1/air-quality`: Mendapatkan data kualitas udara.
-   `POST /api/v1/air-quality`: Mengirim data pengukuran baru.

---
Dibuat untuk keperluan presentasi dan dokumentasi teknis.
