package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

// Usage: ./clockwall NewYork=localhost:8010 Tokyo=localhost:8020
func main() {
	// TODO: parse multiple location - NewYork -> US/NewYork  Tokyo -> Asia/Tokyo
	ports := []string{"localhost:8000", "localhost:8010", "localhost:8020"}
	for _, port := range ports {
		conn, err := net.Dial("tcp", port)
		if err != nil {
			log.Fatal("conn Failed", err)
			continue
		}
		defer conn.Close()
		go mustCopy(os.Stdout, conn)
	}
	// Wait for users send SIGINT to stop the program
	// TODO: How to wait for all goroutines to stop?
	for {
		time.Sleep(time.Second)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		// ? What to do ?
		log.Fatal("Stop copy:", err)
	}
}
