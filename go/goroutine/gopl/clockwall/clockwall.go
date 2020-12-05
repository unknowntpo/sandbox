package main

import (
	"io"
	"log"
	"net"
	"os"
)

// Usage: ./clockwall NewYork=localhost:8010 Tokyo=localhost:8020
func main() {
	// TODO: parse multiple location - NewYork -> US/NewYork  Tokyo -> Asia/Tokyo
	port := "localhost:8000"
	conn, err := net.Dial("tcp", port)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		// ? What to do ?
		log.Fatal(err)
	}
}
