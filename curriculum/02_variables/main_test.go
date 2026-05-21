package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestVariableOutput(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC

	if !strings.Contains(out, "GoStudy") {
		t.Errorf("Missing AppName constant output. Expected 'GoStudy', got: %s", out)
	}
	if !strings.Contains(out, "1.0") {
		t.Errorf("Missing version variable output. Expected '1.0', got: %s", out)
	}
	if !strings.Contains(out, "true") {
		t.Errorf("Missing isAwesome variable output. Expected 'true', got: %s", out)
	}
}
