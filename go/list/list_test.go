package main

import "testing"

func TestPushBack(t *testing.T) {
	want := 1
	l := New()
	l.PushBack(4)
	l.PushBack(3)
	l.PushBack(2)
	l.PushBack(1)
	got := l.Back().Value

	if got != want {
		t.Errorf("Wrong element at tail of list, got %v, want %v", got, want)
	}
}
