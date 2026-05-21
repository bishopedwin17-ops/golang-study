# Instructions: Control Flow

## Objective
Complete the `classify.go` file to implement the `Classify` function.

### Rules:
- Input is an `int` called `n`.
- Return `"negative"` if `n < 0`.
- Return `"zero"` if `n == 0`.
- Return `"small"` if `1 <= n <= 9`.
- Return `"large"` if `n >= 10`.

Also implement `DayType(day string) string`:
- Return `"weekday"` for Monday–Friday.
- Return `"weekend"` for Saturday or Sunday.
- Return `"unknown"` for anything else.

## How to Verify
```
cd modules/03_control_flow
go test -v
```

---

## 🐲 BOSS FIGHT: Piscine Challenges

These exercises are from the 01-edu Piscine. They are strict. 
**Rules:** You can ONLY use the `github.com/01-edu/z01` package to print (specifically `z01.PrintRune`). You cannot use `fmt` or any other standard library package for these functions.

### 1. `printcomb`
Write a function that prints, in ascending order and on a single line: all **unique** combinations of three different digits so that, the first digit is lower than the second, and the second is lower than the third.
These combinations are separated by a comma and a space.
- `000` or `999` are not valid (digits must be different).
- `987` should not be shown (first digit must be less than second, etc).
- Output example: `012, 013, 014, 015, 016, 017, 018, 019, 023, ..., 689, 789` (no newline at the end).

### 2. `printcomb2`
Write a function that prints in ascending order and on a single line: all possible combinations of two different two-digit numbers.
These combinations are separated by a comma and a space.
- Output example: `00 01, 00 02, 00 03, ..., 00 98, 00 99, 01 02, 01 03, ..., 97 98, 97 99, 98 99` (no newline at the end).

Open `piscine.go` to implement them, and test them with `go test -v -run Piscine`.
