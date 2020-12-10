package main

import "fmt"

func main() {
	msgs := make(chan string, 2)

	msgs <- "Ping"
	msgs <- "Pong"
	out1 := <-msgs
	out2 := <-msgs
	fmt.Println(out1)
	fmt.Println(out2)
}
