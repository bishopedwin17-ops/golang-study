## 10. File Operations

### Quest 10.1: File I/O Mastery
**Difficulty:** 🟡 Intermediate
**Skills:** File reading/writing, buffered I/O, permissions

```go
// solutions/10-files/fileio.go
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

// Read file line by line
func readLines(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

// Copy file with buffered I/O
func copyFile(src, dst string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer sourceFile.Close()
    
    destFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destFile.Close()
    
    _, err = io.Copy(destFile, sourceFile)
    return err
}

// Append to file
func appendToFile(filename, text string) error {
    file, err := os.OpenFile(filename, 
        os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()
    
    _, err = file.WriteString(text + "\n")
    return err
}

func main() {
    // Test functions
    lines, err := readLines("test.txt")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Read %d lines\n", len(lines))
    }
    
    copyFile("test.txt", "copy.txt")
    appendToFile("log.txt", "New entry")
}
```

**Tasks:**
- [ ] 1. Read a text file line by line with proper error handling
- [ ] 2. Write to a file with different permissions (0644, 0755)
- [ ] 3. Append content to existing file
- [ ] 4. Copy a file using buffered I/O
- [ ] 5. Implement a log rotator that archives when file exceeds size

### Quest 10.2: File System Explorer
**Difficulty:** 🟡 Intermediate

**Tasks:**
- [ ] 1. List all files recursively in a directory
- [ ] 2. Find duplicate files by computing SHA256 hashes
- [ ] 3. Calculate total directory size
- [ ] 4. Find the 10 largest files in a directory tree
- [ ] 5. Build a concurrent file search tool that finds files by name pattern

---
