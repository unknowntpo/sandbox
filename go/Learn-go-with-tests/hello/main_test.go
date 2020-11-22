package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Eric")
	want := "Hello, Eric"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
