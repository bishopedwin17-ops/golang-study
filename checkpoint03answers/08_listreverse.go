// =============================================================
// EXERCISE: listreverse
//
// QUESTION: Write a function that reverses the order of the 
//           elements of a linked list in place.
//
// CONSTRAINTS:
//           - You MUST modify the actual list pointers (`Next`).
//           - You CANNOT just swap the `Data` inside the nodes.
//           - You must update both `l.Head` and `l.Tail` to 
//             reflect the new reversed state.
// =============================================================

package piscine

// ListReverse reverses the linked list l in place.
func ListReverse(l *List) {
	// ---------------------------------------------------------
	// TODO: 
	// 1. Keep track of three pointers: `prev` (initially nil), 
	//    `current` (initially l.Head), and `next`.
	// 2. Update `l.Tail` to point to the original `l.Head`.
	// 3. Loop while `current != nil`:
	//      a. Save `current.Next` into `next`.
	//      b. Reverse the pointer: `current.Next = prev`.
	//      c. Step forward: `prev = current`, `current = next`.
	// 4. Update `l.Head` to point to `prev`.
	// ---------------------------------------------------------
}
