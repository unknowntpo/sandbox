package main

import "fmt"

// Function fib is a function that returns fibonacci sequence
// fib(0) == 0
// fib(1) == 1
// fib(2) == 1
func fibNormal(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return fibNormal(n-1) + fibNormal(n-2)
}

// Tail recursion of counting fibonacii sequence
// n: current fib number
// pre: previous fib number
// nxt: next fib number
func fibRecTail(n, pre, cur int) int {
	if n == 0 {
		return pre
	}
	return fibRecTail(n-1, cur, pre+cur)
}

func fibIter(n int) int {
	pre := 1
	cur := 1
	result := 1
	for i := n; i > 2; i-- {
		result = pre + cur
		pre = cur
		cur = result
	}
	return result
}
func fibRegister(n int, method string) int {
	if n == 0 {
		return 0
	}
	switch method {
	case "nr":
		// Normal recursion
		return fibNormal(n)
	case "tr":
		// Tail recursion
		// Pre-check of tail recursion
		if n == 1 {
			return 1
		}
		return fibRecTail(n, 0, 1)
	case "it":
		// Iteration
		return fibIter(n)
	}
	return n

}

func main() {
	fibNum := 50
	// Iteration version of fib counting function
	fmt.Println("Iteration:")
	for i := 0; i < fibNum; i++ {
		fmt.Print(fibRegister(i, "it"), " ")
	}
	fmt.Println()

	// Normal recursion fib counting function
	fmt.Println("Normal recursion")
	for i := 0; i < fibNum; i++ {
		fmt.Print(fibRegister(i, "nr"), " ")
	}
	fmt.Println()

	// Tail recursion fib counting function
	fmt.Println("Tail recursion")
	for i := 0; i < fibNum; i++ {
		fmt.Print(fibRegister(i, "tr"), " ")
	}
	fmt.Println()
}
