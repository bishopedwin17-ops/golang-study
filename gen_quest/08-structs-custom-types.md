## 8. Structs & Custom Types

### Quest 8.1: Library Management System
**Difficulty:** 🟡 Intermediate
**Skills:** Structs, methods, custom types

```go
// solutions/08-structs/library.go
package main

import (
    "fmt"
    "time"
)

type Book struct {
    ISBN     string
    Title    string
    Author   string
    Pages    int
    Available bool
    DueDate  *time.Time
}

type Member struct {
    ID            int
    Name          string
    BorrowedBooks []Book
}

// Method on Book
func (b Book) IsLongBook() bool {
    return b.Pages > 300
}

// Method on Book with pointer receiver
func (b *Book) Checkout(dueDate time.Time) error {
    if !b.Available {
        return fmt.Errorf("book %s is not available", b.Title)
    }
    b.Available = false
    b.DueDate = &dueDate
    return nil
}

func (b *Book) Return() {
    b.Available = true
    b.DueDate = nil
}

// String method for pretty printing
func (b Book) String() string {
    return fmt.Sprintf("Book[%s] by %s (%d pages)", 
        b.Title, b.Author, b.Pages)
}

type Library struct {
    Books   []Book
    Members map[int]*Member
}

func (l *Library) AddBook(book Book) {
    l.Books = append(l.Books, book)
}

func (l *Library) FindByAuthor(author string) []Book {
    var result []Book
    for _, book := range l.Books {
        if book.Author == author {
            result = append(result, book)
        }
    }
    return result
}

func main() {
    // Build and test library system
    lib := &Library{
        Members: make(map[int]*Member),
    }
    
    lib.AddBook(Book{
        ISBN:     "978-0134190440",
        Title:    "The Go Programming Language",
        Author:   "Alan Donovan",
        Pages:    380,
        Available: true,
    })
    
    fmt.Println(lib.Books[0])
    fmt.Printf("Is long book? %t\n", lib.Books[0].IsLongBook())
}
```

**Tasks:**
- [ ] 1. Create `Book` struct with: Title, Author, ISBN, Pages, Available
- [ ] 2. Create `Member` struct with: ID, Name, BorrowedBooks slice
- [ ] 3. Implement methods: `Checkout()`, `Return()`, `IsLongBook()`, `String()`
- [ ] 4. Build `Library` struct with search functionality
- [ ] 5. Add fines calculation based on overdue dates

### Quest 8.2: Composition Over Inheritance
**Difficulty:** 🟡 Intermediate

```go
// solutions/08-structs/vehicles.go
package main

import "fmt"

// Base struct
type Vehicle struct {
    Speed int
    Model string
}

func (v Vehicle) Start() {
    fmt.Printf("%s starting...\n", v.Model)
}

// Embedding Vehicle in Car
type Car struct {
    Vehicle
    Doors int
}

// Method overriding
func (c Car) Start() {
    fmt.Printf("Car %s starting with %d doors\n", c.Model, c.Doors)
}

type Boat struct {
    Vehicle
    Length float64
}

type Plane struct {
    Vehicle
    Altitude int
}

// Interface
type Transport interface {
    Start()
    Stop()
}

func (c Car) Stop() {
    fmt.Printf("Car %s stopping\n", c.Model)
}

func (b Boat) Stop() {
    fmt.Printf("Boat %s docking\n", b.Model)
}

func main() {
    car := Car{
        Vehicle: Vehicle{Speed: 200, Model: "Tesla"},
        Doors:   4,
    }
    
    car.Start()  // Uses overridden method
    car.Vehicle.Start()  // Uses base method
    
    // Interface usage
    var t Transport = car
    t.Start()
    t.Stop()
}
```

**Tasks:**
- [ ] 1. Create `Vehicle` struct with Speed and Model
- [ ] 2. Embed it in `Car`, `Boat`, `Plane` structs with unique fields
- [ ] 3. Add unique methods for each (Drive, Sail, Fly)
- [ ] 4. Demonstrate method overriding on embedded struct
- [ ] 5. Create `Transport` interface and implement for all vehicles
- [ ] 6. Write a factory function that creates vehicles based on type string

---
