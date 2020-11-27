package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	// Start a new goroutine with func literal
	go func(msg string) {
		fmt.Println(msg)
	}("going") // passing in args to anonymous func

	time.Sleep(time.Second)
	fmt.Println("done")
}
