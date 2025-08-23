# 📅 Week 5 – Modules & Standard Library (Deep Dive)

This week focuses on **Go modules** and a **deep dive into key standard library packages**.  
You’ll also build a **file watcher / parser** that uses these libraries in combination.

---

## 🎯 Learning Objectives
- [ ] Understand and manage **Go modules** (`go.mod`, `go.sum`) and **vendoring**
- [ ] Practice with stdlib packages: `encoding/json`, `time`, `os`, `regexp`, and `embed`
- [ ] Structure a small project with modules and clean package boundaries
- [ ] Build a **file watcher or parser** using these libraries

---

## 🧰 Go Modules & Project Setup
- [ ] Initialize a module:
  ```bash
  mkdir -p week-5/file-tools && cd week-5/file-tools
  go mod init github.com/<your-username>/file-tools
  ```
- [ ] Add and tidy dependencies (if any third-party libs used):
  ```bash
  go get
  go mod tidy
  ```
- [ ] Vendor dependencies (optional, for reproducible builds):
  ```bash
  go mod vendor
  ```
- [ ] Inspect files:
  - `go.mod`: module path, Go version, required deps
  - `go.sum`: checksums for verification

> Tip: Prefer stdlib wherever possible this week; add third-party deps only if needed.

---

## 📦 Standard Library Deep Dive

### 1) `encoding/json`
- Marshal/unmarshal structs with tags:
  ```go
  type User struct {
      ID   int    `json:"id"`
      Name string `json:"name"`
  }
  ```
- Stream decoders/encoders for big files (`json.Decoder`, `json.Encoder`)

### 2) `time`
- Parsing/formatting timestamps (`time.Parse`, `time.Format`)
- Durations & timers (`time.Duration`, `time.NewTicker`)

### 3) `os`
- File & directory operations (`os.Open`, `os.ReadDir`, `os.Stat`)
- Env vars (`os.Getenv`), exit codes (`os.Exit`)

### 4) `regexp`
- Compile vs MustCompile, find vs replace, submatch extraction

### 5) `embed` (a.k.a. `go:embed`)
- Embed sample inputs or config files at build time:
  ```go
  //go:embed samples/*.json
  var sampleFS embed.FS
  ```

---

## 💻 Project: File Watcher / Parser (Choose One)

### Option A: File Watcher (Stdlib-first)
**Goal:** Monitor a directory and process new/changed `.json` or `.txt` files.

Features:
- Poll a directory every N seconds (use `time.Ticker`)
- Detect new/modified files (`os.ReadDir`, `os.Stat`)
- For `.json`, unmarshal and print summary (`encoding/json`)
- For `.txt`, extract patterns with regex (`regexp`)
- Graceful shutdown via `os.Signal` handling

Suggested structure:
```
week-5/file-tools/
├── cmd/filewatcher/           # main package
│   └── main.go
├── internal/files/            # file ops (stat, read, list)
├── internal/parsers/          # json & text parsers
├── samples/                   # sample files (optionally embedded)
└── go.mod
```

### Option B: Batch Parser (CLI)
**Goal:** Parse a directory of files once and generate a JSON report.

Features:
- CLI flags for `--in` and `--out`
- Read all files under `--in`
- Extract data with regex and produce `report.json`
- Include build-time embedded templates or example inputs via `embed`

---

## ✅ Checkpoints
- [ ] `go mod init` and `go mod tidy` run cleanly
- [ ] Optional: `go mod vendor` creates a `vendor/` folder
- [ ] Project builds with `go build ./...`
- [ ] You can run: `go run ./cmd/filewatcher`
- [ ] JSON parsing and regex extraction both demonstrated

---

## 📝 Deliverables
- Code in `week-5/` (with `cmd/`, `internal/`, `samples/`)
- Notes in `notes/week-5.md`, including:
  - Module concepts you used (tidy, vendor, replace, etc.)
  - Snippets for JSON parsing, `time` usage, regex you wrote
  - What you embedded with `go:embed` (if used)

---

## 📚 Reference
- Modules: https://go.dev/doc/modules
- JSON: https://pkg.go.dev/encoding/json
- Time: https://pkg.go.dev/time
- OS: https://pkg.go.dev/os
- Regexp: https://pkg.go.dev/regexp
- Embed: https://pkg.go.dev/embed

---

✅ End of Week 5: You can structure a Go module properly and use key stdlib packages to build real utilities.
