package main

import "fmt"

func ping(pings chan<- string, msg string) {
	// Send msg to pings
	pings <- msg
	// This will cause compile-time error
	//<-pings
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
	// Send message to receive-only channel will cause error
	//	pings <- "a"
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Hello")
	pong(pings, pongs)
	fmt.Println("messsage received from <-pongs: ", <-pongs)
}
