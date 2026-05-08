package exercises
import ( "bytes"; "io"; "os"; "testing"; "strings" )
func TestPrintComb(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	PrintComb()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	out := buf.String()
	if !strings.HasPrefix(out, "012, 013") || !strings.HasSuffix(out, "789\n") || len(out) != 480 {
		t.Errorf("PrintComb output incorrect. Length: %d", len(out))
	}
}
