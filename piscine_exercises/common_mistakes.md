# 🐛 Common Mistakes & Gotchas

Things that tripped up during solving the 39 checkpoint exercises.
Keep adding to this as you study — it's your personal bug diary.

---

## 1. `byte` vs `rune` — The silent killer

**The mistake:**
```go
// WRONG — s[i] gives a byte, not a rune
for i := 0; i < len(s); i++ {
    ch := s[i] // type: byte
}
```

**The fix:**
```go
// CORRECT — range on a string gives runes automatically
for _, ch := range s {
    // ch is type: rune — handles multi-byte characters safely
}
```

**Why it matters:** ASCII characters (a-z, 0-9) happen to work with bytes.
But the moment there's a special character (é, ñ, etc.), bytes break silently.
**Always use `range` on strings unless you have a specific reason not to.**

---

## 2. Forgetting the trailing `\n`

**The mistake:** Returning `"hello"` instead of `"hello\n"`.

**The fix:**
- If the function returns a string: `return result + "\n"`
- If the program prints: use `fmt.Println(result)` (adds `\n` automatically)
- `fmt.Print(result)` does NOT add `\n` — use it only when `\n` is already in the string

**Exercises this burned:** `wordflip`, `fromto`, `fifthandskip`, `thirdtimeisacharm`

---

## 3. Wrong `os.Args` length check

**The mistake:**
```go
if len(os.Args) != 1 { // WRONG — os.Args[0] is the program name!
```

**The fix:**
```go
// os.Args[0] = program name (always there)
// os.Args[1] = first user argument
// So for 1 user argument:
if len(os.Args) != 2 { ... }
// For 2 user arguments:
if len(os.Args) != 3 { ... }
```

**Exercises this affects:** every single Program-type exercise.

---

## 4. Modifying a slice while using it as both input and output

**The mistake:**
```go
// WRONG — append may modify the original slice
func ConcatSlice(s1, s2 []int) []int {
    return append(s1, s2...) // s1 might get mutated!
}
```

**The fix:**
```go
// SAFE — make a new slice first
result := make([]int, len(s1)+len(s2))
copy(result, s1)
copy(result[len(s1):], s2)
// OR just trust append when you own the result variable
result := append(s1, s2...)
```

**The rule:** `append` is safe when the result is a NEW variable you control.
It's dangerous when the caller still holds a reference to the original.

---

## 5. Map lookup on missing key returns zero value, not a crash

**The thing to know:**
```go
m := make(map[rune]bool)
// Looking up a key that doesn't exist returns false (zero value for bool)
val := m['x'] // val = false — no crash, no error
```

This is actually USEFUL for the `inter` / `union` / `hiddenp` pattern:
```go
if !seen[ch] { // false if ch is not in map → same as "not seen yet"
    seen[ch] = true
}
```

---

## 6. Integer division truncates (doesn't round)

**The mistake:**
```go
n := 7
result := n / 2 // result = 3, NOT 3.5
```

**Why it matters in `fprime`:** when checking `d*d <= n`, make sure both sides
are ints. Mixing int and float causes compile errors in Go.

---

## 7. Stack for bracket matching — order of pop matters

**The mistake:** popping the wrong operand first in `rpncalc`.

```go
// WRONG — a and b are swapped
a := stack[len(stack)-1] // this is actually b (the RIGHT operand)
b := stack[len(stack)-2] // this is actually a (the LEFT operand)
result := a - b // 3 - 1 gives 2, but should be 1 - 3 = -2
```

**The fix:**
```go
// The LAST pushed is the RIGHT operand
b := stack[len(stack)-1]
a := stack[len(stack)-2]
result := a - b // correct: left - right
```

---

## 8. `string(someInt)` is NOT `Itoa`

**The mistake:**
```go
n := 65
s := string(n) // s = "A" — not "65"!
```

`string(rune)` converts a rune/byte to the CHARACTER with that Unicode value.
It does NOT convert the number to its digit string.

**The fix:**
```go
// To get "65" from 65, convert digit by digit:
digit := n % 10     // 5
char := string(rune('0' + digit)) // "5"
```

---

## 9. `len(s)` counts BYTES, not characters

```go
s := "hello"
len(s) // → 5 ✓ (all ASCII, 1 byte each)

s2 := "héllo"
len(s2) // → 6 ✗ (é is 2 bytes in UTF-8, but 1 visible character)
```

**Fix:** use `len([]rune(s))` to count visible characters (runes).
For piscine exercises with ASCII-only input, `len(s)` is fine.

---

## 10. Slice out of bounds — the silent crash

```go
s := []int{1, 2, 3}
// CRASH: index 3 is out of range [0, 2]
x := s[3]

// ALSO CRASH: slicing past the end
sub := s[1:5] // s only has indices 0,1,2
```

**Guard pattern:**
```go
if i < len(s) {
    x := s[i]
}
```

**This bit me in:** `fifthandskip`, `slice`, `revconcatalternate`
