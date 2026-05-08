package task

import "testing"

func TestAdd(t *testing.T) {
	tl := NewTaskList()
	if err := tl.Add("Buy milk"); err != nil {
		t.Fatalf("Add failed: %v", err)
	}
	if err := tl.Add("Learn Go"); err != nil {
		t.Fatalf("Add failed: %v", err)
	}
	tasks := tl.All()
	if len(tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(tasks))
	}
	if tasks[0].ID != 1 {
		t.Errorf("first task ID = %d, want 1", tasks[0].ID)
	}
	if tasks[1].ID != 2 {
		t.Errorf("second task ID = %d, want 2", tasks[1].ID)
	}
}

func TestAddEmptyTitle(t *testing.T) {
	tl := NewTaskList()
	if err := tl.Add(""); err == nil {
		t.Error("Add with empty title should return error")
	}
}

func TestComplete(t *testing.T) {
	tl := NewTaskList()
	tl.Add("Task A")
	if err := tl.Complete(1); err != nil {
		t.Fatalf("Complete(1) failed: %v", err)
	}
	if !tl.All()[0].Done {
		t.Error("task should be marked Done")
	}
	if err := tl.Complete(999); err == nil {
		t.Error("Complete with invalid ID should return error")
	}
}

func TestDelete(t *testing.T) {
	tl := NewTaskList()
	tl.Add("Task A")
	tl.Add("Task B")
	if err := tl.Delete(1); err != nil {
		t.Fatalf("Delete(1) failed: %v", err)
	}
	if len(tl.All()) != 1 {
		t.Errorf("expected 1 task after delete, got %d", len(tl.All()))
	}
	if err := tl.Delete(999); err == nil {
		t.Error("Delete with invalid ID should return error")
	}
}

func TestPending(t *testing.T) {
	tl := NewTaskList()
	tl.Add("Task A")
	tl.Add("Task B")
	tl.Complete(1)

	pending := tl.Pending()
	if len(pending) != 1 {
		t.Errorf("expected 1 pending task, got %d", len(pending))
	}
	if pending[0].Title != "Task B" {
		t.Errorf("pending task = %q, want 'Task B'", pending[0].Title)
	}
}
