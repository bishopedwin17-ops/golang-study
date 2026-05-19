## 15. HTTP & Web

### Quest 15.1: REST API Server
**Difficulty:** 🔴 Advanced
**Skills:** HTTP servers, routing, JSON, middleware

```go
// solutions/15-http/api.go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"
    "sync"
    "time"
)

// Todo model
type Todo struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    CreatedAt time.Time `json:"created_at"`
}

// Todo storage
type TodoStore struct {
    mu    sync.RWMutex
    todos map[int]Todo
    nextID int
}

func NewTodoStore() *TodoStore {
    return &TodoStore{
        todos:  make(map[int]Todo),
        nextID: 1,
    }
}

// HTTP Handlers
func (ts *TodoStore) listTodos(w http.ResponseWriter, r *http.Request) {
    ts.mu.RLock()
    defer ts.mu.RUnlock()
    
    todos := make([]Todo, 0, len(ts.todos))
    for _, todo := range ts.todos {
        todos = append(todos, todo)
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todos)
}

func (ts *TodoStore) createTodo(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    ts.mu.Lock()
    todo.ID = ts.nextID
    todo.CreatedAt = time.Now()
    ts.todos[todo.ID] = todo
    ts.nextID++
    ts.mu.Unlock()
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(todo)
}

func (ts *TodoStore) getTodo(w http.ResponseWriter, r *http.Request) {
    // Extract ID from URL: /api/todos/123
    parts := strings.Split(r.URL.Path, "/")
    if len(parts) < 4 {
        http.Error(w, "invalid URL", http.StatusBadRequest)
        return
    }
    
    id, err := strconv.Atoi(parts[3])
    if err != nil {
        http.Error(w, "invalid ID", http.StatusBadRequest)
        return
    }
    
    ts.mu.RLock()
    todo, exists := ts.todos[id]
    ts.mu.RUnlock()
    
    if !exists {
        http.Error(w, "todo not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

func (ts *TodoStore) updateTodo(w http.ResponseWriter, r *http.Request) {
    parts := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(parts[3])
    if err != nil {
        http.Error(w, "invalid ID", http.StatusBadRequest)
        return
    }
    
    var todo Todo
    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    ts.mu.Lock()
    if _, exists := ts.todos[id]; !exists {
        ts.mu.Unlock()
        http.Error(w, "todo not found", http.StatusNotFound)
        return
    }
    todo.ID = id
    ts.todos[id] = todo
    ts.mu.Unlock()
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(todo)
}

func (ts *TodoStore) deleteTodo(w http.ResponseWriter, r *http.Request) {
    parts := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(parts[3])
    if err != nil {
        http.Error(w, "invalid ID", http.StatusBadRequest)
        return
    }
    
    ts.mu.Lock()
    delete(ts.todos, id)
    ts.mu.Unlock()
    
    w.WriteHeader(http.StatusNoContent)
}

// Middleware
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    }
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        
        next(w, r)
    }
}

func main() {
    store := NewTodoStore()
    
    // Routes
    http.HandleFunc("/api/todos", corsMiddleware(loggingMiddleware(func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            store.listTodos(w, r)
        case http.MethodPost:
            store.createTodo(w, r)
        default:
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        }
    })))
    
    http.HandleFunc("/api/todos/", corsMiddleware(loggingMiddleware(func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            store.getTodo(w, r)
        case http.MethodPut:
            store.updateTodo(w, r)
        case http.MethodDelete:
            store.deleteTodo(w, r)
        default:
            http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
        }
    })))
    
    fmt.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**Tasks:**
- [ ] 1. Create a basic HTTP server with multiple routes
- [ ] 2. Handle different HTTP methods (GET, POST, PUT, DELETE)
- [ ] 3. Parse JSON request bodies and return JSON responses
- [ ] 4. Implement full CRUD API for todo items
- [ ] 5. Add middleware: logging, CORS, authentication
- [ ] 6. Implement pagination and filtering for list endpoints

### Quest 15.2: HTTP Client
**Difficulty:** 🟡 Intermediate

```go
// solutions/15-http/client.go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
)

// HTTP client with retry logic
type Client struct {
    httpClient *http.Client
    maxRetries int
}

func NewClient(timeout time.Duration, maxRetries int) *Client {
    return &Client{
        httpClient: &http.Client{
            Timeout: timeout,
        },
        maxRetries: maxRetries,
    }
}

func (c *Client) GetWithRetry(ctx context.Context, url string) (*http.Response, error) {
    var resp *http.Response
    var err error
    
    for i := 0; i <= c.maxRetries; i++ {
        req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
        if err != nil {
            return nil, err
        }
        
        resp, err = c.httpClient.Do(req)
        if err == nil && resp.StatusCode < 500 {
            return resp, nil
        }
        
        if i < c.maxRetries {
            backoff := time.Duration(i+1) * time.Second
            fmt.Printf("Retry %d after %v\n", i+1, backoff)
            time.Sleep(backoff)
        }
    }
    
    return resp, err
}

// Make API request and decode JSON
func fetchTodo(id int) (*Todo, error) {
    client := NewClient(10*time.Second, 2)
    ctx := context.Background()
    
    resp, err := client.GetWithRetry(ctx, 
        fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", id))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    body, _ := io.ReadAll(resp.Body)
    
    var todo Todo
    if err := json.Unmarshal(body, &todo); err != nil {
        return nil, err
    }
    
    return &todo, nil
}

type Todo struct {
    UserID    int    `json:"userId"`
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Completed bool   `json:"completed"`
}

func main() {
    todo, err := fetchTodo(1)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Todo: %+v\n", todo)
}
```

**Tasks:**
- [ ] 1. Create HTTP client with retry logic and exponential backoff
- [ ] 2. Implement timeout and context cancellation
- [ ] 3. Build rate-limited API consumer
- [ ] 4. Create concurrent downloader for multiple URLs
- [ ] 5. Implement custom HTTP client with connection pooling

---
