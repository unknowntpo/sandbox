package main

import "testing"

func TestAppendInt(t *testing.T) {
	tests := []struct {
		x    []int
		y    int
		want []int
	}{
		{[]int{}, 0, []int{0}},
		{[]int{0}, 1, []int{0, 1}},
	}

	for _, test := range tests {
		ans := test.want
		out := appendInt(test.x, test.y)
		for i := range out {
			if out[i] != ans[i] {
				t.Errorf("append %d to %v gets %v, expecting %v",
					test.y, test.x, ans, test.want)
			}
		}
	}

}
