# 📅 Week 4 – Mini Projects & Consolidation

This week is focused on **bringing everything from Month 1 together**.  
You’ll consolidate Go basics, interfaces, error handling, generics, and testing by building **mini-projects**.

---

## 🎯 Learning Objectives
- [ ] Apply all Go basics in real mini-projects
- [ ] Strengthen understanding of structs, interfaces, and methods
- [ ] Practice error handling with custom errors and panic/recover
- [ ] Use generics in practical scenarios
- [ ] Write comprehensive tests for small projects

---

## 💻 Mini Projects (Hands-On)

1. **Task Manager CLI**
   - Build a CLI-based task manager app
   - Features:
     - Add, list, and delete tasks (store in JSON file)
   - Use structs for `Task`
   - Handle invalid input with proper errors

2. **Student Grades Processor**
   - Read a CSV file of students and grades
   - Calculate average, highest, lowest grades
   - Use generics for numeric operations
   - Write table-driven tests for calculations

3. **Simple HTTP Server**
   - Create a minimal web server using `net/http`
   - Routes:
     - `/hello` → returns “Hello, World!”
     - `/time` → returns server time
   - Demonstrate error handling for invalid requests

4. **Library with Tests**
   - Create a small Go library (e.g., string utils, math utils)
   - Add **unit tests** and measure coverage with `go test -cover`

---

## 📝 Deliverables
- All project code inside `week-4/`
- Notes in `notes/week-4.md` covering:
  - What you built
  - Key learnings
  - Challenges faced
- Update checklist in [root README](../README.md)

---

## 📚 Reference
- [Go by Example – File I/O](https://gobyexample.com/reading-files)
- [Go by Example – JSON](https://gobyexample.com/json)
- [Go by Example – HTTP Servers](https://gobyexample.com/http-servers)
- [Go by Example – Testing](https://gobyexample.com/testing)

---

✅ End of Week 4 (and Month 1): You should now be confident with Go basics, error handling, generics, and testing.
