// =============================================================
// EXERCISE: advancedsortwordarr
//
// QUESTION: Write a function that sorts a slice of strings based 
//           on the function passed as a parameter `f`.
//           The function `f` returns:
//             1  if a > b
//             0  if a == b
//            -1  if a < b
//
// CONSTRAINTS:
//           - You CANNOT use the `sort` package.
//           - This tests your ability to use func types as arguments.
// =============================================================

package piscine

// AdvancedSortWordArr sorts the slice a using comparison function f.
func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	// ---------------------------------------------------------
	// TODO: 
	// 1. Similar to SortWordArr, implement a sorting loop.
	// 2. Instead of `a[i] > a[j]`, use `f(a[i], a[j]) == 1` to 
	//    determine if elements need swapping.
	// ---------------------------------------------------------
}
