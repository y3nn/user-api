# simple  tiny users-api

This project is a minimal 👤 user API written in Go 🦦.  
It uses PostgreSQL 🐘 for storing user data and is easy to run with Docker 🐳.

---

## ✨ Features

- ⚡ User CRUD API in Go (no frameworks)
- 🐘 PostgreSQL database
- 🐳 Easy Docker setup

---

## 🗂️ Project Structure

```txt
user-api/
├── cmd/            # 🚀 Application entry point
├── internal/       
│   ├── repository/ # 🗄️ Database interaction (Postgres)
│   └── model/      # 👤 User model
├── go.mod
├── Dockerfile
└── docker-compose.yaml
```

---

## 🚦 Getting Started

1. 📝 **Copy** `.env.example` **to** `.env` **and set your database values.**

2. 🐳 **Build & start everything using Docker Compose:**
   ```bash
   docker-compose up --build
   ```
   This will run the API (on port 8080) and the Postgres DB.

3. 📡 **API endpoints:**
   - `POST   /users` — ➕ create user
   - `GET    /users/{id}` — 🔍 get user by id
   - `GET    /users` — 📃 list all users

---

That's it 🎉 — clean, simple Go user API with Postgres and Docker!
