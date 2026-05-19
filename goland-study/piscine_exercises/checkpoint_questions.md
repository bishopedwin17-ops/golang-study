# 01-edu / Zone01 Piscine Go — Checkpoint 02 Exercise Pool

> **Source:** [kinoz01/zone01-Piscine – v2/checkpoint](https://github.com/kinoz01/zone01-Piscine/tree/master/v2/checkpoint) · [01-edu/public – subjects](https://github.com/01-edu/public/tree/master/subjects)
>
> This file contains the **actual Checkpoint 02 exercise pool** — the set of problems the exam randomly draws from. During the real exam you are given a random subset from this list under a time limit.

---

## 📌 Checkpoint Rules (Always Apply)

1. **No standard library for output.** Use `github.com/01-edu/z01` (`z01.PrintRune`) unless the exercise explicitly accepts `fmt`.
2. **Programs vs Functions.** Some exercises ask for a `main` program (`go run .`); others ask for a package function (`func Foo(...)`). Read carefully.
3. **Exact output.** The moulinette is byte-exact. Missing `\n`, extra spaces = fail.
4. **No `for` loops where recursion is specified.** Pay attention to constraints.

---

## 🟢 Tier 1 — Basic String & Character Output

### `countalpha`
- **Type:** Function
- **Instructions:** Write a function `CountAlpha()` that takes a `string` and returns the number of **alphabetic** characters in it.
- **Expected function:**
  ```go
  func CountAlpha(s string) int
  ```
- **Example:**
  ```
  CountAlpha("Hello world") → 10
  CountAlpha("H1e2l3l4o")  → 5
  ```

---

### `onlya` / `onlyb` / `onlyf` / `onlyz`
- **Type:** Program (main)
- **Instructions:** Each of these programs prints only the character in its name (`a`, `b`, `f`, or `z`) from the alphabet, using `z01.PrintRune`. The pattern is: print just that single char followed by `\n`.
- **Note:** These are warm-up basics that test your ability to use the `z01` package.

---

### `only1`
- **Type:** Program (main)
- **Instructions:** Print the digit `1` followed by a newline using `z01.PrintRune`.

---

## 🟡 Tier 2 — String Manipulation (Functions)

### `atoi`
- **Type:** Function
- **Instructions:** Simulate Go's `strconv.Atoi`. Transform a number as a `string` into an `int`.
  - Return `0` if the string is not a valid number (e.g. contains non-digits, `++`, `--`).
  - Handle `+` and `-` signs.
  - Only return the `int` (no error return needed).
- **Expected function:**
  ```go
  func Atoi(s string) int
  ```
- **Example:**
  ```
  Atoi("12345")    → 12345
  Atoi("+1234")    → 1234
  Atoi("-1234")    → -1234
  Atoi("012 345")  → 0
  Atoi("Hello!")   → 0
  Atoi("++1234")   → 0
  ```

---

### `itoa`
- **Type:** Function
- **Instructions:** Simulate Go's `strconv.Itoa`. Transform an `int` into its `string` representation.
  - Handle negative numbers.
- **Expected function:**
  ```go
  func Itoa(n int) string
  ```
- **Example:**
  ```
  Itoa(12345)   → "12345"
  Itoa(0)       → "0"
  Itoa(-1234)   → "-1234"
  ```

---

### `itoabase`
- **Type:** Function
- **Instructions:** Write a function that converts an `int` to its string representation in a given base.
- **Expected function:**
  ```go
  func ItoaBase(nbr int, base int) string
  ```
- **Note:** Handle invalid bases (< 2 or > 16) and negative numbers.

---

### `cameltosnakecase`
- **Type:** Function
- **Instructions:** Convert a `string` from `camelCase` to `snake_case`.
  - Return the string **unchanged** if it is not valid camelCase.
  - Valid camelCase rules:
    - Does NOT end on a capital letter (`CamelCasE` → invalid).
    - No two consecutive capitals (`CamelCAse` → invalid).
    - No numbers or punctuation (`camelCase1` → invalid).
  - Accepts both `lowerCamelCase` and `UpperCamelCase`.
- **Expected function:**
  ```go
  func CamelToSnakeCase(s string) string
  ```
- **Example:**
  ```
  CamelToSnakeCase("HelloWorld")     → "Hello_World"
  CamelToSnakeCase("camelToSnake")   → "camel_To_Snake"
  CamelToSnakeCase("CAMELtoSnack")   → "CAMELtoSnack"  (unchanged — two capitals)
  CamelToSnakeCase("hey2")           → "hey2"            (unchanged — has digit)
  ```

---

### `wordflip`
- **Type:** Function
- **Instructions:** Reverse the order of words in a `string`.
  - Ignore multiple spaces between words; trim leading/trailing spaces.
  - If the string is empty or only spaces, return `"Invalid Output"`.
- **Expected function:**
  ```go
  func WordFlip(str string) string
  ```
- **Example:**
  ```
  WordFlip("First second last")   → "last second First\n"
  WordFlip("")                    → "Invalid Output\n"
  WordFlip(" hello  all  you! ")  → "you! all hello\n"
  ```

---

### `countalpha` *(see Tier 1)*

---

## 🟠 Tier 3 — Programs (Programs that take `os.Args`)

### `cleanstr`
- **Type:** Program (main)
- **Instructions:** Takes one `string` argument and prints it with:
  - Exactly **one space** between each word.
  - No leading/trailing spaces or tabs.
  - Followed by a newline.
  - If args ≠ 1 or there are no words → print only a newline.
- **Example:**
  ```console
  $ go run . "  only    it's  harder   "
  only it's harder
  $ go run . "" 
  
  ```

---

### `expandstr`
- **Type:** Program (main)
- **Instructions:** Takes one `string` argument and prints it with exactly **three spaces** between words, no leading/trailing whitespace.
  - A "word" = any sequence of visible (non-whitespace) characters.
  - If args ≠ 1 or no words → print nothing.
- **Example:**
  ```console
  $ go run . "  only  it's harder   "
  only   it's   harder
  ```

---

### `rostring`
- **Type:** Program (main)
- **Instructions:** Takes one `string` and rotates it **one word to the left** (the first word goes to the end). Words are separated by one space in the output.
  - A word = sequence of alphanumeric characters.
  - If args ≠ 1 → print only a newline.
- **Example:**
  ```console
  $ go run . "Let there be light"
  there be light Let
  $ go run . "     AkjhZ zLKIJz , 23y"
  zLKIJz , 23y AkjhZ
  ```

---

### `inter`
- **Type:** Program (main)
- **Instructions:** Takes two `string`s and prints, **without duplicates**, the characters that appear in **both** strings, in the order they appear in the **first** string. Followed by `\n`.
  - If args ≠ 2 → print nothing.
- **Example:**
  ```console
  $ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
  padinto
  ```

---

### `union`
- **Type:** Program (main)
- **Instructions:** Takes two `string`s and prints, **without duplicates**, the characters that appear in **either** string, in the order they appear on the command line. Followed by `\n`.
  - If args ≠ 2 → print only `\n`.
- **Example:**
  ```console
  $ go run . zpadinton paqefwtdjetyiytjneytjoeyjnejeyj
  zpadintoqefwjy
  ```

---

### `addprimesum`
- **Type:** Program (main)
- **Instructions:** Takes a positive integer argument and prints the **sum of all prime numbers ≤ that integer** followed by `\n`.
  - If args ≠ 1 or argument is not a positive number → print `0\n`.
- **Example:**
  ```console
  $ go run . 5
  10
  $ go run . 7
  17
  $ go run . -2
  0
  ```

---

### `fprime`
- **Type:** Program (main)
- **Instructions:** Takes a positive `int` and displays its **prime factors** in ascending order, separated by `*`, followed by `\n`.
  - If args ≠ 1 or argument is invalid or has no prime factors → print nothing.
- **Example:**
  ```console
  $ go run . 225225
  3*3*5*5*7*11*13
  $ go run . 42
  2*3*7
  $ go run . 1
  (nothing)
  ```

---

### `wdmatch`
- **Type:** Program (main)
- **Instructions:** Takes two `string`s and checks if the **first** can be written using characters from the **second**, in the same order (not necessarily consecutive).
  - If yes → print the first string followed by `\n`.
  - If no → print nothing.
  - If args ≠ 2 → print nothing.
- **Example:**
  ```console
  $ go run . faya fgvvfdxcacpolhyghbreda
  faya
  $ go run . faya fgvvfdxcacpolhyghbred
  (nothing)
  ```

---

### `hiddenp`
- **Type:** Program (main)
- **Instructions:** Takes two `string`s. Checks if `s1` is **hidden** in `s2` — i.e. every character of `s1` appears in `s2` in the same order (not necessarily consecutive).
  - Print `1\n` if hidden, `0\n` if not.
  - Empty `s1` → always `1`.
  - If args ≠ 2 → print nothing.
- **Example:**
  ```console
  $ go run . "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6"
  1
  $ go run . "DD" "DABC"
  0
  ```

---

### `brackets`
- **Type:** Program (main)
- **Instructions:** Takes any number of `string` arguments. For each, if brackets are **correctly paired and ordered** → print `OK\n`, else print `Error\n`.
  - Bracket types: `()`, `[]`, `{}`
  - All other characters are ignored.
  - A string with no brackets → `OK`.
  - No arguments → print nothing.
- **Example:**
  ```console
  $ go run . '(johndoe)'
  OK
  $ go run . '([)]'
  Error
  $ go run . '' '{[(0 + 0)(1 + 1)](3*(-1)){()}}'
  OK
  OK
  ```

---

## 🔴 Tier 4 — Math & Number Functions

### `findprevprime`
- **Type:** Function
- **Instructions:** Returns the first prime number **≤** the given `int`. Returns `0` if none exists.
- **Expected function:**
  ```go
  func FindPrevPrime(nb int) int
  ```
- **Example:**
  ```
  FindPrevPrime(5) → 5
  FindPrevPrime(4) → 3
  ```

---

## 🟣 Tier 5 — Advanced / Algorithmic

### `rpncalc`
- **Type:** Program (main)
- **Instructions:** Takes a `string` containing an equation in **Reverse Polish Notation (RPN)** and evaluates it.
  - Operators: `+`, `-`, `*`, `/`, `%`
  - Operands and operators separated by at least one space.
  - Print result + `\n`. If invalid or not exactly 1 argument → print `Error\n`.
- **Example:**
  ```console
  $ go run . "1 2 * 3 * 4 +"
  10
  $ go run . "1 2 3 4 +"
  Error
  ```

---

## 📋 Full Pool Index (from kinoz01/zone01-Piscine v2/checkpoint)

These are all exercises confirmed in the checkpoint pool. The ones with full descriptions above are the most common exam picks:

| Exercise | Type | Topic |
|---|---|---|
| `addprimesum` | Program | Prime numbers / math |
| `atoi` | Function | String → int conversion |
| `brackets` | Program | Stack / bracket matching |
| `cameltosnakecase` | Function | String case conversion |
| `canjump` | Function | Array/slice logic |
| `checknumber` | Function | Number validation |
| `chunk` | Function | Slice operations |
| `cleanstr` | Program | String whitespace normalization |
| `concatalternate` | Function | Slice/string interleaving |
| `concatslice` | Function | Slice concat |
| `countalpha` | Function | Character counting |
| `expandstr` | Program | String spacing |
| `fifthandskip` | Function | Slice indexing |
| `findpairs` | Function | Pair finding |
| `findprevprime` | Function | Prime numbers |
| `fprime` | Program | Prime factorization |
| `fromto` | Function | Range generation |
| `grouping` | Function | Slice grouping |
| `hiddenp` | Program | Subsequence check |
| `inter` | Program | String intersection |
| `iscapitalized` | Function | String check |
| `itoa` | Function | Int → string conversion |
| `itoabase` | Function | Base conversion |
| `notdecimal` | Function | Number check |
| `only1` / `onlya` / `onlyb` / `onlyf` / `onlyz` | Program | Basic z01 output |
| `printmemory` | Program | Hex/binary output |
| `printrevcomb` | Program | Combinations (descending) |
| `revconcatalternate` | Function | Reverse interleave |
| `reversestrcap` | Function | Capitalize reversed |
| `romannumbers` | Function | Roman numeral conversion |
| `rostring` | Program | Word rotation |
| `rpncalc` | Program | Reverse Polish Notation |
| `saveandmiss` | Function | Slice filtering |
| `slice` | Function | Slice manipulation |
| `thirdtimeisacharm` | Function | Pattern/logic |
| `union` | Program | String union |
| `wdmatch` | Program | Subsequence match |
| `weareunique` | Function | Slice dedup |
| `wordflip` | Function | Word reversal |
| `zipstring` | Function | String zip/interleave |
