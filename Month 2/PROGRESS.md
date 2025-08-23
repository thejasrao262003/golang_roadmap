# 📊 Progress Tracker – Month 2 (Concurrency & Standard Library)

This document tracks my **progress for Month 2** of the Go learning roadmap.  
I will update weekly with notes, completed projects, and learnings.

---

## ✅ Overview
Month 2 dives into **Go’s concurrency model and essential standard library packages**.  
By the end of this month, I should be confident structuring Go projects with modules, using key stdlib tools, and writing safe concurrent programs.

---

## 📅 Weekly Progress

### Week 5 – Modules & Standard Library
**Goals:**
- Learn Go modules (`go.mod`, `go.sum`), vendoring
- Practice stdlib packages: `encoding/json`, `time`, `os`, `regexp`, `embed`

**Project:**
- [ ] File watcher / parser using stdlib packages

**Notes:**
- Key learnings:
- Challenges:
- Links to code:


### Week 6 – Concurrency (Goroutines & Channels)
**Goals:**
- Launch goroutines
- Communicate with channels (buffered vs unbuffered)
- Use `select` for multiplexing
- Run with race detector (`-race`)

**Projects:**
- [ ] Concurrent counter
- [ ] Channel pipeline
- [ ] Worker pool basics

**Notes:**
- Key learnings:
- Challenges:
- Links to code:


### Week 7 – Concurrency Patterns (Sync, Context, Pipelines)
**Goals:**
- Coordinate tasks with `sync.WaitGroup` and `sync.Mutex`
- Manage cancellations and deadlines with `context`
- Chain goroutines into pipelines

**Projects:**
- [ ] Concurrent web scraper (basic)
- [ ] Rate limiter (token bucket)
- [ ] Multi-stage channel pipeline

**Notes:**
- Key learnings:
- Challenges:
- Links to code:


### Week 8 – Concurrency Capstone Project
**Goals:**
- Build a larger project consolidating all concurrency tools
- Add production concerns: logging, graceful shutdown, retries
- Test with `-race`, add benchmarks, profiling (`pprof`)

**Projects:**
- [ ] Concurrent web scraper (worker pool, rate limiting, retries, context)
- [ ] OR: Token bucket rate limiter (library + demo)

**Notes:**
- Key learnings:
- Challenges:
- Links to code:


---

## 🎯 Month 2 Goals Recap
- [ ] Comfortably use Go modules & manage dependencies
- [ ] Leverage standard library effectively
- [ ] Use goroutines, channels, and select confidently
- [ ] Apply sync primitives and context cancellation
- [ ] Write production-style concurrent code with tests, benchmarks, profiling

---

✅ End of Month 2: I should be able to design **safe, concurrent Go programs** and structure projects with modules + stdlib tools.
