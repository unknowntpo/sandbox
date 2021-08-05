package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	serve()
}

func serve() {
	shutdown := make(chan struct{})
	go graceful(shutdown)
	fmt.Println("Starting server")
	for {
		select {
		// Means server has already shutted down by graceful().
		case <-shutdown:
			fmt.Println("\nstopped server")
			return
		default:
		}
	}
}

func graceful(shutdown chan struct{}) {
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT)

	// Signal incoming.
	<-quit

	// Simulate resource cleanup of server.
	shutdownProcess := func() {
		time.Sleep(2 * time.Second)
	}
	// spinner
	go spinner("Shutting down server...")

	shutdownProcess()

	shutdown <- struct{}{}
	return
}

func spinner(work string) {
	for {
		for _, s := range []string{"⢹", "⢺", "⢼", "⣸", "⣇", "⡧", "⡗", "⡏"} {
			fmt.Printf("\r%s %s", s, work)
			time.Sleep(30 * time.Millisecond)
		}
	}
}
