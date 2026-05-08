package storage

import (
	"encoding/json"
	"os"

	"golang-study/modules/15_final_project/task"
)

// Storage defines the interface for task persistence.
type Storage interface {
	Load() ([]task.Task, error)
	Save(tasks []task.Task) error
}

// JSONStorage persists tasks in a JSON file.
type JSONStorage struct {
	FilePath string
}

// Load reads tasks from the JSON file.
// If the file does not exist, return an empty slice (not an error).
// TODO: Implement.
func (s *JSONStorage) Load() ([]task.Task, error) {
	// Hint: use os.ReadFile, then json.Unmarshal
	// Hint: handle os.ErrNotExist
	// Your code here
	return nil, nil
}

// Save writes tasks to the JSON file.
// TODO: Implement.
func (s *JSONStorage) Save(tasks []task.Task) error {
	// Hint: use json.MarshalIndent, then os.WriteFile with perm 0644
	// Your code here
	return nil
}

// keep
var _ = json.Marshal
var _ = os.ReadFile
