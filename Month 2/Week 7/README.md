# 📅 Week 7 – Concurrency Patterns (Sync, Context & Pipelining)

This week is about mastering **concurrency control tools and patterns** in Go.  
You’ll learn how to coordinate goroutines with `sync`, manage cancellations with `context`, and apply concurrency to real-world problems.

---

## 🎯 Learning Objectives
- [ ] Deepen understanding of **goroutines, channels, select**
- [ ] Learn **sync tools**: `WaitGroup`, `Mutex`
- [ ] Use **context** for cancellation and deadlines
- [ ] Build **pipelines** with goroutines & channels
- [ ] Practice race detection with `go run -race`

---

## 📖 Core Concurrency Topics

### 1. Sync Tools
- `sync.WaitGroup` to wait for multiple goroutines
- `sync.Mutex` to protect shared state

### 2. Context
- Create with `context.Background()`
- Derive with `context.WithCancel`, `context.WithTimeout`, `context.WithDeadline`
- Pass context through goroutines to stop work gracefully

### 3. Pipelining
- Chain goroutines together, each processing and passing results to the next stage via channels
- Use select + context to gracefully terminate pipelines

### 4. Race Detection
- Compile/run with `-race` to catch data races early

---

## 💻 Projects (Hands-On)

1. **Concurrent Web Scraper**
   - Input: list of URLs
   - Launch goroutines to fetch each URL concurrently
   - Use `sync.WaitGroup` to wait for completion
   - Protect shared slice/map with `sync.Mutex`
   - Add context cancellation if runtime exceeds 5s

2. **Rate Limiter**
   - Implement a rate limiter using channels and `time.Ticker`
   - Allow only N requests per second

3. **Pipeline Example**
   - Build a 3-stage pipeline:
     - Generator → Processor → Consumer
   - Each stage in its own goroutine, connected by channels
   - Use context to cancel pipeline gracefully

---

## 📝 Deliverables
- Code in `week-7/` for scraper, rate limiter, and pipeline
- Notes in `notes/week-7.md` covering:
  - Usage of WaitGroup vs Mutex
  - How context simplified goroutine cancellation
  - Example pipeline flow
- Update checklist in [root README](../README.md)

---

## 📚 Reference
- [Go by Example – WaitGroups](https://gobyexample.com/waitgroups)
- [Go by Example – Mutexes](https://gobyexample.com/mutexes)
- [Go Blog – Context Package](https://go.dev/blog/context)
- [Go Blog – Concurrency Pipelines](https://go.dev/blog/pipelines)

---

✅ End of Week 7: You should be confident coordinating concurrent tasks using WaitGroups, Mutexes, Context, and building pipelines.
