<div align="center">

<img src="https://capsule-render.vercel.app/api?type=waving&color=0:000000,40:00ADD8,100:000000&height=220&section=header&text=go-journey&fontSize=72&fontColor=ffffff&fontAlignY=40&desc=Learning%20Go%20the%20right%20way%20%E2%80%94%20one%20concept%20at%20a%20time&descAlignY=62&descSize=18&descColor=00ADD8&animation=fadeIn" width="100%"/>

</div>

<div align="center">

[![Typing SVG](https://readme-typing-svg.demolab.com?font=JetBrains+Mono&weight=600&size=16&pause=900&color=00ADD8&center=true&vCenter=true&width=650&lines=Structs+%7C+Methods+%7C+Interfaces+%7C+Goroutines;Data+Structures+built+from+scratch+in+Go;From+zero+to+idiomatic+Go+%F0%9F%9A%80;A+Software+Architect+learning+the+language+of+scale)](https://git.io/typing-svg)

</div>

<br/>

<div align="center">

![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Status](https://img.shields.io/badge/Status-Active%20Learning-brightgreen?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-white?style=for-the-badge)
![Concepts](https://img.shields.io/badge/Concepts-Growing%20Daily-00ADD8?style=for-the-badge)

</div>

---

## 📌 About This Repo

This repository documents my **personal learning journey with Go** — from the very first `Hello, World!` to complex data structures, idiomatic patterns, and concurrent systems.

I'm a Software Architect and Fullstack Developer with 20+ years of experience, coming from a TypeScript/Node.js/Java background. This repo exists to:

- Solidify foundational Go concepts through hands-on code
- Build data structures and algorithms from scratch in Go
- Explore what makes Go different: simplicity, performance, and concurrency
- Keep a clean, organized reference for future projects

> _"The best way to learn a language is to build things that hurt a little."_

---



## 📚 Concepts Covered

<br/>

| # | Topic | Status | Notes |
|---|-------|--------|-------|
| 01 | Variables, Types & Constants | ✅ Done | `var`, `:=`, typed vs inferred |
| 02 | Control Flow | ✅ Done | `if`, `for`, `switch`, no `while` |
| 03 | Functions & Multiple Returns | ✅ Done | First-class functions |
| 04 | Arrays & Slices | ✅ Done | Difference, capacity, `append` |
| 05 | Maps | ✅ Done | Key-value, nil safety |
| 06 | Structs | ✅ Done | Value types, embedding |
| 07 | Methods | ✅ Done | Value vs pointer receivers |
| 08 | Interfaces | 🔄 In Progress | Duck typing, composition |
| 09 | Pointers | 🔄 In Progress | `*`, `&`, nil handling |
| 10 | Error Handling | 🔄 In Progress | `error` type, wrapping |
| 11 | Data Structures | 🔄 In Progress | Stack, Queue, Linked List |
| 12 | Goroutines & Channels | 📌 Upcoming | Concurrency primitives |
| 13 | Packages & Modules | ✅ Done | `go mod`, local imports |
| 14 | Generics | 📌 Upcoming | Type parameters |

---

## 🧠 Key Insights So Far

```go
// Go forces you to be explicit. No exceptions, no magic.
// Errors are values — not something to throw and forget.

result, err := doSomething()
if err != nil {
    return fmt.Errorf("context: %w", err)
}

// Interfaces are implicit. If your type has the methods, it satisfies the interface.
// No "implements" keyword. Just behavior.

// Goroutines are cheap. The runtime handles the rest.
go func() {
    // This runs concurrently. No threads, no callbacks, no drama.
}()
```

---

## 🏗️ Data Structures — Built from Scratch

Each data structure is implemented without external packages — pure Go, with tests.

```
Stack      → LIFO using slices + generics
Queue      → FIFO with enqueue/dequeue methods
LinkedList → nodes with pointer chaining
BinaryTree → recursive insert/search/traverse
HashTable  → manual hashing + collision handling
```

---

## 🚀 How to Run

```bash
# Clone the repo
git clone https://github.com/lucaspnbrs/go-journey.git
cd go-journey

# Run any example
cd 03-structs/defining-structs
go run main.go

# Run tests
go test ./...
```

> Requires **Go 1.21+**. Install at [go.dev/dl](https://go.dev/dl)

---

## 🌱 Why Go?

Coming from TypeScript and Java, Go feels like a disciplined reset:

- **No inheritance** — composition only
- **No exceptions** — errors are explicit return values
- **No magic** — what you write is what runs
- **Built-in concurrency** — goroutines are first-class
- **Compiled and fast** — close to C-level performance with high-level ergonomics

For systems that need to scale — and stay readable — Go is the right tool.

---

## 👤 Author

**Lucas Barros** — Software Architect & Fullstack Developer  
Founding Partner @ [JL.Code](https://github.com/lucaspnbrs) & Code.AI

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0A66C2?style=flat-square&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/lucas-barros-30a22330a)
[![Portfolio](https://img.shields.io/badge/Portfolio-000000?style=flat-square&logo=vercel&logoColor=white)](https://lucasbarrosdev-swart.vercel.app/)
[![Gmail](https://img.shields.io/badge/Gmail-EA4335?style=flat-square&logo=gmail&logoColor=white)](mailto:lucaspnbrrs@gmail.com)

---

<div align="center">

<img src="https://capsule-render.vercel.app/api?type=waving&color=0:000000,50:00ADD8,100:000000&height=120&section=footer" width="100%"/>

<sub><code>// Simple, fast, concurrent. Go is not a language you learn once — it's one you keep understanding deeper.</code></sub>

</div>
