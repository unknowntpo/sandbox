package main

import "fmt"

func a() {
	i := 0
	// A deferred function's arguments are evaluated when
	// the defer statement is evaluated.
	// Ref:
	defer fmt.Println(i)
	i++
	return
}

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

// Why return value is 2?
func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	fmt.Println("a:")
	a()
	fmt.Println("b:")
	b()
	fmt.Println("\nc:", c())
}
