package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"test for even target with nums contains target / 2",
			[]int{3, 2, 4},
			6, []int{1, 2},
		},
		{"test for even target with nums contains two element == target / 2",
			[]int{3, 3},
			6, []int{0, 1},
		},
	}

	for _, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		if !sameWithoutOrder(got, tt.want) {
			t.Errorf("Wrong answer, got %v, want %v", got, tt.want)
		}

	}
}

// sameWithoutOrder check whether a, b has same element, the order of element doesn't need to be the same.
func sameWithoutOrder(a, b []int) bool {
	t.Helper()
	if len(a) != len(b) {
		return false
	}
	for _, va := range a {
		isContain := false
		for _, vb := range b {
			if va == vb {
				isContain = true
				break
			}
		}
		if isContain == false {
			return false
		}
	}
	return true
}
