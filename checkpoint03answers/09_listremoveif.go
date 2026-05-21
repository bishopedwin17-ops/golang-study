// =============================================================
// EXERCISE: listremoveif
//
// QUESTION: Write a function that removes all elements from a 
//           linked list where the data matches the `data_ref`.
//
// EDGE CASES:
//           - What if the element to remove is the VERY FIRST 
//             element (the Head)? You have to move the Head.
//           - What if there are multiple elements in a row that 
//             need to be removed?
//           - What if the list becomes completely empty? Update Tail.
// =============================================================

package piscine

// ListRemoveIf removes nodes with Data == data_ref.
func ListRemoveIf(l *List, data_ref interface{}) {
	// ---------------------------------------------------------
	// TODO: 
	// 1. Handle removing nodes from the START of the list first.
	//    While `l.Head != nil && l.Head.Data == data_ref`:
	//       `l.Head = l.Head.Next`
	// 2. Once the Head is secure, iterate through the rest of the 
	//    list using a `current` pointer.
	// 3. If `current.Next.Data == data_ref`, skip over it by doing:
	//       `current.Next = current.Next.Next`
	// 4. Make sure to update `l.Tail` if the last element was removed.
	// ---------------------------------------------------------
}
