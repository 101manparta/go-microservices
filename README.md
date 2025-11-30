# Go Microservices Demo

Proyek ini adalah demo arsitektur microservices sederhana menggunakan bahasa Go.  
Terdiri dari beberapa service yang berjalan secara terpisah dan API Gateway sebagai pintu masuk utama.

---

## Struktur Project

go-microservices-demo/
├── gateway/
│ └── main.go # API Gateway service
├── product-service/
│ └── main.go # Service produk dengan endpoint GET dan POST /products
├── user-service/
│ └── main.go # Service user dengan endpoint GET dan POST /users (contoh)
└── README.md

yaml
Copy code

---

## Cara Setup & Menjalankan

### Prasyarat

- Go (minimal versi 1.18)
- Git (opsional, untuk clone repo)
- Curl atau Postman untuk testing API

---

### Clone Repository (jika belum)

```bash
git clone https://github.com/username/go-microservices-demo.git
cd go-microservices-demo
Menjalankan Service
1. Product Service
bash
Copy code
cd product-service
go run main.go
Service ini berjalan di port 8082.
Endpoint utama:

GET /products - Mendapatkan list produk

POST /products - Menambah produk baru (kirim JSON body)

2. User Service
bash
Copy code
cd user-service
go run main.go
Service ini berjalan di port 8081.
Endpoint utama:

GET /users

POST /users

3. API Gateway
bash
Copy code
cd gateway
go run main.go
Gateway berjalan di port 8080, bertindak sebagai proxy ke service di atas.

Contoh Request API
Product Service
GET semua produk

bash
Copy code
curl http://localhost:8082/products
POST produk baru

bash
Copy code
curl -X POST http://localhost:8082/products \
 -H "Content-Type: application/json" \
 -d '{"id":"1","name":"Produk A","price":15000}'
User Service
GET semua user

bash
Copy code
curl http://localhost:8081/users
POST user baru

bash
Copy code
curl -X POST http://localhost:8081/users \
 -H "Content-Type: application/json" \
 -d '{"id":"1","name":"User A"}'
API Gateway
Bisa juga akses endpoint lewat gateway, misalnya:

bash
Copy code
curl http://localhost:8080/products
curl http://localhost:8080/users
Struktur Kode Singkat
Setiap service menggunakan net/http package Go untuk membuat HTTP server sederhana.

Data disimpan di memory menggunakan map dan dilindungi sync.RWMutex.

API Gateway menggunakan reverse proxy (atau custom handler) untuk meneruskan request ke masing-masing service.

Tips
Jalankan service secara terpisah di terminal berbeda.

Pastikan port tidak bentrok (misal Jenkins atau aplikasi lain).

Gunakan tools seperti Postman untuk eksplorasi API lebih mudah.

Lisensi
Proyek ini dibagikan tanpa lisensi khusus (public domain). Gunakan dan modifikasi bebas.

