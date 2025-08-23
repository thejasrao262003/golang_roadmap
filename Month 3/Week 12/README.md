# 📅 Week 12 – Deployment & Polish (Final Week)

This week is about **finishing touches**: packaging, deploying, and polishing your Go capstone project.  
You’ll containerize the service, optionally deploy it, add CI/CD, and ensure documentation + test coverage are production-ready.

---

## 🎯 Objectives
- [ ] Finalize capstone project with clean code, docs, and tests
- [ ] Add full module support and run `go mod tidy`
- [ ] Ensure **comprehensive test coverage** (`go test -cover`)
- [ ] Dockerize service and CLI (multi-stage build)
- [ ] Deploy locally with Docker Compose or to cloud (AWS/GCP)
- [ ] Add profiling hooks (`pprof`) for performance
- [ ] Polish documentation (README, usage, design notes)

---

## 🏗 Deployment Steps

### 1. Dockerization
- Multi-stage Dockerfile (build + runtime)
- Build image:
  ```bash
  docker build -t capstone:latest .
  docker run -p 8080:8080 capstone:latest
  ```

### 2. Docker Compose (optional)
- Combine service + dependencies (e.g., Postgres, Redis)
- Example `docker-compose.yml`:
  ```yaml
  version: "3.9"
  services:
    capstone:
      build: .
      ports:
        - "8080:8080"
  ```

### 3. Cloud Deployment (stretch goal)
- AWS ECS (Fargate) or EKS (Kubernetes)
- GCP Cloud Run (serverless containers)
- Or run on bare EC2 with `systemd`

---

## 🧪 Testing & Benchmarking

- Run tests with race detection and coverage:
  ```bash
  go test ./... -v -race -cover
  ```
- Benchmark performance:
  ```bash
  go test ./... -bench=. -benchmem
  ```
- Add profiling routes (optional):
  ```go
  import _ "net/http/pprof"
  go http.ListenAndServe("localhost:6060", nil)
  ```

---

## 📄 Documentation Polish

- Update root `README.md` with:
  - Project overview
  - CLI usage examples
  - REST API endpoints (with sample curl requests)
  - Docker build & run instructions
  - Test & benchmark commands
- Add architecture diagram (ASCII or draw.io export)
- Write notes on design tradeoffs in `notes/week-12.md`

---

## ✅ Acceptance Checklist
- [ ] Tests: all green with `-race`, coverage > 80%
- [ ] Benchmarks executed, results documented
- [ ] Docker image builds and runs locally
- [ ] Deployment tested (local Compose or cloud)
- [ ] Documentation polished and complete
- [ ] Repo organized (`cmd/`, `internal/`, `pkg/`, notes, READMEs)
- [ ] Capstone marked as complete in progress log

---

## 📝 Deliverables
- Final code in `week-12/` (or finalize in `week-11/capstone/`)
- Dockerfile + docker-compose.yml
- Deployment notes (cloud/local)
- Updated root README with polished docs
- Notes in `notes/week-12.md` covering:
  - Deployment experience
  - Benchmark/profiling results
  - Lessons learned from 3-month journey

---

## 📚 References
- Dockerizing Go: https://docs.docker.com/language/golang/
- Go pprof: https://pkg.go.dev/net/http/pprof
- Go Modules: https://go.dev/ref/mod
- AWS ECS: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/docker-basics.html
- GCP Cloud Run: https://cloud.google.com/run/docs

---

🎯 End of Week 12: You’ll have a **fully polished, tested, Dockerized Go project**, with documentation and deployment to the cloud.
