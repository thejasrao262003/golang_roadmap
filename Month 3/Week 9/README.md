# 📅 Week 9 – Web & CLI Basics

Month 3 begins with **real-world interfaces**: building CLIs and basic web servers.  
You’ll practice with CLI frameworks and Go’s `net/http` package, and write tests for web handlers.

---

## 🎯 Learning Objectives
- [ ] Learn CLI frameworks (Cobra, urfave/cli, bubbletea)
- [ ] Build a CLI tool with flags, arguments, and subcommands
- [ ] Create a web service using Go’s `net/http`
- [ ] Write custom handlers, use routers and middleware
- [ ] Test HTTP code with `httptest`

---

## 📖 Core Topics

### CLI Tools
- Cobra basics:
  - Root command + subcommands
  - Flags and arguments (`--name`, `--port`)
  - Auto-generated help
- Alternatives: urfave/cli, bubbletea (for TUI apps)

### Web Services
- `net/http` essentials:
  - `http.HandleFunc`, `http.ListenAndServe`
  - Custom `Handler` interface implementation
- Routers:
  - Use stdlib `http.ServeMux` or third-party like [chi](https://github.com/go-chi/chi)
- Middleware:
  - Logging, auth, request timing

### Testing
- Use `httptest.NewRecorder()` and `httptest.NewRequest()`
- Verify response codes, headers, and body

---

## 💻 Project: CLI + REST API Demo

**Goal:** Build a CLI tool that fetches data from a public REST API.

### Features
- CLI built with Cobra:
  - `fetch` subcommand with flags (`--url`, `--output`)
- Make HTTP GET requests to API endpoint
- Parse JSON response (`encoding/json`)
- Display results in terminal or save to file

### Example Usage
```bash
go run cmd/apicli/main.go fetch --url "https://api.github.com/users/octocat" --output user.json
```

### Suggested layout
```
week-9/apicli/
├── cmd/apicli/main.go      # entry point for CLI
├── internal/fetcher/       # HTTP fetch + JSON decode
├── internal/cli/           # Cobra setup
└── go.mod
```

---

## 📝 Deliverables
- CLI project in `week-9/apicli/`
- Notes in `notes/week-9.md` covering:
  - CLI vs web server patterns
  - How httptest works for handlers
  - Example test cases written
- Update checklist in [root README](../README.md)

---

## 📚 References
- Cobra: https://github.com/spf13/cobra
- urfave/cli: https://github.com/urfave/cli
- Bubbletea (TUI): https://github.com/charmbracelet/bubbletea
- Go net/http: https://pkg.go.dev/net/http
- Go httptest: https://pkg.go.dev/net/http/httptest

---

✅ End of Week 9: You should be able to build both CLI tools and basic web services in Go, and write tests for HTTP handlers.
