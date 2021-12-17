package main

import "testing"

func BenchmarkFalsesharing(b *testing.B) {
	// prepare slice
	veryBigSlice := make([]int, 10000)
	for i := range veryBigSlice {
		veryBigSlice[i] = i
	}
	b.Run("false sharing", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			false_sharing(veryBigSlice)
		}
	})
	b.Run("not false sharing", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			not_false_sharing(veryBigSlice)
		}
	})
}
