# 📅 Week 8 – Concurrency Project (Capstone of Month 2)

This week you’ll **ship a production‑style concurrent utility** that consolidates goroutines, channels, sync primitives, context cancellation, and testing.  
Pick **one** primary project (or do both if time permits).

---

## 🚀 Project Options

### Option A: Concurrent Web Scraper (Robust)
**Goal:** Fetch many pages concurrently, extract data, and write a single report.

**Requirements**
- Concurrency:
  - Worker pool (configurable size via flag or env)
  - Job / result channels, `sync.WaitGroup`
- Context:
  - `--timeout` to cancel globally (e.g., 10s)
  - Per‑request timeout using `context.WithTimeout`
- Networking:
  - HTTP client with timeouts; retries with backoff
- Robustness:
  - Rate limiting (token bucket via `time.Ticker` or buffered channel)
  - Deduplicate URLs (concurrent‑safe set/map)
- Output:
  - Write JSON/CSV summary (status code, content length, title, error)
- Observability:
  - Structured logs (JSON lines) and metrics (counters, timings)

**Suggested layout**
```
week-8/scraper/
├── cmd/scraper/main.go
├── internal/httpx/        # HTTP client, retry/backoff
├── internal/limiter/      # rate limiter
├── internal/scrape/       # worker pool, parse/extract
├── internal/report/       # JSON/CSV writers
└── go.mod
```

---

### Option B: Token‑Bucket Rate Limiter (Library + Demo)
**Goal:** Implement a reusable **rate limiter** library with examples.

**Requirements**
- Token bucket with refill (`time.Ticker`), `Allow()`/`Wait()` APIs
- Context‑aware `Wait(ctx)` that cancels cleanly
- Benchmarks for different burst/throughput settings
- Example programs:
  - CLI that limits API calls from a file of requests
  - HTTP middleware that limits requests per client IP

**Suggested layout**
```
week-8/ratelimit/
├── cmd/demo/              # CLI demo
├── cmd/httpdemo/          # HTTP middleware demo
├── limiter/               # library (public API)
├── limiter_test.go        # unit & race tests
└── go.mod
```

---

## 🧪 Testing & Quality Gates (for either project)
- Unit tests (table‑driven) for core components
- Concurrency sanity:
  - `go test -race ./...`
  - Use `-run TestName -count=50` for flake hunt
- Benchmarks:
  - `go test -bench=. -benchmem ./...`
- Profiling (optional but recommended):
  - `go test -cpuprofile cpu.out -memprofile mem.out`
  - Integrate `net/http/pprof` in a debug build tag
- Lint & format:
  - `gofmt -s -w .` and `go vet ./...`
  - `golangci-lint run` (if configured)

---

## 🔧 Implementation Hints
- Use `select` with:
  - `<-ctx.Done()` for cancellation
  - `<-ticker.C` for pacing
  - channel sends/receives for pipeline steps
- Backoff template:
  ```go
  for i := 0; i < maxRetries; i++ {
      err := do()
      if err == nil { break }
      time.Sleep(time.Duration(1<<i) * time.Millisecond) // exp backoff
  }
  ```
- Safe maps: guard with `sync.Mutex`/`sync.RWMutex`, or use a worker that owns the map

---

## ✅ Acceptance Checklist
- [ ] Builds with `go build ./...`
- [ ] `go test ./...` passes (with `-race`)
- [ ] Benchmarks included and documented
- [ ] README with run examples and flags
- [ ] Graceful shutdown via context
- [ ] Logs are structured and helpful

---

## 📝 Deliverables
- Code under `week-8/<project>/`
- `README.md` with usage examples, flags, and architecture diagram (ASCII OK)
- Notes in `notes/week-8.md` summarizing:
  - Concurrency model (workers, channels, select)
  - Cancellation strategy and rate limiting
  - Testing/bench/profile results & learnings

---

## 📚 Reference
- Go Blog – Pipelines: https://go.dev/blog/pipelines
- Context: https://pkg.go.dev/context
- http.Client timeouts: https://pkg.go.dev/net/http#Client
- pprof: https://pkg.go.dev/net/http/pprof

---

🎯 End of Week 8: You’ve shipped a **concurrent, tested, cancelable** Go utility/library and gathered performance insights.
