package exercises
import ( "bytes"; "io"; "os"; "testing"; "strings" )
func TestPrintComb2(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	PrintComb2()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	out := buf.String()
	if !strings.HasPrefix(out, "00 01, 00 02") || !strings.HasSuffix(out, "98 99\n") || len(out) != 34200 {
		t.Errorf("PrintComb2 output incorrect. Length: %d", len(out))
	}
}
