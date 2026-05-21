# Instructions: Collections

## Objective

Open `collections.go` and implement the three functions:

1. **`Sum(nums []int) int`**  
   Return the sum of all elements using a `for range` loop.

2. **`CountWords(sentence string) map[string]int`**  
   Split the sentence into words using `strings.Fields(sentence)`.  
   Return a map of `word → frequency`.  
   Remember to `import "strings"` at the top.

3. **`Filter(nums []int, keep func(int) bool) []int`**  
   Return a new slice containing only elements where `keep(element)` is `true`.  
   Use `append` to build the result.

## How to Verify
```
cd modules/04_collections
go test -v
```

---

## 🐲 BOSS FIGHT: Piscine Challenges

**Rules:** You can ONLY use the `github.com/01-edu/z01` package to print (specifically `z01.PrintRune`). You cannot use `fmt` or `strings` or `strconv` for these functions.

### 1. `printstr`
Write a function that prints one by one the characters of a `string` on the screen.
- Expected: `func PrintStr(s string)`

### 2. `strlen`
Write a function that counts the characters of a string and returns that count.
- Expected: `func StrLen(s string) int`

### 3. `strrev`
Write a function that reverses a string and returns it.
- Expected: `func StrRev(s string) string`

### 4. `atoi`
Write a function that simulates the behaviour of the `Atoi` function in Go. `Atoi` transforms a number represented as a `string` in a number represented as an `int`.
- It must handle `+` and `-` signs. 
- It returns `0` if the string is not a valid number (contains non-digits other than a leading sign).
- Expected: `func Atoi(s string) int`

Open `piscine.go` to implement them, and test them with `go test -v -run Piscine`.
