package storage

import (
	"os"
	"path/filepath"
	"testing"

	"golang-study/modules/15_final_project/task"
)

func TestJSONStorageRoundTrip(t *testing.T) {
	// Use a temp file
	dir := t.TempDir()
	s := &JSONStorage{FilePath: filepath.Join(dir, "tasks.json")}

	// Load from non-existent file — should return empty list
	tasks, err := s.Load()
	if err != nil {
		t.Fatalf("Load from non-existent file failed: %v", err)
	}
	if len(tasks) != 0 {
		t.Errorf("expected 0 tasks, got %d", len(tasks))
	}

	// Save some tasks
	input := []task.Task{
		{ID: 1, Title: "Buy milk", Done: false},
		{ID: 2, Title: "Learn Go", Done: true},
	}
	if err := s.Save(input); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// Load them back
	loaded, err := s.Load()
	if err != nil {
		t.Fatalf("Load after Save failed: %v", err)
	}
	if len(loaded) != 2 {
		t.Fatalf("expected 2 tasks after load, got %d", len(loaded))
	}
	if loaded[0].Title != "Buy milk" {
		t.Errorf("loaded[0].Title = %q, want 'Buy milk'", loaded[0].Title)
	}
	if !loaded[1].Done {
		t.Error("loaded[1].Done should be true")
	}
}

func TestJSONStorageFileCreated(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "tasks.json")
	s := &JSONStorage{FilePath: path}

	if err := s.Save([]task.Task{}); err != nil {
		t.Fatalf("Save failed: %v", err)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error("JSON file was not created")
	}
}
