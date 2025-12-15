<!-- @format -->

# Golang Microservices dengan gRPC

Tutorial lengkap untuk membangun microservices menggunakan gRPC di Golang. Project ini membuat 2 service sebagai contoh: **User Service** dan **Order Service**.

## 📋 Daftar Isi

- [Persyaratan](#persyaratan)
- [Struktur Project](#struktur-project)
- [Instalasi](#instalasi)
- [Cara Menjalankan](#cara-menjalankan)
- [Testing](#testing)
- [API Documentation](#api-documentation)
- [Arsitektur](#arsitektur)

## 🔧 Persyaratan

Sebelum memulai, pastikan sudah terinstall:

- **Go** (v1.21 atau lebih baru)

  ```bash
  go version
  ```

- **Protocol Buffers Compiler (protoc)**

  **macOS:**

  ```bash
  brew install protobuf
  ```

  **Linux (Ubuntu/Debian):**

  ```bash
  sudo apt update
  sudo apt install -y protobuf-compiler
  ```

  **Windows:**
  Download dari [GitHub Releases](https://github.com/protocolbuffers/protobuf/releases)

- **Go Plugins untuk protoc**

  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```

- **Tambahkan Go bin ke PATH**

  **macOS/Linux (bash):**

  ```bash
  echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.bashrc
  source ~/.bashrc
  ```

  **macOS (zsh):**

  ```bash
  echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.zshrc
  source ~/.zshrc
  ```

## 📁 Struktur Project

```
grpc-microservices/
├── proto/
│   ├── user/
│   │   ├── user.proto           # Proto definition untuk User Service
│   │   ├── user.pb.go           # Generated Go code
│   │   └── user_grpc.pb.go      # Generated gRPC code
│   └── order/
│       ├── order.proto          # Proto definition untuk Order Service
│       ├── order.pb.go          # Generated Go code
│       └── order_grpc.pb.go     # Generated gRPC code
├── user-service/
│   ├── main.go                  # Entry point User Service
│   └── server.go                # Business logic User Service
├── order-service/
│   ├── main.go                  # Entry point Order Service
│   ├── server.go                # Business logic Order Service
│   └── client.go                # gRPC client untuk User Service
├── test-client/
│   └── main.go                  # Client untuk testing semua service
├── generate.sh                  # Script untuk generate proto files
├── go.mod                       # Go module definition
├── go.sum                       # Go dependencies checksum
└── README.md                    # Dokumentasi project
```

## 🚀 Instalasi

### 1. Clone atau Buat Project

```bash
mkdir grpc-microservices
cd grpc-microservices
```

### 2. Initialize Go Module

```bash
go mod init grpc-microservices
```

### 3. Buat Struktur Folder

```bash
mkdir -p proto/user proto/order user-service order-service test-client
```

### 4. Install Dependencies

```bash
go get google.golang.org/grpc
go get google.golang.org/protobuf
go get github.com/google/uuid
```

### 5. Generate Proto Files

Buat file `generate.sh` di root project:

```bash
chmod +x generate.sh
./generate.sh
```

Atau generate manual:

```bash
# Set PATH
export PATH="$PATH:$(go env GOPATH)/bin"

# Generate User proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/user/user.proto

# Generate Order proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/order/order.proto
```

### 6. Tidy Dependencies

```bash
go mod tidy
```

## ▶️ Cara Menjalankan

Buka **3 terminal** berbeda dan jalankan setiap service:

### Terminal 1 - User Service

```bash
cd user-service
go run .
```

Output yang diharapkan:

```
User Service is running on port :50051
```

### Terminal 2 - Order Service

```bash
cd order-service
go run .
```

Output yang diharapkan:

```
Successfully connected to User Service
Order Service is running on port :50052
```

### Terminal 3 - Test Client

```bash
cd test-client
go run main.go
```

Output yang diharapkan:

```
=== Test 1: Create User ===
User created: id:"..." name:"John Doe" email:"john@example.com" ...

=== Test 2: Get User ===
User retrieved: ...

=== All tests completed successfully! ===
```

## 🧪 Testing

### Testing Manual dengan gRPC Client

Anda bisa menggunakan tools seperti:

1. **grpcurl** (Command line tool)

   ```bash
   # Install
   brew install grpcurl  # macOS

   # List services
   grpcurl -plaintext localhost:50051 list

   # Call CreateUser
   grpcurl -plaintext -d '{"name":"John","email":"john@example.com"}' \
     localhost:50051 user.UserService/CreateUser
   ```

2. **BloomRPC** atau **Postman** (GUI tool)

### Testing dengan Test Client

Jalankan test client yang sudah disediakan:

```bash
cd test-client
go run main.go
```

Test client akan melakukan:

1. ✅ Create User
2. ✅ Get User by ID
3. ✅ Create Order
4. ✅ Create Another Order
5. ✅ Get Order by ID
6. ✅ Get Orders by User ID
7. ✅ List All Users
8. ✅ List All Orders

## 📚 API Documentation

### User Service (Port: 50051)

#### CreateUser

```protobuf
rpc CreateUser (CreateUserRequest) returns (UserResponse)
```

**Request:**

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "phone": "08123456789"
}
```

**Response:**

```json
{
  "user": {
    "id": "uuid-generated",
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "08123456789"
  },
  "message": "User created successfully"
}
```

#### GetUser

```protobuf
rpc GetUser (GetUserRequest) returns (UserResponse)
```

**Request:**

```json
{
  "id": "user-uuid"
}
```

#### ListUsers

```protobuf
rpc ListUsers (ListUsersRequest) returns (ListUsersResponse)
```

**Request:**

```json
{
  "page": 1,
  "limit": 10
}
```

### Order Service (Port: 50052)

#### CreateOrder

```protobuf
rpc CreateOrder (CreateOrderRequest) returns (OrderResponse)
```

**Request:**

```json
{
  "user_id": "user-uuid",
  "product_name": "Laptop",
  "quantity": 1,
  "total_price": 15000000
}
```

**Response:**

```json
{
  "order": {
    "id": "uuid-generated",
    "user_id": "user-uuid",
    "product_name": "Laptop",
    "quantity": 1,
    "total_price": 15000000,
    "status": "pending"
  },
  "message": "Order created successfully"
}
```

#### GetOrder

```protobuf
rpc GetOrder (GetOrderRequest) returns (OrderResponse)
```

#### GetOrdersByUserId

```protobuf
rpc GetOrdersByUserId (GetOrdersByUserIdRequest) returns (ListOrdersResponse)
```

**Request:**

```json
{
  "user_id": "user-uuid"
}
```

#### ListOrders

```protobuf
rpc ListOrders (ListOrdersRequest) returns (ListOrdersResponse)
```

## 🏗️ Arsitektur

### User Service (Port 50051)

- **Fungsi:** Mengelola data user (create, get, list)
- **Storage:** In-memory (map) - data hilang saat restart
- **ID Generator:** UUID v4
- **Concurrency:** Menggunakan sync.RWMutex untuk thread-safety

### Order Service (Port 50052)

- **Fungsi:** Mengelola data order (create, get, list, get by user)
- **Storage:** In-memory (map)
- **Komunikasi:** Dapat berkomunikasi dengan User Service via gRPC client
- **Verifikasi:** Bisa memverifikasi user existence sebelum create order

### Komunikasi Antar Service

```
┌─────────────────┐
│   Test Client   │
└────────┬────────┘
         │
         ├─────────────────┐
         │                 │
         ▼                 ▼
┌─────────────────┐ ┌─────────────────┐
│  User Service   │ │  Order Service  │
│   Port: 50051   │ │   Port: 50052   │
└─────────────────┘ └────────┬────────┘
         ▲                   │
         └───────────────────┘
         gRPC Client Connection
```

Order Service memiliki gRPC client untuk berkomunikasi dengan User Service. Ini menunjukkan pola **service-to-service communication** di microservices.

### Protocol Buffers

Protocol Buffers (protobuf) adalah format serialisasi data yang:

- **Lebih kecil** dibanding JSON
- **Lebih cepat** untuk serialize/deserialize
- **Type-safe** dengan strong typing
- **Language-agnostic** - bisa generate code untuk berbagai bahasa
- **Backward compatible** dengan versioning

## 📝 License

MIT License - feel free to use this project for learning purposes.

## 🤝 Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

---

**Happy Coding! 🚀**

Made with ❤️ using Go and gRPC
