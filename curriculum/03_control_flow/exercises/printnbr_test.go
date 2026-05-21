package exercises
import ( "bytes"; "io"; "os"; "testing"; "strconv" )
func TestPrintNbr(t *testing.T) {
	tests := []int{0, 1, -1, 123, -123, 2147483647, -2147483648}
	for _, tc := range tests {
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		PrintNbr(tc)
		w.Close()
		os.Stdout = old
		var buf bytes.Buffer
		io.Copy(&buf, r)
		expected := strconv.Itoa(tc)
		if buf.String() != expected {
			t.Errorf("PrintNbr(%d) printed %q, expected %q", tc, buf.String(), expected)
		}
	}
}
