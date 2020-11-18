package sort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var tests = []struct {
		in   []int
		want []int
	}{
		{nil, nil},         // test for nil slice
		{[]int{}, []int{}}, // test for empty slice
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},
		{[]int{5, 2, 2, 3, 1}, []int{1, 2, 2, 3, 5}},
	}
	for _, test := range tests {
		if ans := bubbleSort(test.in); !sort.IntsAreSorted(ans) {
			t.Errorf("Slice is not sorted, get %v, want %v\n", test.want, ans)
		}
	}
}

func BenchmarkBubbleSort5Ele(b *testing.B) {
	var in = []int{5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		_ = bubbleSort(in)
	}
}

func BenchmarkBubbleSort1000Ele(b *testing.B) {
	var in []int
	var size int

	size = 1000
	in = make([]int, 1000)
	for i := 0; i < size; i++ {
		in[i] = rand.Int() % 50
	}
	// reset timer before actual bench
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = bubbleSort(in)
	}
}
