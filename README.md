# Go Order API

A RESTful API built with Go (Golang) for managing users, products, and orders.  
This project demonstrates clean architecture, JWT authentication, and proper separation of concerns, suitable for portfolio and real-world backend services.

---

## 🚀 Features

- User authentication (JWT)
- Create & list orders
- Product validation
- Clean architecture (controller → service → repository)
- MySQL with GORM auto-migration
- Environment-based configuration

---

## 🧰 Tech Stack

- **Language**: Go
- **Router**: Chi
- **ORM**: GORM
- **Database**: MySQL
- **Auth**: JWT
- **Config**: godotenv (.env)

---

## 📁 Project Structure

```
cmd/
 └─ api/              # Application entry point
internal/
 ├─ config/           # Database initialization
 ├─ controller/       # HTTP controllers
 ├─ middleware/       # Auth & context middleware
 ├─ models/           # GORM models
 ├─ pkg/auth/         # JWT manager
 ├─ repositories/     # Data access layer
 ├─ routes/           # Route registration
 └─ services/         # Business logic
```

---

## 🛠 Requirements

- Go 1.20+
- MySQL

---

## ⚙️ Setup & Run

### 1. Clone repository
```bash
git clone https://github.com/fahmiabd/go-order-api.git
cd go-order-api
```

### 2. Create `.env`
```env
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASS=password
DB_NAME=go_order_api

JWT_SECRET=supersecretkey
```

### 3. Install dependencies
```bash
go mod tidy
```

### 4. Run application
```bash
go run ./cmd/api
```

Server runs on:
```
http://localhost:8080
```

---

## 🔐 API Endpoints

### Auth
| Method | Endpoint | Description |
|------|----------|------------|
| POST | `/login` | Login & get JWT |

### Orders (Protected)
Header:
```
Authorization: Bearer <token>
```

| Method | Endpoint | Description |
|------|----------|------------|
| POST | `/orders` | Create order |
| GET | `/orders` | List user orders |

---

## 🗄 Auto Migration

On startup, the application automatically migrates:
- User
- Product
- Order

---

## 🧠 Notes

- Passwords are stored hashed
- JWT tokens have expiration
- `.env` file should never be committed

---

## 📄 License

This project is open-source and free to use as learning material or backend template.
