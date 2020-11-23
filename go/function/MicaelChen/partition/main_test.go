package main

import (
	"testing"
)

func checkeq(t *testing.T, m []int, n []int) bool {
	t.Helper()
	if len(m) != len(n) {
		return false
	}

	for i := 0; i < len(m); i++ {
		if m[i] != n[i] {
			return false
		}
	}
	return true
}
func TestPartition(t *testing.T) {

	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	wantFit := []int{1, 3, 5, 7, 9}
	wantNotFit := []int{0, 2, 4, 6, 8}

	// Test for partition of odd, even number
	gotFit, gotNotFit := Partition(original, func(elem int) bool {
		return (elem & 1) == 1
	})

	if !checkeq(t, gotFit, wantFit) {
		t.Errorf("Unexpected slice: Fit, got %v, want %v\n", gotFit, wantFit)
	}
	if !checkeq(t, gotNotFit, wantNotFit) {
		t.Errorf("Unexpected slice: NotFit , got %v, want %v\n", gotNotFit, wantNotFit)
	}

}
