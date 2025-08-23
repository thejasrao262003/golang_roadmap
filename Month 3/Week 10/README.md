# 📅 Week 10 – REST APIs & Middleware

This week focuses on building **REST APIs in Go** with routing, middleware, and testing.  
You’ll design clean HTTP handlers, implement middleware, and write integration tests.

---

## 🎯 Learning Objectives
- [ ] Build REST APIs using `net/http` and router libraries (chi or gorilla/mux)
- [ ] Organize code with `cmd/`, `internal/`, and `pkg/` structure
- [ ] Write reusable middleware (logging, auth, error handling)
- [ ] Test APIs with `httptest` (unit + integration style)
- [ ] Understand JSON request/response handling in REST endpoints

---

## 📖 Core Topics

### REST APIs
- CRUD patterns: GET, POST, PUT, DELETE
- JSON request parsing (`json.NewDecoder`)
- JSON response writing (`json.NewEncoder`)

### Routers
- Stdlib `http.ServeMux`
- Third-party: [chi](https://github.com/go-chi/chi), [gorilla/mux](https://github.com/gorilla/mux)

### Middleware
- Logging middleware (request/response time)
- Authentication placeholder (check header)
- Error handler for consistent JSON errors

### Testing
- `httptest.NewRecorder()` and `httptest.NewRequest()`
- Table-driven tests for multiple HTTP methods
- Integration-style test with in-memory server (`httptest.NewServer`)

---

## 💻 Project: REST API Service

**Goal:** Implement a basic REST API service for managing “Tasks”.

### Features
- Endpoints:
  - `GET /tasks` → list tasks
  - `POST /tasks` → create task
  - `GET /tasks/{id}` → get by ID
  - `PUT /tasks/{id}` → update task
  - `DELETE /tasks/{id}` → delete task
- Middleware:
  - Logging requests
  - Error handling
- Testing:
  - Unit tests for handlers
  - Integration tests with `httptest.NewServer`

### Suggested layout
```
week-10/taskapi/
├── cmd/taskapi/main.go     # server entry point
├── internal/handlers/      # REST handlers
├── internal/models/        # task struct, in-memory store
├── internal/middleware/    # logging, error handling
└── go.mod
```

---

## 📝 Deliverables
- REST API project in `week-10/taskapi/`
- Tests for all endpoints
- Notes in `notes/week-10.md` covering:
  - Routing patterns
  - Middleware use cases
  - Testing strategies used
- Update checklist in [root README](../README.md)

---

## 📚 References
- net/http: https://pkg.go.dev/net/http
- chi router: https://github.com/go-chi/chi
- httptest: https://pkg.go.dev/net/http/httptest
- JSON encoding/decoding: https://pkg.go.dev/encoding/json

---

✅ End of Week 10: You should be able to build, structure, and test a REST API in Go with middleware and proper JSON handling.
