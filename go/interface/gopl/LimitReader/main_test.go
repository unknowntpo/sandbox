package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	b := &bytes.Buffer{}
	r := strings.NewReader("1234567890")

	lr := limitReader(r, 4)
	// use limitReader() to convert it to limitedReader
	// read from it
	n, _ := io.Copy(b, lr)
	// verify bytes in buffer
	got := b.String()
	want := "1234"
	if n != int64(4) {
		t.Errorf("wrong number of bytes to read, got %d, want %d", n, 4)
	}
	if got != want {
		t.Errorf("content of bytes does not match, got %s, want %s", got, want)
	}
}
