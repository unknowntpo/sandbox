package divide

import "testing"

func TestDivide(t *testing.T) {
	var tests = []struct {
		dividend, divisor int32
		want              int32
	}{
		{-2147483648, 1, -2147483648},
		{-2147483648, -1, 2147483647},
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
