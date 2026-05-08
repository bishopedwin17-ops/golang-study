# Instructions: Functions

## Objective

Open `functions.go` and implement the four functions:

1. **`Divide(a, b float64) (float64, error)`**  
   Return `a/b, nil` normally.  
   Return `0, fmt.Errorf("cannot divide by zero")` when `b == 0`.

2. **`SumAll(nums ...int) int`**  
   Accept variadic int args. Return the total.

3. **`MakeAdder(n int) func(int) int`**  
   Return a closure that adds `n` to its argument.  
   e.g. `MakeAdder(5)(3)` → `8`

4. **`MinMax(nums []int) (min, max int)`**  
   Use named return values. Find the min and max in one pass.  
   Use a `naked return` at the end.

## How to Verify
```
cd modules/05_functions
go test -v
```
