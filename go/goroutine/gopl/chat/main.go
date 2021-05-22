package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type clientChan chan string

var (
	messages = make(chan string)
	quit     = make(chan string)
)

func broadcaster(clients map[string]clientChan) {
	for {
		select {
		case msg := <-messages:
			for _, clientChan := range clients {
				clientChan <- msg
			}
		case who := <-quit:
			delete(clients, who)
			log.Printf("client %s is leaving...", who)
			for _, clientChan := range clients {
				clientChan <- fmt.Sprintf("%s: Good bye !", who)
			}
		}
	}
}

func serve() {
	clients := make(map[string]clientChan)
	go broadcaster(clients)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Register client
		// addr: conn.RemoteAddr().String()
		addr := conn.RemoteAddr().String()
		clients[addr] = make(clientChan)
		log.Printf("Client from %s joined\n", addr)

		go handleConn(conn, clients[addr])
	}

}

func handleConn(c net.Conn, ch clientChan) {
	defer c.Close()
	// Listen to connection from client and write message from client to broadcaster
	go func() {
		scanner := bufio.NewScanner(c)
		for scanner.Scan() {
			// got msg from c
			msg := scanner.Text()
			log.Println(msg)
			messages <- msg
		}
		quit <- c.RemoteAddr().String()
	}()
	// Listen to broadcast message from broadcaster, and write it back to client.
	for msg := range ch {
		io.WriteString(c, msg+"\n")
	}
}

func main() {
	serve()
}
