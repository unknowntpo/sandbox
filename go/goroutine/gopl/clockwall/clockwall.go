package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type clock struct {
	location, hostAddr string
}

// Usage: ./clockwall NewYork=localhost:8010 Tokyo=localhost:8020
func main() {
	// TODO: parse multiple location - NewYork -> US/NewYork  Tokyo -> Asia/Tokyo
	//ports := []string{"localhost:8000", "localhost:8010", "localhost:8020"}
	clocks := parseArgs()
	for _, clock := range clocks {
		// TODO: Display location of corresponding clock
		_ = clock.location
		port := clock.hostAddr
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

func parseArgs() []clock {
	// loop through os.Args[1:]
	var clocks []clock
	for _, s := range os.Args[1:] {
		clockSlice := strings.Split(s, "=")
		clocks = append(clocks, clock{
			location: clockSlice[0],
			hostAddr: clockSlice[1],
		})
	}
	return clocks
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		// ? What to do ?
		log.Fatal("Stop copy:", err)
	}
}
