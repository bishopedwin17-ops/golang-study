package exercises
import ( "bytes"; "io"; "os"; "testing" )
func TestPrintStr(t *testing.T) {
	tests := []string{"Hello", "World", "123", ""}
	for _, tc := range tests {
		old := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		PrintStr(tc)
		w.Close()
		os.Stdout = old
		var buf bytes.Buffer
		io.Copy(&buf, r)
		if buf.String() != tc {
			t.Errorf("PrintStr(%q) printed %q", tc, buf.String())
		}
	}
}
