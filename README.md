# 🚀 Go Learning Roadmap (3 Months)

Welcome! 👋  
This repository documents my **journey to becoming a Go (Golang) pro in 3 months**.  

I come from a Python & C++ background (basic syntax, problem solving), and this repo is where I’ll track my progress, notes, and projects as I master Go for backend, concurrency, and production-ready systems.

---

## 📌 Goals
- Learn idiomatic Go from basics → concurrency → production.
- Build weekly projects to reinforce concepts.
- Document all progress in a structured way (commits, notes, code).
- At the end: deploy a full **Go REST API service** with Docker & tests.

---

## 🛠 Tech Stack
- **Language**: Go 1.22+
- **Tools**: `go run`, `go build`, `go test`, `go mod`, `gofmt`
- **Frameworks & Libraries**:
  - Web: `net/http`, [chi](https://github.com/go-chi/chi)
  - CLI: [Cobra](https://github.com/spf13/cobra)
  - DB: `database/sql`, [gorm](https://gorm.io/)
  - Config: [viper](https://github.com/spf13/viper)
- **Deployment**: Docker, systemd / Kubernetes (optional)

---

## 🗂 Repo Structure

go-roadmap/
├── month-1/
│ ├── week-1/ # Basics, syntax, setup
│ ├── week-2/ # Structs, interfaces, errors
│ ├── week-3/ # Testing, generics
│ └── week-4/ # Mini projects
├── month-2/
│ ├── week-5/ # Modules, std lib
│ ├── week-6/ # Concurrency (goroutines, channels)
│ ├── week-7/ # Sync, context, worker pools
│ └── week-8/ # Concurrency project
├── month-3/
│ ├── week-9/ # Web & CLI basics
│ ├── week-10/ # REST APIs, testing
│ ├── week-11/ # Capstone build
│ └── week-12/ # Deployment & polish
├── notes/ # Daily/weekly notes & learnings
├── projects/ # Independent larger projects
└── README.md
---

## 📅 Timeline

### Month 1 – Core Go Foundations
- [ ] Week 1: Setup, syntax, variables, control flow
- [ ] Week 2: Structs, methods, interfaces, error handling
- [ ] Week 3: Generics, table-driven tests
- [ ] Week 4: Mini projects (CLI calculator, JSON parser, file I/O)

### Month 2 – Concurrency & Standard Library
- [ ] Week 5: Go modules, std lib (os, time, encoding/json, regex)
- [ ] Week 6: Goroutines, channels, select
- [ ] Week 7: Sync, context, worker pools, cancellation
- [ ] Week 8: Concurrency project (scraper, rate limiter)

### Month 3 – Backend, CLI & Production
- [ ] Week 9: CLI tools, net/http
- [ ] Week 10: REST APIs, middleware, testing with httptest
- [ ] Week 11: Capstone project build
- [ ] Week 12: Dockerize + deploy

---

## 📚 Reference Resources
- 📖 [The Go Programming Language (Donovan & Kernighan)](https://www.gopl.io/)
- 🛠 [roadmap.sh/golang](https://roadmap.sh/golang)
- 🎥 Go by Example → https://gobyexample.com/
- 📦 Effective Go → https://go.dev/doc/effective_go
- 💡 Go Proverbs (Rob Pike)

---

## ✅ Progress Log
Each week I’ll commit:
- Code exercises
- Mini-projects
- Notes in `notes/`
- Checklist updates in this README

---

## 🏆 Final Capstone
At the end of 3 months → a **production-ready REST API**:
- User authentication
- CRUD operations
- SQL/NoSQL database
- Unit & integration tests
- Dockerized & deployed to cloud

---

✨ Follow along as I document my Go journey here.  