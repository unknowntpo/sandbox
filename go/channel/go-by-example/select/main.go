package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	fmt.Println("Start two goroutines")
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "goroutine:1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "goroutine:2"
	}()

	// What does for-loop do?
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Got message from", msg1)
		case msg2 := <-c2:
			fmt.Println("Got message from", msg2)
		}
	}

}
