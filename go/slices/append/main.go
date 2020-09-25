// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 88.

// Append illustrates the behavior of the built-in append function.
package main

import "fmt"

func appendslice(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to expand the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

//!+append
func appendInt(x []int, y int) []int {
	// if cap(x) >= len(x) + 1
	// just let x[len(x)] == y
	// can we append len(x) directly? <
	// else ,
	// means cap(x) < len(x) + 1
	// allocate new slices with length cap(x) * 2
	// copy from old slice x to new slice z

	if cap(x) >= len(x)+1 {
		z := x[:len(x)+1]
		z[len(x)] = y
		return z
	}
	// if we reach here, mean no enough space for new element
	zlen := cap(x) + 1
	zcap := len(x) + 1
        // prevent 2 * 0
        if zcap < 2*len(x) {
            zcap = 2*len(x)
        }
        var z []int // how to declare a nil slice?
	z = make([]int, zlen, zcap)
	copy(z, x)
        z[len(x)] = y
	return z
}

//!-append

//!+growth
func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

//!-growth

/*
//!+output
0  cap=1   [0]
1  cap=2   [0 1]
2  cap=4   [0 1 2]
3  cap=4   [0 1 2 3]
4  cap=8   [0 1 2 3 4]
5  cap=8   [0 1 2 3 4 5]
6  cap=8   [0 1 2 3 4 5 6]
7  cap=8   [0 1 2 3 4 5 6 7]
8  cap=16  [0 1 2 3 4 5 6 7 8]
9  cap=16  [0 1 2 3 4 5 6 7 8 9]
//!-output
*/
