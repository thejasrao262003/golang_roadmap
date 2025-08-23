# 📅 Week 6 – Concurrency (Goroutines & Channels)

This week is focused on **Go’s concurrency model** – one of the language’s biggest strengths.  
You’ll learn to use goroutines, channels, and the `select` statement to build concurrent programs.

---

## 🎯 Learning Objectives
- [ ] Understand what **goroutines** are and how to launch them
- [ ] Learn how to use **channels** for communication between goroutines
- [ ] Differentiate between **buffered vs unbuffered channels**
- [ ] Use the **`select` statement** for multiplexing channel operations
- [ ] Practice race detection with `go run -race`

---

## 📖 Core Concurrency Topics

### Goroutines
- Launch with `go f()`, run independently of main thread
- Lightweight compared to OS threads

### Channels
- Create with `make(chan T)`
- Send (`ch <- val`) and receive (`val := <-ch`)
- Close channels (`close(ch)`) and check with `v, ok := <-ch`

### Buffered vs Unbuffered Channels
- Unbuffered: synchronous send/receive
- Buffered: allows limited async behavior

### Select Statement
- Multiplex over multiple channels
- Default case for non-blocking operations

### Race Detection
- Use `go run -race main.go` to detect race conditions

---

## 💻 Projects (Hands-On)

1. **Concurrent Counter**
   - Launch multiple goroutines incrementing a shared counter
   - Demonstrate race conditions
   - Fix with a channel-based solution

2. **Channel Pipeline**
   - Create a pipeline of goroutines:
     - One generates numbers
     - Another squares them
     - Another prints results
   - Communicate only via channels

3. **Worker Pool**
   - Implement a worker pool pattern using goroutines + channels
   - Workers fetch “jobs” from a channel, process them, and send results back

4. **Timeout Example**
   - Use `select` with `time.After` to add a timeout to a goroutine operation

---

## 📝 Deliverables
- Code in `week-6/` demonstrating goroutines, channels, select
- Notes in `notes/week-6.md` covering:
  - Difference between goroutines & threads
  - Buffered vs unbuffered behavior observed
  - Example pipeline/worker pool code
- Update checklist in [root README](../README.md)

---

## 📚 Reference
- [Go by Example – Goroutines](https://gobyexample.com/goroutines)
- [Go by Example – Channels](https://gobyexample.com/channels)
- [Go by Example – Channel Buffering](https://gobyexample.com/channel-buffering)
- [Go by Example – Select](https://gobyexample.com/select)
- [Go Blog – Go Concurrency Patterns](https://go.dev/blog/pipelines)

---

✅ End of Week 6: You should be comfortable launching goroutines, using channels, and building simple concurrent patterns in Go.
