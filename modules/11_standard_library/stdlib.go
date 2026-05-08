package stdlib

import (
	"encoding/json"
	"strings"
)

// WordCount returns a map of word → count from a sentence.
// Words should be lowercased before counting.
// Hint: strings.Fields, strings.ToLower
// TODO: Implement.
func WordCount(s string) map[string]int {
	// Your code here
	return nil
}

// User is a JSON-serializable user struct.
// TODO: Add json struct tags so that:
//   Name  → "name"
//   Email → "email"
//   Age   → "age" (omit if zero)
type User struct {
	Name  string
	Email string
	Age   int
}

// MarshalUser converts a User to a JSON string.
// TODO: Implement using json.Marshal.
func MarshalUser(u User) (string, error) {
	// Your code here
	return "", nil
}

// UnmarshalUser parses a JSON string into a User.
// TODO: Implement using json.Unmarshal.
func UnmarshalUser(data string) (User, error) {
	// Your code here
	return User{}, nil
}

// keep — suppress unused import
var _ = strings.ToLower
var _ = json.Marshal
