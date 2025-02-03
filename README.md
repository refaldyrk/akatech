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

## SOAP
### How To Use?
Cek File `wsdl.go` lalu pakai function nya seperti dibawah ini
```go
func main() {
	userID := "12345"
	response, err := callSOAPService(userID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User Details:\n")
	fmt.Printf("ID: %s\n", response.User.ID)
	fmt.Printf("First Name: %s\n", response.User.FirstName)
	fmt.Printf("Last Name: %s\n", response.User.LastName)
	fmt.Printf("Email: %s\n", response.User.Email)
}
```

## RabbitMQ
- Membuat Sender
```go
//Send Message
	err = u.channel.PublishWithContext(context.Background(), "", u.queue.Name, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(newUser.UserID),
	})

	if err != nil {
		return 0, nil
	}
```
- Membuat Consumer
```go 
msgs, err := configBase.AMQPChannel.Consume(
		configBase.Q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
```

Saat membuat user baru maka akan mengirimkan message baru ke RabbitMQ, dan otomatis akan di consume juga pada bersamaan\n
Example:
```shell
2025/02/03 10:01:21 Received a message: cug330dp6e2sk7fqdka0
```