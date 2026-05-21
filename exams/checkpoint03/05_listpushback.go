// =============================================================
// EXERCISE: listpushback
//
// QUESTION: Write a function that inserts a new element `NodeL` at 
//           the end of the list.
//
// STRUCTURES:
//           The standard Piscine linked list structures are provided.
//           A List contains a Head (start) and Tail (end).
//           A NodeL contains Data (interface{}) and Next (pointer).
//
// EDGE CASES:
//           - What if the list is completely empty (l.Head == nil)?
//             You must update BOTH the Head and the Tail.
// =============================================================

package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

// ListPushBack inserts a new element at the end of the list.
func ListPushBack(l *List, data interface{}) {
	// ---------------------------------------------------------
	// TODO: 
	// 1. Create a new NodeL instance using `&NodeL{Data: data}`.
	// 2. Check if the list is empty (`l.Head == nil`).
	// 3. If empty, point both `l.Head` and `l.Tail` to the new node.
	// 4. If not empty, attach the new node to the current `Tail.Next`.
	// 5. Update `l.Tail` to be the new node.
	// ---------------------------------------------------------
}
