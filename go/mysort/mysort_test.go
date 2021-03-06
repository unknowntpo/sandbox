package mysort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

var testSort = []struct {
	name string
	in   []int
	want []int
}{
	{"nil slice", nil, nil},           // test for nil slice
	{"empty slice", []int{}, []int{}}, // test for empty slice
	{"normal slice", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
	{"zero value slice", []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},
	{"normal slice - 2", []int{5, 2, 2, 3, 1}, []int{1, 2, 2, 3, 5}},
}

func TestBubbleSort(t *testing.T) {
	for _, tt := range testSort {
		t.Run(tt.name, func(t *testing.T) {
			if ans := bubbleSort(tt.in); !(sort.IntsAreSorted(ans) && reflect.DeepEqual(ans, tt.want)) {
				t.Errorf("Slice is not sorted or answer not correct, get %v, want %v\n", ans, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tt := range testSort {
		t.Run(tt.name, func(t *testing.T) {
			if ans := mergeSort(tt.in); !(sort.IntsAreSorted(ans) && reflect.DeepEqual(ans, tt.want)) {
				t.Errorf("Slice is not sorted or answer not correct, get %v, want %v\n", ans, tt.want)
			}
		})
	}

}

// Test for merging two sorted slice
func TestMerge(t *testing.T) {
	var tests = []struct {
		name        string
		left, right []int
		want        []int
	}{
		{name: "nil slice", left: nil, right: nil, want: nil},
		{name: "empty slice", left: []int{}, right: []int{}, want: []int{}},
		{name: "normal sorted slice", left: []int{1, 2, 3}, right: []int{3, 4, 5}, want: []int{1, 2, 3, 3, 4, 5}},
		{name: "zero value slice", left: []int{0, 0, 0}, right: []int{0, 0}, want: []int{0, 0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ans := merge(tt.left, tt.right); !reflect.DeepEqual(ans, tt.want) {
				t.Errorf("Wrong answer, get %v, want %v\n", ans, tt.want)
			}
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	b.Run("Bench slice with 5 elements", func(b *testing.B) {
		var in = []int{5, 4, 3, 2, 1}
		for i := 0; i < b.N; i++ {
			_ = bubbleSort(in)
		}
	})

	b.Run("Bench slice with 1000 elements", func(b *testing.B) {
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
	})
}

func BenchmarkMergeSort(b *testing.B) {
	b.Run("Bench slice with 5 elements", func(b *testing.B) {
		var in = []int{5, 4, 3, 2, 1}
		for i := 0; i < b.N; i++ {
			_ = mergeSort(in)
		}
	})

	b.Run("Bench slice with 1000 elements", func(b *testing.B) {
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
			_ = mergeSort(in)
		}
	})
}

func BenchmarkMerge(b *testing.B) {
	b.Run("Bench slice with 5 elements", func(b *testing.B) {
		var left = []int{1, 2, 3, 4, 5}
		var right = []int{1, 2, 3, 4, 5}
		for i := 0; i < b.N; i++ {
			_ = merge(left, right)
		}
	})

	b.Run("Bench slice with 1000 elements", func(b *testing.B) {
		var size int

		size = 1000
		left := make([]int, 1000)
		right := make([]int, 1000)
		for i := 0; i < size; i++ {
			left[i] = i
			right[i] = i + 1
		}
		// reset timer before actual bench
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = merge(left, right)
		}
	})
}
