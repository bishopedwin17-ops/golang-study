package stdlib

import (
	"strings"
	"testing"
)

func TestWordCount(t *testing.T) {
	result := WordCount("go is great and go is fast")
	if result["go"] != 2 {
		t.Errorf("'go' count = %d, want 2", result["go"])
	}
	if result["is"] != 2 {
		t.Errorf("'is' count = %d, want 2", result["is"])
	}
	if result["great"] != 1 {
		t.Errorf("'great' count = %d, want 1", result["great"])
	}

	// Test case insensitivity
	result2 := WordCount("Go go GO")
	if result2["go"] != 3 {
		t.Errorf("case insensitive 'go' count = %d, want 3", result2["go"])
	}
}

func TestMarshalUser(t *testing.T) {
	u := User{Name: "Alice", Email: "alice@example.com", Age: 30}
	s, err := MarshalUser(u)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(s, `"name":"Alice"`) {
		t.Errorf("marshaled JSON missing 'name' field: %s", s)
	}
	if !strings.Contains(s, `"email":"alice@example.com"`) {
		t.Errorf("marshaled JSON missing 'email' field: %s", s)
	}
}

func TestMarshalUserOmitsZeroAge(t *testing.T) {
	u := User{Name: "Bob", Email: "bob@example.com"} // Age = 0
	s, err := MarshalUser(u)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.Contains(s, `"age"`) {
		t.Errorf("marshaled JSON should omit zero age, got: %s", s)
	}
}

func TestUnmarshalUser(t *testing.T) {
	jsonStr := `{"name":"Charlie","email":"charlie@example.com","age":22}`
	u, err := UnmarshalUser(jsonStr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.Name != "Charlie" {
		t.Errorf("Name = %q, want Charlie", u.Name)
	}
	if u.Age != 22 {
		t.Errorf("Age = %d, want 22", u.Age)
	}
}
