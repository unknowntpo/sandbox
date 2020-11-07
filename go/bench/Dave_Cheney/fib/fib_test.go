package fib

import "testing"

func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(10)
	}
}
func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(20)
	}
}
func BenchmarkFib30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(30)
	}
}

func BenchmarkFib40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(40)
	}
}
