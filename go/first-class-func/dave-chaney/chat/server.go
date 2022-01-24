package main

import (
	"io"
	"log"
	"net"
)

type Mux struct {
	add     chan net.Conn
	remove  chan net.Addr
	sendMsg chan string
	conns   map[net.Addr]net.Conn
}

func (m *Mux) Add(conn net.Conn) {
	m.add <- conn
}

func (m *Mux) Remove(addr net.Addr) {
	m.remove <- addr
}

func (m *Mux) SendMsg(msg string) error {
	m.sendMsg <- msg
	return nil
}

func (m *Mux) loop() error {
	for {
		select {
		case conn := <-m.add:
			log.Println("loop got a connection", conn.RemoteAddr())
			m.conns[conn.RemoteAddr()] = conn
		case addr := <-m.remove:
			delete(m.conns, addr)
		case msg := <-m.sendMsg:
			log.Println("got a message", msg)
			for _, conn := range m.conns {
				io.WriteString(conn, msg)
			}
		}
	}
}

const Addr string = ":8080"

func main() {
	m := &Mux{
		add:     make(chan net.Conn, 3),
		remove:  make(chan net.Addr, 3),
		sendMsg: make(chan string, 3),
		conns:   make(map[net.Addr]net.Conn),
	}

	ln, err := net.Listen("tcp", Addr)
	if err != nil {
		log.Printf("failed to listen to port %s", Addr)
	}

	// Listener goroutine
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Printf("failed to accept connection: %v", err)
				continue
			}
			log.Println("receive a connection from ", conn.RemoteAddr())
			m.add <- conn
		}
	}()

	log.Println("chat server is running at ", Addr)

	m.loop()
}
