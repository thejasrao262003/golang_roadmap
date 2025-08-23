# 📅 Week 2 – Structs, Interfaces & Error Handling

This week is focused on **learning Go’s core building blocks for abstraction and reliability**:  
structs, interfaces, and idiomatic error handling.

---

## 🎯 Learning Objectives
- [ ] Understand **structs** and methods in Go
- [ ] Learn about **interfaces** and how Go achieves polymorphism without inheritance
- [ ] Master Go’s **error handling** patterns (multi-return values, wrapping, custom errors)
- [ ] Practice using `panic` and `recover` responsibly

---

## 📖 Core Language Topics
- [ ] **Structs**
  - Defining and instantiating structs
  - Struct methods (value vs pointer receivers)
  - Nested structs & composition
- [ ] **Interfaces**
  - Implicit implementation
  - Interface variables & type assertions
  - Empty interface `interface{}` and `any`
  - Using `fmt.Stringer` as an example
- [ ] **Error Handling**
  - Returning `error` as the last return value
  - Creating custom errors with `errors.New` and `fmt.Errorf`
  - Error wrapping and unwrapping (`errors.Is`, `errors.As`)
  - `panic` vs `recover` vs error returns

---

## 💻 Projects (Hands-On)
1. **Bank Account Struct**
   - Create a `BankAccount` struct with fields: `owner`, `balance`
   - Add methods: `Deposit`, `Withdraw`, `GetBalance`
   - Ensure negative withdrawals throw an error

2. **Shape Interface**
   - Define a `Shape` interface with `Area()` and `Perimeter()`
   - Implement for `Rectangle` and `Circle`
   - Demonstrate polymorphism with a slice of `Shape`

3. **Custom Error Example**
   - Build a small function that parses user input (e.g., string → int)
   - Return proper errors with `fmt.Errorf` for invalid inputs
   - Showcase error wrapping

---

## 📝 Deliverables
- Code for each project in `week-2/`
- Notes in `notes/week-2.md` covering:
  - Key struct & interface insights
  - Error handling patterns learned
  - Differences noticed from Python & C++ OOP
- Update checklist in [root README](../README.md)

---

## 📚 Reference
- [Effective Go – Errors](https://go.dev/doc/effective_go#errors)
- [Go by Example – Structs](https://gobyexample.com/structs)
- [Go by Example – Interfaces](https://gobyexample.com/interfaces)
- [Go by Example – Errors](https://gobyexample.com/errors)

---

✅ End of Week 2: You should be confident using structs, interfaces, and handling errors in an idiomatic Go style.
