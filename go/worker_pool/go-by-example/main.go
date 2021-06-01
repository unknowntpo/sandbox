package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, ID int, jobqueue <-chan int, result chan<- int) {
	for n := range jobqueue {
		fmt.Printf("worker %d start working:\n", ID)
		time.Sleep(1 * time.Second)
		result <- n * n
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	numJobs := 5
	numWorkers := 3
	result := make(chan int, numJobs)
	jobqueue := make(chan int, numJobs)

	// Fire up workers
	for wID := 0; wID < numWorkers; wID++ {
		go worker(&wg, wID, jobqueue, result)
	}

	// Send number to jobqueue
	for i := 0; i < numJobs; i++ {
		wg.Add(1)
		jobqueue <- i
	}

	wg.Wait()
	// After finish all job, close the channel
	close(jobqueue)

	for n := 0; n < numJobs; n++ {
		fmt.Println(<-result)
	}
}
