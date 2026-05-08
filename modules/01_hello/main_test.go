package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMainOutput(t *testing.T) {
	// Keep backup of the real stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	// Restore stdout
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old
	out := <-outC

	expected := "Hello, Gopher!"
	if !strings.Contains(out, expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", out, expected)
	}
}
