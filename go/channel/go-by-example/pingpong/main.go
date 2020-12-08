package main

import "fmt"

func main() {
	msg := make(chan string)

	go func() { msg <- "Bomb" }()

	out := <-msg
	fmt.Println(out)
}
