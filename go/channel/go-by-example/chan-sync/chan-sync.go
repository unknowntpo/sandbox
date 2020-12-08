package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println("\rdone             ")

	// Tell caller of worker that the job is done
	done <- true
}

func spinner() {
	for {
		for _, r := range []string{".  ", ".. ", "..."} {
			fmt.Printf("\rworking %s", r)
			time.Sleep(200 * time.Millisecond)
		}
	}
}
func main() {
	// Use done channel to indicate the job is done
	done := make(chan bool, 1)
	// gopl 8.1 spinner
	go spinner()
	go worker(done)

	// Blocked until someone send data through channel
	// and then drop the data.
	<-done
}
