# 📘 Month 3 – Testing, Web, CLIs & Production Readiness

Month 3 focuses on **building production-ready applications** in Go:  
you’ll design real-world CLIs, REST APIs, add middleware, write tests & benchmarks, and deploy your final capstone project.

---

## 📅 Weekly Breakdown

### Week 9: Web & CLI Basics
- CLI tools: Cobra, urfave/cli, bubbletea
- Build CLI with flags, subcommands
- Web basics with `net/http` (handlers, ServeMux, routers)
- Middleware introduction (logging, request handling)
- Testing HTTP code with `httptest`
- **Project:** CLI tool that calls a REST API

### Week 10: REST APIs & Middleware
- REST API design with JSON in/out
- Router libraries (chi, gorilla/mux)
- Middleware: logging, error handling, auth stub
- Testing REST APIs with `httptest` (unit + integration)
- **Project:** Task Manager REST API

### Week 11: Capstone Build
- Build combined service + CLI project (e.g., Task Tracker)
- Use `go:embed` for assets/templates
- Add background jobs with goroutines
- Write tests, table-driven tests, and benchmarks
- Package with modules, ensure coverage
- Dockerize app (multi-stage build)

### Week 12: Deployment & Polish
- Finalize project with full test coverage
- Docker Compose for local deployment
- Optional cloud deployment (AWS/GCP)
- Add profiling hooks (`pprof`) for performance insights
- Polish documentation: README, architecture diagram, usage notes

---

## 🎯 Goals for Month 3
- Comfortably build CLIs and REST APIs in Go
- Write reusable middleware and test HTTP handlers
- Use `go:embed` for embedding assets and configs
- Incorporate concurrency in real-world apps (background jobs)
- Achieve test coverage and benchmarking confidence
- Containerize and deploy Go apps in a production-style setup

---

## 📝 Deliverables
- Weekly READMEs under `month-3/week-*`
- Projects for CLI, REST API, and Capstone
- Notes under `notes/week-*`
- Polished final README and documentation
- Optional deployment notes for Docker/Cloud

---

## 📚 References
- Cobra: https://github.com/spf13/cobra
- urfave/cli: https://github.com/urfave/cli
- Bubbletea: https://github.com/charmbracelet/bubbletea
- net/http: https://pkg.go.dev/net/http
- chi router: https://github.com/go-chi/chi
- httptest: https://pkg.go.dev/net/http/httptest
- Dockerizing Go: https://docs.docker.com/language/golang/
- pprof: https://pkg.go.dev/net/http/pprof

---

✅ End of Month 3: You’ll have built, tested, Dockerized, and optionally deployed a **production-grade Go service + CLI**, completing the 3‑month roadmap.  
🎉 Congratulations — you’ve reached Go Pro status!
