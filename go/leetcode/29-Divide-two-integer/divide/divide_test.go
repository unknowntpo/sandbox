package divide

import "testing"

func TestDivide(t *testing.T) {
	var tests = []struct {
		dividend, divisor int32
		want              int32
	}{
		{-2147483648, 1, -2147483648},
		{-2147483648, -1, 2147483647},
		{-2147483648, 2, -1073741824},
		{10, 3, 3},
		{7, -3, -2},
		{0, 1, 0},
		{1, 1, 1},
	}
	for _, test := range tests {
		if ans := divide(test.dividend, test.divisor); ans != test.want {
			t.Errorf("divide(%d, %d) = %d; want %d", test.dividend, test.divisor, ans, test.want)
		}
	}
}
func TestNormalDivide(t *testing.T) {
	var tests = []struct {
		dividend, divisor int32
		want              int32
	}{
		{-2147483648, 1, -2147483648},
		{-2147483648, -1, 2147483647},
		{-2147483648, 2, -1073741824},
		{10, 3, 3},
		{7, -3, -2},
		{0, 1, 0},
		{1, 1, 1},
	}
	for _, test := range tests {
		if ans := normalDivide(test.dividend, test.divisor); ans != test.want {
			t.Errorf("divide(%d, %d) = %d; want %d", test.dividend, test.divisor, ans, test.want)
		}
	}
}
func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = divide(-2147483648, 1)
	}
}

func BenchmarkNormalDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = normalDivide(-2147483648, 1)
	}
}
