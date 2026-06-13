# 🚀 Go User API

A production-grade RESTful API built with **Go**, **GoFiber**, **PostgreSQL**, **SQLC**, **Uber Zap**, and **go-playground/validator**.

Manages users with `name` and `dob` (date of birth) — age is calculated **dynamically** at request time using Go's `time` package.

---

## 🛠 Tech Stack

| Layer        | Technology                     |
|--------------|-------------------------------|
| HTTP Server  | GoFiber v2                    |
| Database     | PostgreSQL                    |
| DB Layer     | SQLC (generated queries)      |
| Logging      | Uber Zap                      |
| Validation   | go-playground/validator v10   |
| Config       | godotenv                      |
| Migrations   | golang-migrate                |
| Docker       | Docker + Docker Compose       |

---

## 📁 Project Structure

```
go-user-api/
├── cmd/
│   └── server/
│       └── main.go           # Entry point
├── config/
│   └── config.go             # Config loader
├── db/
│   ├── migrations/           # SQL migration files
│   ├── query/                # SQLC query definitions
│   └── sqlc/                 # SQLC generated code
├── internal/
│   ├── handler/              # HTTP handlers
│   ├── repository/           # Data access layer
│   ├── service/              # Business logic + age calculation
│   ├── routes/               # Route registration
│   ├── middleware/           # Request ID + Duration middleware
│   ├── models/               # Request/Response structs
│   └── logger/               # Uber Zap wrapper
├── .env                      # Environment variables
├── sqlc.yaml                 # SQLC config
├── Dockerfile                # Multi-stage Docker build
├── docker-compose.yml        # App + PostgreSQL services
└── README.md
```

---

## ⚙️ Setup

### 1. Prerequisites

- Go 1.22+
- PostgreSQL 15+
- [sqlc](https://sqlc.dev/): `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
- [golang-migrate](https://github.com/golang-migrate/migrate): `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

### 2. Clone & Install

```bash
git clone https://github.com/yourusername/go-user-api.git
cd go-user-api
go mod tidy
```

### 3. Configure Environment

Edit `.env` with your PostgreSQL credentials:

```env
APP_PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=userdb
```

### 4. Create Database

```sql
CREATE DATABASE userdb;
```

### 5. Run Migrations

```bash
migrate -path db/migrations \
  -database "postgres://postgres:password@localhost:5432/userdb?sslmode=disable" up
```

### 6. Generate SQLC Code (if modifying queries)

```bash
sqlc generate
```

### 7. Run the Server

```bash
go run cmd/server/main.go
```

Server starts at `http://localhost:3000`

---

## 🐳 Docker

```bash
# Start both app and PostgreSQL
docker-compose up --build

# Stop
docker-compose down
```

---

## 📡 API Reference

### POST /users — Create User

```bash
curl -X POST http://localhost:3000/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice", "dob": "1990-05-10"}'
```

**Response `201`:**
```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}
```

---

### GET /users/:id — Get User by ID

```bash
curl http://localhost:3000/users/1
```

**Response `200`:**
```json
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}
```

---

### GET /users — List All Users (with Pagination)

```bash
curl "http://localhost:3000/users?page=1&limit=10"
```

**Response `200`:**
```json
[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 35
  }
]
```

---

### PUT /users/:id — Update User

```bash
curl -X PUT http://localhost:3000/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice Updated", "dob": "1991-03-15"}'
```

**Response `200`:**
```json
{
  "id": 1,
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

---

### DELETE /users/:id — Delete User

```bash
curl -X DELETE http://localhost:3000/users/1
```

**Response: `204 No Content`**

---

## ❌ Error Responses

| Status | Body                                |
|--------|-------------------------------------|
| 400    | `{"error": "invalid request"}`      |
| 404    | `{"error": "user not found"}`       |
| 500    | `{"error": "internal server error"}`|

---

## 🧪 Unit Tests

```bash
go test ./internal/service/... -v -run TestCalculateAge
```

Tests cover:
- ✅ Birthday already passed this year
- ✅ Birthday is exactly today
- ✅ Birthday not yet reached this year

---

## 🔒 Validation Rules

| Field | Rules                        |
|-------|------------------------------|
| name  | required                     |
| dob   | required, format: YYYY-MM-DD |

Invalid requests return `400 {"error": "invalid request"}`.

---

## 📝 Middleware

| Middleware        | Description                                          |
|-------------------|------------------------------------------------------|
| `X-Request-ID`    | UUID injected in every request & response header     |
| Request Duration  | Logs `GET /users/1 completed in 4ms` via Uber Zap   |

---

## ✅ Submission Checklist

- [x] GoFiber
- [x] PostgreSQL
- [x] SQLC
- [x] Uber Zap
- [x] Validator
- [x] CRUD APIs
- [x] Dynamic Age Calculation
- [x] Error Handling (400/404/500)
- [x] Logging (create/update/delete/errors)
- [x] Docker + Docker Compose
- [x] Pagination (`?page=1&limit=10`)
- [x] Request ID Middleware
- [x] Request Duration Middleware
- [x] Unit Test (CalculateAge)
- [x] README
