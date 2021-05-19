package main

import "fmt"

func main() {
	natural := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10; x++ {
			natural <- x
		}
		close(natural)
	}()

	// Squarer
	go func() {
		for {
			x, ok := <-natural
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	// Printer
	for x := range squares {
		fmt.Println(x)
	}
}
