package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Case: k * n < len(s) < 2k * n
Given
s = "abcdefg"
a b c d e f g
0 1 2 3 4 5 6
k = 2
2k = 4

expect = "bacdfeg"

Case: k * n <  len(s) < k * 2n
Given
s = "abcdefghi"

a b c d e f g h i j
0 1 2 3 4 5 6 7 8 9
k = 2
2k = 4

expect = "bacdfeghji"



a


*/
// 0 1 2 3 4
// 0   2   4

func reverseString2(s string, k int) string {
	// want: 0 2k 4k 8k
	// for every 2k
	/*
		for i := 0; i < 2*k; i = i * 2 * k {
			// for every k
			for j := 0; j < k; j++ {

			}

		}
	*/
	forEach := func(s string, f func(string)) {
		for i := range s {
			f(s[i : i+1])
		}
	}

	forEach(s, func(c string) {
		fmt.Printf("%s\n", c)
	})

	return ""
}

func TestReverseString2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		k     int
		want  string
	}{
		{"Case: k * n < len(s) < 2k * n", "abcdefg", 2, "bacdfeg"},
		{"Case: k * n <  len(s) < k * 2n", "abcdefghij", 2, "bacdfeghji"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseString2(tt.input, tt.k)
			assert.Equal(t, tt.want, got)
		})
	}

}
