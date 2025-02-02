# Akatech

## Membuat REST API
### Running

- Up Compose (optional)
```bash
docker compose -f postgresql.yaml up -d
```
- Change .env
```bash
cp .env.example .env
```
- Run Go Application
```bash
go run main.go
```
### 1. **Error Handling dan Logging:**
- **Database Connection:**
    - Menggunakan `panic` untuk menghentikan aplikasi jika gagal terhubung ke database PostgreSQL.
- **API Endpoint:**
    - **JSON Binding:** Error ditangani jika JSON tidak valid dengan status `400`.
    - **Validasi Data:** Memastikan `name` dan `email` tidak kosong. Error jika kosong, status `400`.
    - **Duplikasi Email:** Mengecek email yang sudah ada di database. Jika duplikat, status `500`.
    - **Penyimpanan Data:** Jika gagal menyimpan user, akan mengembalikan error status `500`.

- **Logging:**
    - Gin Gonic Support

### 2. **Menghindari SQL Injection:**
- GORM menggunakan **parameterized queries** untuk menghindari SQL Injection. Contoh:
  ```go
  db.Where("field = ?", value).First(&model)
