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
	if len(os.Args) != 2 {
		// TODO:
		fmt.Fprintf(os.Stderr, "Usage: ./clock <port>\n")
		os.Exit(1)
	}
	// TODO: validate the port specified from user
	//TODO: parse time zone
	listener, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for {
		// accept the connection
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err) // e.g. connection aborted
			break
		}
		go handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// send time in format
		_, err := io.WriteString(c, time.Now().Format("Mon Jan 2 15:04:05 -0700 MST 2006\n"))
		if err != nil {
			return // e.g. client disconnected
		}
		// delay for 1 seconds
		time.Sleep(1 * time.Second)
	}
}
