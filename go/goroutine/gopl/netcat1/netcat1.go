package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// connect to the addr on the named network
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

// What does mustCopy() do?
// mustCopy() do io.Copy() from an io.Reader (src) to io.Writer (dst),
// and handled the error, the logic of 'must' is found in template.Must()
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
