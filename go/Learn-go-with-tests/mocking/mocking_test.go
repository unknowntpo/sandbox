package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	want := "3"

	buf := &bytes.Buffer{}
	Countdown(buf)
	got := buf.String()

	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
