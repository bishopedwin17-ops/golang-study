// =============================================================
// EXERCISE: btreeinsertdata
//
// QUESTION: Write a function that inserts data into a binary 
//           search tree in the correct sorted position.
//
// RULES OF BINARY SEARCH TREES:
//           - If the new data is LESS than the current node's data,
//             it goes to the LEFT child.
//           - If the new data is GREATER OR EQUAL, it goes to the
//             RIGHT child.
//
// EDGE CASES:
//           - If the `root` is nil, return the new Node as the 
//             new root.
//           - Ensure the `Parent` pointer of the newly inserted 
//             node is correctly assigned.
// =============================================================

package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

// BTreeInsertData inserts data into the BST and returns the root.
func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	// ---------------------------------------------------------
	// TODO: 
	// 1. If `root == nil`, create and return a `&TreeNode{Data: data}`.
	// 2. If `data < root.Data`:
	//      - Check if `root.Left == nil`. If so, insert it there and 
	//        set the Parent.
	//      - Otherwise, recursively call `BTreeInsertData(root.Left, data)`.
	// 3. Do the same logic for `root.Right` if `data >= root.Data`.
	// 4. Return the original `root` at the end.
	// ---------------------------------------------------------
	return nil
}
