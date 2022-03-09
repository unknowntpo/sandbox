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
	fmt.Printf("type of s[0]: %T", s[0])
	forEach := func(s string, f func(byte)) {
		for _, b := range []byte(s) {
			f(b)
		}
	}
	forEach(s, func(b byte) {
		fmt.Printf("%s ", string(b))
	})

	fmt.Println("\n********for each k")
	forEachK := func(s string, k int, f func(byte)) {
		for i := 0; i < len(s); i = i + k {
			fmt.Printf("%s ", string(s[i]))
		}
	}

	forEachK(s, k, func(b byte) {
		fmt.Printf("%s ", string(b))
	})

	fmt.Println("\n*******for each 2k")
	forEach2K := func(s string, k int, f func(byte)) {
		for i := 0; i < len(s); i = i + 2*k {
			fmt.Printf("%s ", string(s[i]))
		}
	}

	forEach2K(s, k, func(b byte) {
		fmt.Printf("%s ", string(b))
	})

	// TODO: What does this means ?
	forEach2K(s, k, func(b byte) {
		forEachK(s, k, func(b byte) {
			fmt.Printf("%s ", string(b))
		})
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
