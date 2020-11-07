package main

import "testing"

func BenchmarkPrintInt2String01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printInt2String01(100)
	}
}

func BenchmarkPrintInt2String02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printInt2String02(int64(100))
	}
}

func BenchmarkPrintInt2String03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printInt2String03(100)
	}
}
