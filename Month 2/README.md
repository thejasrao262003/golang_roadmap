# 📘 Month 2 – Concurrency & Standard Library

Month 2 dives into **Go’s concurrency model and essential standard library packages**.  
By the end of this month, you’ll be confident structuring Go projects with modules, using key stdlib tools, and writing safe concurrent programs.

---

## 📅 Weekly Breakdown

### Week 5: Modules & Standard Library
- Learn Go modules (`go.mod`, `go.sum`), dependency management, vendoring
- Standard library deep dive:
  - `encoding/json`, `time`, `os`, `regexp`, `embed`
- **Project:**
  - File watcher or parser using stdlib packages

### Week 6: Concurrency (Goroutines & Channels)
- Goroutines and lightweight concurrency
- Channels: buffered vs unbuffered
- Select statement for multiplexing
- Race detection with `-race`
- **Projects:**
  - Concurrent counter
  - Channel pipeline
  - Worker pool basics

### Week 7: Concurrency Patterns (Sync, Context, Pipelines)
- `sync.WaitGroup` and `sync.Mutex` for coordination & safety
- `context` package for cancellation and deadlines
- Build pipelines with goroutines + channels
- **Projects:**
  - Concurrent web scraper (basic version)
  - Rate limiter (simple token bucket)
  - Multi-stage channel pipeline

### Week 8: Concurrency Capstone Project
- Choose a **larger project** to consolidate skills:
  - Concurrent web scraper with worker pool, rate limiting, retries, context cancellation
  - Token bucket rate limiter (library + demo)
- Emphasis on:
  - Production patterns (logging, graceful shutdown)
  - Testing (`-race`, table-driven, benchmarks)
  - Profiling (`pprof`)

---

## 🎯 Goals for Month 2
- Comfortably use Go modules & manage dependencies
- Leverage standard library effectively for real utilities
- Understand goroutines, channels, and select deeply
- Apply sync primitives and context cancellation
- Write production-style concurrent code with tests and profiling

---

## 📝 Deliverables
- Weekly READMEs in `month-2/week-*`
- Projects under `week-5`, `week-6`, `week-7`, `week-8`
- Notes under `notes/week-*`
- Updated progress checklist in [root README](../README.md)

---

## 📚 References
- Go Modules: https://go.dev/doc/modules
- Go Blog – Concurrency Patterns: https://go.dev/blog/pipelines
- Go by Example (channels, goroutines, sync, context): https://gobyexample.com/
- Effective Go – Concurrency: https://go.dev/doc/effective_go#concurrency

---

✅ End of Month 2: You’ll be confident designing **safe, concurrent Go programs** and structuring real-world projects with modules and the standard library.
