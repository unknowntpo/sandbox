package main

import "fmt"

// facNaive returns result of n!, if n < 0, returns 0.
func facNaive(n int) int {
	if n < 0 {
		return 0
	}
	return facNaiveRec(n)
}

func facNaiveRec(n int) int {
	if n <= 1 {
		return 1
	}
	return n * facNaive(n-1)
}

// facDP returns result of n!, if n < 0, returns 0.
func facDP(n int) int {
	if n < 0 {
		return 0
	}
	return facDPRec(n)
}

var cache = make(map[int]int)

func facDPRec(n int) int {
	if n <= 1 {
		return 1
	}
	if facN, ok := cache[n]; ok {
		return facN
	} else {
		facN = n * facDPRec(n-1)
		cache[n] = facN
		return facN
	}
}
func main() {
	fmt.Println("vim-go")
}
