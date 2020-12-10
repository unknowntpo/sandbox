package main

import "testing"

func TestFibRegister(t *testing.T) {
	var tests = []struct {
		fibnum int
		method string
		want   int
	}{
		{30, "nr", 832040},
		{30, "tr", 832040},
		{30, "it", 832040},
		{10, "nr", 55},
		{10, "tr", 55},
		{10, "it", 55},
	}
	for _, test := range tests {
		ans := fibRegister(test.fibnum, test.method)
		if ans != test.want {
			t.Errorf("fibRegister(%d, '%s') = %d, want %d\n", test.fibnum, test.method, ans, test.want)
		}
	}
}

var fibNum int = 50

func BenchmarkFibNormal(b *testing.B) {
	//fibRegister(fibNum, "it")
	// package level variable vs. _ ?
	for i := 0; i < b.N; i++ {
		_ = fibNormal(fibNum)
	}
}

func BenchmarkFibRecTail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fibRecTail(fibNum, 0, 1)
	}

}
func BenchmarkFibIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fibIter(fibNum)

	}
}

//fibRegister(i, "nr")
//fibRegister(i, "tr")
