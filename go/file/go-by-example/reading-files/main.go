package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	// Read files and print the content
	dat, err := ioutil.ReadFile("dat")
	check(err)
	fmt.Println(dat)

	// Open the file using os.Open())
	f, err := os.Open("dat")
	check(err)

	// Read first 5 bytes and print the data
	fmt.Println("Reading 5 btyes ...")
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, b1[:n1])

	// Seek to known location of file
	// pos 5 is new-line char
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// Seek to another known location of file
	// Practice using io.ReadAtLeast
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	// Read at least 2 bytes for f to b3
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// Rewind back to first position of file
	_, err = f.Seek(0, 0)
	check(err)

	// Test reading with bufio package
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close()
}
