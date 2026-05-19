# Lecture 11: The Standard Library

Go's standard library is massive and high quality. Here are the packages you'll use most.

---

## fmt — Formatted I/O

```go
fmt.Println("Hello")                    // prints + newline
fmt.Printf("Name: %s, Age: %d\n", n, a) // formatted
fmt.Sprintf("Value: %v", x)             // returns string
fmt.Errorf("error: %w", err)            // wraps error

// Common format verbs:
// %v  — default format (any value)
// %T  — type of value
// %d  — integer
// %f  — float (%.2f = 2 decimal places)
// %s  — string
// %q  — quoted string
// %b  — binary, %x — hex
```

---

## strings

```go
import "strings"

strings.Contains("golang", "go")     // true
strings.HasPrefix("hello", "he")     // true
strings.HasSuffix("world", "ld")     // true
strings.ToUpper("hello")             // "HELLO"
strings.ToLower("HELLO")             // "hello"
strings.TrimSpace("  hi  ")          // "hi"
strings.Split("a,b,c", ",")          // ["a","b","c"]
strings.Join([]string{"a","b"}, "-") // "a-b"
strings.Replace("foo", "o", "0", -1) // "f00" (-1 = all)
strings.Count("cheese", "e")         // 3
strings.Index("golang", "lang")      // 2
```

---

## strconv — String Conversions

```go
import "strconv"

n, err := strconv.Atoi("42")       // "42" → 42
s := strconv.Itoa(42)              // 42 → "42"
f, err := strconv.ParseFloat("3.14", 64)
b, err := strconv.ParseBool("true")
```

---

## encoding/json — JSON

```go
import "encoding/json"

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age,omitempty"` // omit if zero
}

// Marshal (struct → JSON bytes)
u := User{Name: "Alice", Email: "alice@example.com"}
data, err := json.Marshal(u)
fmt.Println(string(data)) // {"name":"Alice","email":"alice@example.com"}

// Unmarshal (JSON bytes → struct)
jsonStr := `{"name":"Bob","email":"bob@example.com","age":25}`
var u2 User
err = json.Unmarshal([]byte(jsonStr), &u2)
```

---

## os & io — File I/O

```go
import "os"

// Write a file
err := os.WriteFile("hello.txt", []byte("Hello!\n"), 0644)

// Read a file
data, err := os.ReadFile("hello.txt")
fmt.Println(string(data))

// Environment variables
home := os.Getenv("HOME")
os.Setenv("MY_VAR", "value")

// Command line args (os.Args[0] is the program name)
args := os.Args[1:]
```

---

## net/http — Building HTTP Servers

```go
import "net/http"

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### HTTP Client
```go
resp, err := http.Get("https://api.example.com/data")
if err != nil { log.Fatal(err) }
defer resp.Body.Close()

body, err := io.ReadAll(resp.Body)
fmt.Println(string(body))
```

---

## time

```go
import "time"

now := time.Now()
fmt.Println(now.Format("2006-01-02 15:04:05")) // Go's reference time!
tomorrow := now.Add(24 * time.Hour)
duration := time.Since(start) // elapsed time

time.Sleep(500 * time.Millisecond)
```

> Go's time format uses a specific reference date: `Mon Jan 2 15:04:05 MST 2006`. Memorize the magic numbers: **1-2-3-4-5-6** (month-day-hour-minute-second-year).
