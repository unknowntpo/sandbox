package main

import (
	"fmt"
	"time"
)

func main() {
	// Why does infinite loop stop ?
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacii (%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		// back quote v.s single quote double quote?
		for _, r := range `-/|\` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
