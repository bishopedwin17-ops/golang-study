package exercises
import ( "bytes"; "io"; "os"; "testing" )
func TestHelloPiscine(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	HelloPiscine()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	if buf.String() != "Hello Piscine!\n" {
		t.Errorf("Expected 'Hello Piscine!\\n', got %q", buf.String())
	}
}
