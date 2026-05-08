# 01-edu Piscine Golang Checkpoint Exercises

This is a compilation of past checkpoint and quest exercises from the 01-edu / Zone01 Piscine curriculum. 

### Rules of the Piscine:
1. **No Standard Library:** Most of these checkpoints **forbid** importing anything other than `github.com/01-edu/z01`. You cannot use `fmt.Println`, `strconv.Atoi`, or `strings.Reverse`. You have to build them yourself!
2. **Runes are Everything:** You must understand that a `string` in Go is a read-only slice of bytes, but you often need to iterate over it as `rune`s (characters) to handle things properly.
3. **Strict Formatting:** The "moulinette" (the automated grader) expects exactly the output requested. An extra space or missing newline (`\n`) means a failure.

---

## 🟢 Level 1: Basics & Characters
These focus on basic syntax, types, and printing using the `z01` package.

### 1. `printstr`
*   **Instructions:** Write a function that prints one by one the characters of a `string` on the screen.
*   **Expected function:** `func PrintStr(s string)`

### 2. `printnbr`
*   **Instructions:** Write a function that prints an `int` passed in parameter. All possible values of type `int` have to go through. You cannot convert to `int64`.
*   **Expected function:** `func PrintNbr(n int)`

---

## 🟡 Level 2: Pointers & Basic Math
These introduce pointers (memory addresses) and basic mathematical operations.

### 3. `swap`
*   **Instructions:** Write a function that takes two pointers to an `int` (`*int`) and swaps their contents.
*   **Expected function:** `func Swap(a *int, b *int)`

### 4. `divmod`
*   **Instructions:** Write a function that will be formatted as below. This function will divide the `int` `a` and `b`. The result of this division will be stored in the `int` pointed by `div`. The remainder of this division will be stored in the `int` pointed by `mod`.
*   **Expected function:** `func DivMod(a int, b int, div *int, mod *int)`

---

## 🟠 Level 3: Loops & Logic
These require nested loops and careful formatting.

### 5. `printcomb`
*   **Instructions:** Write a function that prints, in ascending order and on a single line: all unique combinations of three different digits so that, the first digit is lower than the second, and the second is lower than the third. (e.g., `012, 013, 014...`). `000` and `987` are invalid.
*   **Expected function:** `func PrintComb()`

### 6. `printcomb2`
*   **Instructions:** Write a function that prints in ascending order and on a single line: all possible combinations of two different two-digit numbers. (e.g., `00 01, 00 02...`).
*   **Expected function:** `func PrintComb2()`

---

## 🔴 Level 4: Strings & Conversions
These involve manipulating strings, runes, and converting between types.

### 7. `strlen`
*   **Instructions:** Write a function that counts the characters of a string and returns that count.
*   **Expected function:** `func StrLen(s string) int`

### 8. `strrev`
*   **Instructions:** Write a function that reverses a string and returns it.
*   **Expected function:** `func StrRev(s string) string`

### 9. `atoi`
*   **Instructions:** Write a function that simulates the behaviour of the `Atoi` function in Go. `Atoi` transforms a number represented as a `string` in a number represented as an `int`. It must handle `+` and `-` signs. It returns `0` if the string is invalid.
*   **Expected function:** `func Atoi(s string) int`

---

## 🟣 Level 5: Recursion & Algorithms
These exercises explicitly require recursion or optimization.

### 10. `recursivefactorial`
*   **Instructions:** Write a **recursive** function that returns the factorial of the `int` passed as parameter. Errors (non-possible values or overflows) should return `0`. `for` loops are forbidden.
*   **Expected function:** `func RecursiveFactorial(nb int) int`

### 11. `fibonacci`
*   **Instructions:** Write a **recursive** function that returns the value at the position `index` in the fibonacci sequence. The sequence starts: 0, 1, 1, 2, 3... Negative index returns `-1`. `for` loops are forbidden.
*   **Expected function:** `func Fibonacci(index int) int`

### 12. `isprime`
*   **Instructions:** Write a function that returns `true` if the `int` passed as parameter is a prime number. Otherwise it returns `false`. The function must be optimized in order to avoid time-outs with the tester.
*   **Expected function:** `func IsPrime(nb int) bool`

---

## 🟤 Level 6: Linked Lists (Advanced)
Later in the Piscine, you have to build data structures from scratch.

### 13. `sortedlistmerge`
*   **Instructions:** Write a function `SortedListMerge` that merges two sorted linked lists into one sorted linked list.
*   **Expected function:** `func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI`
