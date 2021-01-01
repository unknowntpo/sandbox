package main

import (
	"fmt"
	"io"
	"os"
)

const (
	MAXSIZE = 1024
)

// cat read the file named name and output its content
// to out.
func cat(name string, out io.Writer) {
	b := make([]byte, MAXSIZE)
	// open file
	f, err := os.Open(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail to cat file: %v", err)
		os.Exit(1)
	}

	for {
		_, err := f.Read(b)
		if err == io.EOF {
			break
		}

		_, err = out.Write(b)
		if err != nil {
			// What might be the error durint Write() ?
			// ErrShortWrite?
			break
		}
	}
}

func main() {
	cat("out.txt", os.Stdout)
}
