# 📅 Week 1 – Setup & Language Basics

This week is focused on **getting comfortable with Go’s tooling and core syntax**.  
Goal: Be able to write, build, and run small Go programs confidently.

---

## 🎯 Learning Objectives
- [ ] Install & configure Go toolchain (`go run`, `go build`, `go test`, `go mod`)
- [ ] Learn about environment variables & `GOPATH`, `GOROOT`
- [ ] Get comfortable with `gofmt`, `goimports`, and linters (`golangci-lint`)
- [ ] Understand Go syntax: variables, types, control flow, composite types
- [ ] Learn pointers in Go (differences from C++)

---

## 🛠 Setup & Tooling
- [ ] Install Go 1.22+ → [Download](https://go.dev/dl/)
- [ ] Verify installation: `go version`
- [ ] Setup workspace:
  ```bash
  mkdir -p ~/go/week-1
  cd ~/go/week-1
  go mod init week1
  ```
- [ ] Install helper tools:
  ```bash
  go install golang.org/x/tools/cmd/goimports@latest
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  ```
- [ ] Configure editor (VSCode or GoLand) for auto-formatting & linting

---

## 📖 Core Language Topics
- [ ] Variables, constants, and type inference (`:=`)
- [ ] Basic types: `int`, `float`, `string`, `bool`
- [ ] Control flow: `if`, `for`, `switch`, `range`
- [ ] Composite types: arrays, slices, maps
- [ ] Functions (single & multiple return values)
- [ ] Pointers (`*`, `&`) and how they differ from C++

---

## 💻 Projects (Hands-On)
1. **Hello World**  
   - Basic print program with `fmt.Println`.
   - Show version info using `runtime.Version()`.

2. **CLI Parser**  
   - Build a CLI program that takes a name as input (`os.Args`) and prints:  
     `Hello, <name>!`

3. **JSON File Reader**  
   - Read a local JSON file (e.g., `data.json`).
   - Parse into a Go struct using `encoding/json`.
   - Print results to console.

---

## 📝 Deliverables
- Code for each mini project stored in `week-1/`
- Notes in `notes/week-1.md` covering:
  - Key syntax takeaways
  - Differences you noticed vs Python/C++
  - Gotchas encountered
- Update progress checklist in [root README](../README.md)

---

## 📚 Reference
- [Go Tour](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)

---

✅ End of Week 1: You should be comfortable writing & running simple Go programs with proper formatting and dependency management.
