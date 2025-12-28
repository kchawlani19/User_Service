# 🧩 User Service – Go Backend (Clean Architecture)

A backend **User Management Service** built using **pure Go (no frameworks)**, following **clean layered architecture** and **configuration-driven design**.  
The project demonstrates how to build **production-ready REST APIs** with interchangeable storage layers.

---

## 🚀 Features

- User CRUD APIs (Create, Read, Update, Delete)
- Clean layered architecture:
  - Handler
  - Service
  - Repository
- Repository interface with multiple implementations:
  - In-Memory Store (for development/testing)
  - MySQL Store (for production-like setup)
- Configuration-based store switching (no code changes)
- Concurrency-safe in-memory store using mutex
- No frameworks – only Go standard library

---

## 🏗️ Architecture Overview

```
Client
  |
  v
HTTP Handler
  |
  v
Service (Business Logic)
  |
  v
Repository Interface
  |
  v
-------------------------
|                       |
In-Memory Store     MySQL Store
```

---

## 📂 Project Structure

```
go-basic-user-service/
│
├── main.go
│
├── model/
│   └── user.go
│
├── handler/
│   └── user_handler.go
│
├── service/
│   └── user_service.go
│
├── repository/
│   ├── user_repository.go
│   ├── inmemory_user_repository.go
│   └── db_user_repository.go
│
├── database/
│   └── db.go
│
├── go.mod
├── go.sum
```

---

## ⚙️ Configuration (Store Switching)

The application selects the storage implementation **at runtime** using an environment variable.

### Default (In-Memory Store)

```bash
go run main.go
```

### MySQL Store

```powershell
$env:STORE_TYPE="db"
go run main.go
```

No code changes are required to switch between stores.

---

## 🗄️ Database Schema (MySQL)

```sql
CREATE DATABASE userdb;
USE userdb;

CREATE TABLE users (
    id INT PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
```

---

## 🔌 API Endpoints

### Create User
```
POST /users
```

### Get User by ID
```
GET /users/{id}
```

### Update User
```
PUT /users/{id}
```

### Delete User
```
DELETE /users/{id}
```

---

## 🧠 Design Decisions

- **Handler Layer**
  - Handles HTTP concerns only
- **Service Layer**
  - Contains business rules and validations
- **Repository Layer**
  - Responsible only for data access
- **Main Function**
  - Performs dependency wiring and configuration-based decisions

---

## 🎯 Why This Project?

This project demonstrates:

- Strong backend fundamentals
- Clean separation of concerns
- Interface-driven design
- Configuration-based behavior switching
- Production-style Go backend development

---

## 🧑‍💻 Author

**Khushi Chawlani**  

---

## 📌 Future Enhancements

- Add more fields to user model
- Pagination support
- Structured logging
- Authentication & authorization
