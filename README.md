# Stunting-Backend
**Tim: Lumino Dev**

Anggota:
- Nur Muhammad Ainul Yaqin
- Kurnia Cahya Febryanto
- Christhoper Marcelino Mamahit


## Deskripsi
StuntFree adalah platform monitoring tumbuh kembang anak usia dini dengan konsep To-Do-List sebagai solusi preventif stunting untuk orang tua.

## Struktur Folder dan File
```bash
.
|-- Dockerfile
|-- README.md
|-- cmd
|   `-- app
|       `-- main.go
|-- configs
|   `-- config.go
|-- go.mod
|-- go.sum
|-- pkg
|   |-- delivery
|   |   `-- http
|   |       `-- ${name}_handler.go
|   |-- middleware
|   |   `-- middleware.go
|   |-- models
|   |   `-- ${name}.go
|   |-- repository
|   |   `-- ${name}_repository.go
|   `-- usecase
|       `-- ${name}_usecase.go
`-- test
    `-- ${name}_usecase_test.go
     -- ${name}_handler_test.go
```

### Penjelasan Struktur Folder dan File Penting

- **Dockerfile**: File konfigurasi untuk membuat Docker image dari aplikasi ini.
- **cmd/app/main.go**: Titik awal aplikasi, di sini semua komponen disatukan dan server HTTP dimulai.
- **configs/config.go**: Mengandung fungsi untuk memulai dan mengkonfigurasi database.
- **pkg**: Direktori ini berisi semua kode utama aplikasi.
    - **delivery**: Menghandle request HTTP dan meresponsnya.
    - **middleware**: Mengandung middleware yang dapat digunakan di seluruh aplikasi.
    - **models**: Mendefinisikan model `${name}` dan interface untuk `${name}Repository` dan `${name}Usecase`.
    - **repository**: Mengatur semua interaksi dengan sumber data.
    - **usecase**: Tempat penulisan logika bisnis utama.
- **test**: Direktori ini berisi semua kode test aplikasi.

## Teknologi yang Digunakan
- Golang
- GORM (ORM library untuk Golang)
- MySQL
- Docker
- Git

## Cara Install
Pastikan Go, Docker dan MySQL sudah terinstall pada mesin Anda. Jika sudah, ikuti langkah-langkah berikut:
1. Clone repository ini.
2. Buka terminal dan masuk ke direktori tempat clone repository.
3. Jalankan `go mod tidy` untuk mendownload semua dependencies.
4. Buat file `.env` berdasarkan file `.env.example`, lalu sesuaikan dengan konfigurasi.
5. Jalankan `go run cmd/app/main.go` untuk memulai aplikasi.

## Cara Menggunakan
Setelah aplikasi berjalan, bisa mengakses API melalui endpoint yang sudah ditentukan.
Link API Postman yaitu https://documenter.getpostman.com/view/28422242/2s946bBui9

## Cara Running
Ada banyak metode
1. `go run cmd/app/main.go` untuk running biasa dengan default autoMigrate dan database local
2. `go run cmd/app/main.go --seed` untuk running dengan autoMigrate dan seed data
3. `go run cmd/app/main.go --prod` untuk running dengan database production PlanetScale
4. `go run cmd/app/main.go --seed --prod` untuk running dengan database production PlanetScale dan seed data

## Hak Cipta
© 2023 Kurnia Cahya Febryanto, Christhoper Marcelino Mamahit, dan Nur Muhammad Ainul Yaqin.

