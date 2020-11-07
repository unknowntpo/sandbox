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
