// =============================================================
// HELPER: visualizers
// USE THESE to debug your linked lists and binary trees.
// Example: PrintList(myList) inside your main function to see
// what your pointers are doing!
// =============================================================

package piscine

import "fmt"

// PrintList prints a linked list in a readable format.
func PrintList(l *List) {
	if l == nil || l.Head == nil {
		fmt.Println("List is empty (Head is nil)")
		return
	}
	current := l.Head
	for current != nil {
		fmt.Printf("[%v] -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

// PrintTree prints a binary tree visually (tilted 90 degrees)
func PrintTree(root *TreeNode, level int) {
	if root == nil {
		return
	}
	// Print right branch
	PrintTree(root.Right, level+1)
	
	// Print current node
	for i := 0; i < level; i++ {
		fmt.Print("      ")
	}
	fmt.Printf("|---- %s\n", root.Data)
	
	// Print left branch
	PrintTree(root.Left, level+1)
}
