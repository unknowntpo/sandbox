// The example is from pkg io documentation
// https://golang.org/pkg/io/#LimitReader
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type limitedReader struct {
	R io.Reader // underlying reader
	N int64     // n bytes remained
}

// Our implementation of io.LimitReader,
// It read n bytes from an io.Reader r, and return another io.Reader
// that reads from r but reports an end-of-file condition after n
// bytes.
func limitReader(r io.Reader, n int64) io.Reader {
	return r
}
func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 4)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}

}
