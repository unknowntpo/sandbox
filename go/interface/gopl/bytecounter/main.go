package main

import "fmt"

type ByteCounter int

// Implement our own Write method, in order to fit the io.Writer interface
func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	// What do we return?
	return len(p), nil
}
func main() {
	var c ByteCounter
	// Implicitly use pointer receiver, ref:
	// [go spec](https://golang.org/ref/spec#Method_sets)
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0              // reset the counter
	var name = "Dolly" // c can take multiple type of input
	fmt.Fprintf(&c, "Hello, %s", name)
	fmt.Println(c)
}
