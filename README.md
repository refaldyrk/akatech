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

## Unit Test
### Penjelasan
- Anda Bisa Jalankan Unit Test Dengan command berikut 
```bash
go test -v ./utest
```
maka anda melihat seperti ini
```shell
=== RUN   TestAdd
=== RUN   TestAdd/Add
=== RUN   TestAdd/Add#01
=== RUN   TestAdd/Add#02
=== RUN   TestAdd/Add#03
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/Add (0.00s)
    --- PASS: TestAdd/Add#01 (0.00s)
    --- PASS: TestAdd/Add#02 (0.00s)
    --- PASS: TestAdd/Add#03 (0.00s)
=== RUN   TestDivide
=== RUN   TestDivide/Divide
=== RUN   TestDivide/Divide#01
=== RUN   TestDivide/Divide#02
=== RUN   TestDivide/Divide#03
--- PASS: TestDivide (0.00s)
    --- PASS: TestDivide/Divide (0.00s)
    --- PASS: TestDivide/Divide#01 (0.00s)
    --- PASS: TestDivide/Divide#02 (0.00s)
    --- PASS: TestDivide/Divide#03 (0.00s)
PASS
ok      akatech/utest   0.348s
```
- Pentingnya Unit Test
    1. **Meningkatkan Keandalan:** Memastikan setiap fungsi berfungsi sesuai harapan, mengurangi kemungkinan bug.
    2. **Deteksi Masalah Dini:** Menemukan dan memperbaiki masalah sejak awal, sebelum mencapai tahap produksi.
    3. **Memudahkan Refactoring:** Membantu pengembang mengubah kode tanpa merusak fungsionalitas yang ada.
    4. **Meningkatkan Kepercayaan:** Memberikan rasa aman saat mengubah atau menambahkan fitur baru.
    5. **Dokumentasi Kode:** Bertindak sebagai dokumentasi untuk bagaimana fungsi seharusnya bekerja.