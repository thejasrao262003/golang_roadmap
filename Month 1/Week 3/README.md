# 📅 Week 3 – Advanced Core Concepts (Methods, Generics & Testing)

This week is focused on **deepening Go fundamentals**:  
methods, advanced interfaces, error handling, and starting with generics & testing.

---

## 🎯 Learning Objectives
- [ ] Understand methods (value vs pointer receivers)
- [ ] Explore deeper use of interfaces & type assertions
- [ ] Practice error handling with `panic` and `recover`
- [ ] Learn the basics of **generics** in Go (functions & types)
- [ ] Write unit tests & table-driven tests with Go’s `testing` package

---

## 📖 Core Language Topics
- [ ] **Methods**
  - Difference between value and pointer receivers
  - When to use which
- [ ] **Interfaces (Advanced)**
  - Type assertions (`x.(T)`)
  - Type switches
- [ ] **Error Handling**
  - Proper use of `panic` and `recover`
  - Best practices (avoid overusing panic)
- [ ] **Generics**
  - Generic functions
  - Type constraints
  - Generic structs & types
- [ ] **Testing**
  - Using Go’s `testing` package
  - Writing table-driven tests
  - Running tests with `go test -v`
  - Coverage with `go test -cover`

---

## 💻 Projects (Hands-On)
1. **Methods Practice**
   - Extend Week 2’s `BankAccount` with methods using pointer receivers
   - Add transfer functionality between accounts

2. **Generic Utility Function**
   - Implement a generic function `Find[T]` that searches for an element in a slice
   - Practice type constraints

3. **Error Handling + Panic/Recover**
   - Write a function that reads a file and recovers gracefully if the file doesn’t exist

4. **Testing Library**
   - Build a small library with error-wrapped functions (e.g., math utils, string utils)
   - Write **unit tests** and **table-driven tests** for coverage

---

## 📝 Deliverables
- Code for each project in `week-3/`
- Notes in `notes/week-3.md` covering:
  - Insights on methods, generics, testing
  - When to use pointer receivers
  - Example table-driven tests
- Update checklist in [root README](../README.md)

---

## 📚 Reference
- [Go Tour – Methods](https://go.dev/tour/methods/1)
- [Go Generics Guide](https://go.dev/doc/tutorial/generics)
- [Go by Example – Testing](https://gobyexample.com/testing)
- [Effective Go – Testing](https://go.dev/doc/effective_go#testing)

---

✅ End of Week 3: You should be comfortable with methods, generics, and able to write unit & table-driven tests for your code.
