package main

import (
	"fmt"
	"os"
	"strconv"

	"golang-study/modules/15_final_project/storage"
	"golang-study/modules/15_final_project/task"
)

func main() {
	store := &storage.JSONStorage{FilePath: "tasks.json"}

	tl := task.NewTaskList()
	tasks, err := store.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading tasks:", err)
		os.Exit(1)
	}
	tl.LoadFrom(tasks)

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: tasks <add|list|done|delete> [args...]")
		os.Exit(0)
	}

	cmd := args[0]

	switch cmd {
	case "add":
		// TODO: Get the title from args[1] and call tl.Add()
		// Print a confirmation message.
		// Your code here

	case "list":
		// TODO: Print all pending tasks in format: [ID] Title
		// If no pending tasks, print "No pending tasks."
		// Your code here

	case "done":
		// TODO: Parse args[1] as an int ID, call tl.Complete()
		// Your code here

	case "delete":
		// TODO: Parse args[1] as an int ID, call tl.Delete()
		// Your code here

	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		os.Exit(1)
	}

	// Save after every command
	if err := store.Save(tl.All()); err != nil {
		fmt.Fprintln(os.Stderr, "Error saving tasks:", err)
		os.Exit(1)
	}
}

// keep
var _ = strconv.Atoi
