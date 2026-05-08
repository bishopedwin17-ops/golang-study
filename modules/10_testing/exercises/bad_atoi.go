package exercises

func BadAtoi(s string) int {
	if s == "" { return 1 } // Bug 1
	return 0 // Bug 2
}
