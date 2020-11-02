package main

import "fmt"

// Function fib is a function that returns fibonacci sequence
// fib(0) == 0
// fib(1) == 1
// fib(2) == 1
func fib_normal(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return fib_normal(n-1) + fib_normal(n-2)
}

// Tail recursion of counting fibonacii sequence
// n: current fib number
// pre: previous fib number
// nxt: next fib number
func fib_rec_tail(n, pre, cur int) int {
	if n == 0 {
		return pre
	}
	return fib_rec_tail(n-1, cur, pre+cur)
}

func fibRegister(n int) int {
	if false { /* Choose fib normal */
		return fib_normal(n)
	}
	// TODO: implement fib recur tail
	if true {
		// Pre-check of fib
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		return fib_rec_tail(n, 0, 1)
	}
	return n
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fibRegister(i))
	}
}
