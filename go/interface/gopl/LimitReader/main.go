// The example is from pkg io documentation
// https://golang.org/pkg/io/#LimitReader
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

// A LimitedReader reads from R but limits the amount of data returned to just
// N bytes. Each call to Read updates N to reflect the new amount remaining.
// Read returns EOF when N <= 0 or when the underlying R returns EOF.
type limitedReader struct {
	R io.Reader // underlying reader
	N int64     // n bytes remained
}

// Read reads from l.R, but limits the amount of data returned to just
// N bytes. Each calll to Read updates N to reflect the new amount remaining.
// Read returns EOF when N <= 0 or when the underlying R returns EOF.
func (l *limitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		err = io.EOF
		return
	}

	// Limit the amount of bytes we read from l.R to l.N
	// Ref: https://golang.org/src/io/io.go?s=15184:15241#L438
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}

	// read n bytes from underlying io.Reader l.R
	n, err = l.R.Read(p)
	if err != nil {
		return 0, err
	}
	// l.N += n
	l.N -= int64(n)
	return n, nil
}

// LimitReader returns a io.Reader that reads from r but stops with EOF after n
// bytes. The underlying implementation is a *LimitedReader.
func limitReader(r io.Reader, n int64) io.Reader {
	// copy elements in r to *limitedReader l
	l := &limitedReader{r, n}
	return l
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := limitReader(r, 4)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}

}
