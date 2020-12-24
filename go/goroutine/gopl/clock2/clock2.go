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

	// validate the location specified by $TZ
	_, err := time.LoadLocation(os.Getenv("TZ"))
	if err != nil {
		log.Fatal(err)
	}

	// listen to the conn
	l, err := net.Listen("tcp", ":"+os.Args[2])
	if err != nil {
		log.Fatalf("fail to listen to port %s", os.Args[2])
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		timezone := os.Getenv("TZ")
		t, err := timeIn(time.Now(), timezone)
		if err != nil {
			// fail to resolve TZ
			log.Print(err)
			return
		}
		_, err = io.WriteString(c, timezone+" : "+t.Format("Mon Jan 2 15:04:05 -0700 MST 2006")+"\n")
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func timeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}
