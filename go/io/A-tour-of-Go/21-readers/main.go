package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("At: %v, %s", e.When, e.What)
}
func main() {
	err := readFromReader()
	fmt.Println(err)
}

// readFromReader creates an Reader and read from it to buffer,
// if EOF, return my custom error structure to indicate when and
// what happend.
func readFromReader() error {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)

		fmt.Printf("n = %v, err = %v\t", n, err)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			return &MyError{
				When: time.Now(),
				What: "encounter EOF",
			}
		}
	}
}
