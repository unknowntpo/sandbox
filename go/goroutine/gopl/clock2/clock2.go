package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// TODO: Use flags package to handle argument list
	if len(os.Args) != 3 {
		// TODO:
		fmt.Fprintf(os.Stderr, "Usage: TZ=<timezone> ./clock2 -port <port>\n")
		os.Exit(1)
	}

	listener, err := net.Listen("tcp", ":"+os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	for {
		// accept the connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err) // e.g. connection aborted
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// send time in format
		timezone := os.Getenv("TZ")
		_, err := io.WriteString(c, timezone+": "+time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006\n"))
		if err != nil {
			return // e.g. client disconnected
		}
		// delay for 1 seconds
		time.Sleep(1 * time.Second)
	}
}
