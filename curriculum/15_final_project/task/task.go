package task

import (
	"errors"
	"time"
)

// Task represents a single to-do item.
// TODO: Add JSON struct tags.
type Task struct {
	ID        int
	Title     string
	Done      bool
	CreatedAt time.Time
}

// TaskList manages a list of tasks.
type TaskList struct {
	tasks  []Task
	nextID int
}

// NewTaskList creates a new TaskList.
func NewTaskList() *TaskList {
	return &TaskList{nextID: 1}
}

// LoadFrom replaces the internal task list (used when loading from storage).
func (tl *TaskList) LoadFrom(tasks []Task) {
	tl.tasks = tasks
	tl.nextID = 1
	for _, t := range tasks {
		if t.ID >= tl.nextID {
			tl.nextID = t.ID + 1
		}
	}
}

// All returns all tasks.
func (tl *TaskList) All() []Task {
	return tl.tasks
}

// Add creates a new task with an auto-increment ID and appends it.
// TODO: Implement. Return an error if title is empty.
func (tl *TaskList) Add(title string) error {
	// Your code here
	return nil
}

// Complete marks the task with the given ID as done.
// Return an error if the ID is not found.
// TODO: Implement.
func (tl *TaskList) Complete(id int) error {
	// Your code here
	return errors.New("not implemented")
}

// Delete removes the task with the given ID.
// Return an error if not found.
// TODO: Implement.
func (tl *TaskList) Delete(id int) error {
	// Your code here
	return errors.New("not implemented")
}

// Pending returns only tasks where Done == false.
// TODO: Implement.
func (tl *TaskList) Pending() []Task {
	// Your code here
	return nil
}
