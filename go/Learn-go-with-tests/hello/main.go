package main

import "fmt"

const engHelloPrefix = "Hello, "

func Hello(name string) string {
	return engHelloPrefix + name
}
func main() {
	fmt.Println(Hello("world"))
}
