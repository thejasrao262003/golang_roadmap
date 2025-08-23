# 📅 Week 11 – Capstone Build (Service + CLI)

This week you’ll **build the capstone**: a small production-style web service with an accompanying CLI.  
You’ll use `go:embed`, JSON (marshal/unmarshal), robust error handling, **concurrency** for background jobs, **tests + benchmarks**, and package everything with modules and a Docker image.

---

## 🎯 Objectives
- [ ] Implement a **Task Tracker** service (HTTP) with a **CLI** client
- [ ] Use `go:embed` for seeding templates or sample data
- [ ] Marshal/unmarshal JSON for API contracts and storage
- [ ] Add a background job using **goroutines** (e.g., periodic cleanup or scheduler)
- [ ] Write unit tests, table-driven tests, and **benchmarks**
- [ ] Ensure module tidiness + coverage (`go mod tidy`, `go test -cover`)
- [ ] Dockerize the app; run locally, optionally deploy to cloud
- [ ] Ship a polished README with usage examples

---

## 🏗 Suggested Architecture

```
week-11/capstone/
├── cmd/
│   ├── server/            # HTTP server entrypoint
│   │   └── main.go
│   └── cli/               # CLI client entrypoint (Cobra or urfave/cli)
│       └── main.go
├── internal/
│   ├── api/               # handlers, router, request/response models
│   ├── tasks/             # domain: Task repo, service, validation
│   ├── jobs/              # background jobs (goroutines + context)
│   ├── storage/           # JSON file storage (or in-memory)
│   ├── middleware/        # logging, request ID, recover
│   └── version/           # build-time version info
├── assets/                # embedded seed data / templates
├── Dockerfile
├── docker-compose.yml     # optional (for local run)
└── go.mod
```

---

## 🔌 Features

### Web Service (HTTP)
- Endpoints (JSON in/out):
  - `GET /tasks` → list tasks
  - `POST /tasks` → create task
  - `GET /tasks/{id}` → fetch one
  - `PUT /tasks/{id}` → update
  - `DELETE /tasks/{id}` → delete
- Middleware:
  - request logging, recovery, request id header
- Error handling:
  - consistent JSON error envelope: `{ "error": "...", "code": ... }`

### CLI Client
- Commands:
  - `tasks list`
  - `tasks add --title "..." --due "2025-08-31"`
  - `tasks done --id <id>`
  - `tasks rm --id <id>`
- Targets the HTTP service; supports `--server http://localhost:8080`

### Concurrency – Background Job
- Goroutine that runs on a `time.Ticker` (e.g., every minute)
- Examples:
  - auto-archive completed tasks older than N days
  - send periodic summary to logs
- Controlled via `context.Context` and graceful shutdown on SIGINT/SIGTERM

### go:embed
- Embed initial `assets/seed.json` and optional HTML templates
- Use embedded seed data to pre-populate storage (one-time init)

---

## 🧪 Testing & Benchmarks

- **Unit tests** for handlers, services, and storage
- **Table-driven tests** for CRUD validation and error paths
- **Integration test** with `httptest.NewServer`
- **Benchmarks** for hot code paths (e.g., JSON parse, in-memory filtering)

Commands:
```bash
go test ./... -v -race
go test ./... -cover
go test ./... -bench=. -benchmem
```

---

## 🧰 Example Code Snippets

**Embed seed data**
```go
//go:embed assets/seed.json
var seedFS embed.FS
```

**Graceful shutdown**
```go
ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
defer stop()
srv := &http.Server{Addr: ":8080", Handler: router}
go func() {
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
        log.Fatal(err)
    }
}()
<-ctx.Done()
shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
_ = srv.Shutdown(shutdownCtx)
```

**Background job**
```go
func (j *Janitor) Run(ctx context.Context, d time.Duration) {
    ticker := time.NewTicker(d)
    defer ticker.Stop()
    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            j.Clean()
        }
    }
}
```

---

## 🐳 Dockerization

**Dockerfile**
```dockerfile
# Build stage
FROM golang:1.22 AS build
WORKDIR /app
COPY . .
RUN go mod download && go build -o server ./cmd/server && go build -o cli ./cmd/cli

# Run stage
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /app/server /usr/local/bin/server
COPY --from=build /app/cli /usr/local/bin/cli
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/server"]
```

**docker-compose.yml (optional)**
```yaml
version: "3.9"
services:
  capstone:
    build: .
    ports:
      - "8080:8080"
    environment:
      - LOG_LEVEL=info
```

Run:
```bash
# Build & run in Docker
docker build -t capstone .
docker run -p 8080:8080 capstone

# Or with compose
docker compose up --build
```

---

## ✅ Acceptance Checklist
- [ ] HTTP service implements CRUD with JSON contracts
- [ ] CLI works against the service (list/add/done/rm)
- [ ] Background job runs, is cancelable, and logs actions
- [ ] `go:embed` used for seed/templates
- [ ] Tests pass with `-race` and coverage reported
- [ ] Benchmarks included and documented
- [ ] Docker image builds and runs locally
- [ ] README documents commands, flags, and endpoints

---

## 📝 Deliverables
- Code in `week-11/capstone/`
- `README.md` with:
  - run commands (local & Docker)
  - CLI usage examples
  - API endpoints (with sample requests)
  - notes on design and tradeoffs
- Notes in `notes/week-11.md` summarizing:
  - design choices (storage, routing, concurrency)
  - test/bench/profile insights
  - deployment notes

---

## 📚 References
- embed: https://pkg.go.dev/embed
- net/http: https://pkg.go.dev/net/http
- httptest: https://pkg.go.dev/net/http/httptest
- Cobra: https://github.com/spf13/cobra
- Distroless images: https://github.com/GoogleContainerTools/distroless

---

🎯 End of Week 11: A **feature-complete, tested, Dockerized** service + CLI, ready for final polish and deployment next week.
