# Lecture 13: Packages & Modules

Understanding how Go organizes code is essential for working on real projects.

---

## Packages

Every `.go` file belongs to a package. Packages are the unit of code organization.

```
myproject/
├── go.mod
├── main.go          (package main)
├── handlers/
│   └── users.go     (package handlers)
└── models/
    └── user.go      (package models)
```

### Exported vs Unexported
- **Uppercase** names are **exported** (public): `User`, `GetUser`, `MaxRetries`
- **Lowercase** names are **unexported** (private): `user`, `getUser`, `maxRetries`

```go
// models/user.go
package models

type User struct {
    Name  string // exported — accessible from other packages
    email string // unexported — only accessible within models package
}

func NewUser(name, email string) *User {
    return &User{Name: name, email: email}
}
```

---

## Modules (`go.mod`)

A module is a collection of packages. `go.mod` defines the module name and dependencies.

```
module github.com/yourname/myproject

go 1.21

require (
    github.com/some/library v1.2.3
)
```

### Common Commands
```bash
go mod init github.com/yourname/myproject  # create a new module
go get github.com/some/library@v1.2.3     # add a dependency
go mod tidy                                # clean up unused dependencies
go mod download                            # download all dependencies
```

---

## Internal Packages

Code in a directory named `internal/` can only be imported by code in the parent directory:

```
myproject/
├── main.go
└── internal/
    └── config/
        └── config.go  ← only importable by myproject code
```

---

## Init Functions

Each package can have an `init()` function that runs automatically before `main()`. Useful for setup:

```go
package mypackage

var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("postgres", os.Getenv("DB_URL"))
    if err != nil {
        log.Fatal("failed to connect to DB:", err)
    }
}
```

> Don't overuse `init()`. Prefer explicit initialization when possible.

---

## Blank Identifier for Side Effects

Some packages register themselves via `init()`. Import them with `_`:
```go
import _ "github.com/lib/pq" // registers the postgres driver
```

---

## Real Project Layout (Standard)

```
myapp/
├── cmd/
│   └── server/
│       └── main.go       ← entry points go here
├── internal/
│   ├── handlers/
│   │   └── user.go
│   └── models/
│       └── user.go
├── pkg/
│   └── utils/            ← shared utilities (exportable)
├── go.mod
└── go.sum
```
