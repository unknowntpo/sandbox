package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	want := `3
2
1
Go!`

	buf := &bytes.Buffer{}
	spySleeper := &SpySleeper{}
	Countdown(buf, spySleeper)
	got := buf.String()

	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
