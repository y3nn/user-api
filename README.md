# 🌟 yoyo — Clean & Minimal Go Architecture Example

**yoyo** is a lightweight, framework-free example that demonstrates how to structure Go applications in a clean, understandable, and maintainable way.  
It serves as a practical reference for developers who want to build services with clarity, simplicity, and architectural discipline — without unnecessary abstractions.

---

## ✨ Key Ideas

- 🧩 **Clarity first** — simple, readable project layout  
- 📦 **Logical package boundaries** — clear separation of responsibilities  
- ⚡ **Minimalism** — no frameworks, no magic, only Go  
- 🧠 **Easy to understand** — suitable as a teaching/reference example  
- 🔧 **Extendable** — structure scales naturally as the project grows  

---

## 📁 Project Structure

```txt
yoyo/
├── cmd/            # Application entry points
│   └── app/        # Main service executable
├── internal/       # Core domain & business logic
│   ├── service/    # Application services
│   ├── repo/       # Storage & repositories
│   └── model/      # Entities & domain models
├── pkg/            # Shared utilities (optional)
└── go.mod
