package power

import "testing"

func TestPower(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},  // test for base == 0
		{10, 0, 1}, // test for power == 0
		{2, 3, 8},
		{2, 4, 16},
		{3, 4, 81},
	}
	t.Run("Iteration version of power", func(t *testing.T) {
		for _, test := range tests {
			if ans := powerIter(test.a, test.b); ans != test.want {
				t.Errorf("%d ^ %d = %d, want %d", test.a, test.b, ans, test.want)
			}
		}
	})

	t.Run("Recursion version of power", func(t *testing.T) {
		for _, test := range tests {
			if ans := powerRec(test.a, test.b); ans != test.want {
				t.Errorf("%d ^ %d = %d, want %d", test.a, test.b, ans, test.want)
			}
		}
	})
}
func BenchmarkPowerIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = powerIter(10, 1000)
	}
}

func BenchmarkPowerRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = powerRec(10, 1000)
	}
}
