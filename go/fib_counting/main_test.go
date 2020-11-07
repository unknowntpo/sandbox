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
	result := fibNormal(fibNum)
	b.Logf("fib normal: fib[%d] = %d", fibNum, result)
}

func BenchmarkFibRecTail(b *testing.B) {
	result := fibRecTail(fibNum, 0, 1)
	b.Logf("fib rec tail: fib[%d] = %d", fibNum, result)
}
func BenchmarkFibIter(b *testing.B) {
	result := fibIter(fibNum)
	b.Logf("fib rec tail: fib[%d] = %d", fibNum, result)
}

//fibRegister(i, "nr")
//fibRegister(i, "tr")
