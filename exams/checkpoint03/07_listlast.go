// =============================================================
// EXERCISE: listlast
//
// QUESTION: Write a function that returns the data from the last 
//           element of a linked list.
//
// EDGE CASES:
//           - If the list is empty (l.Head == nil), you must 
//             return `nil`.
// =============================================================

package piscine

// ListLast returns the data from the last element of l.
func ListLast(l *List) interface{} {
	// ---------------------------------------------------------
	// TODO: 
	// 1. Check if the list is empty. If so, return nil.
	// 2. Because `List` has a `Tail` pointer, you can access the 
	//    last node directly via `l.Tail.Data`.
	// 3. If the List struct didn't have a Tail, you would have to
	//    iterate until `current.Next == nil`.
	// ---------------------------------------------------------
	return nil
}
