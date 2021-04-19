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

func TestPushFront(t *testing.T) {
	want := 1
	l := New()
	l.PushFront(4)
	l.PushFront(3)
	l.PushFront(2)
	l.PushFront(1)

	got := l.Front().Value

	if got != want {
		t.Errorf("Wrong element at head of list, got %v, want %v", got, want)
	}

}
