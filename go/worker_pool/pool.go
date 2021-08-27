package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type jobFunc func()

type Pool struct {
	// we use jobChan to communicate between caller of Pool and Pool.
	jobChan chan jobFunc
	// we use workerChan to send job from Pool to workers.
	workerChan chan jobFunc
	maxJobs    int
	maxWorkers int
	// use waitgroup to wait for workers done its job and retire.
	wg sync.WaitGroup
}

// Init inits goroutine pool with capacity of jobchan and workerchan
func New(maxJobs, maxWorkers int) *Pool {
	p := &Pool{
		jobChan:    make(chan jobFunc, maxJobs),
		workerChan: make(chan jobFunc, maxWorkers),
		maxJobs:    maxJobs,
		maxWorkers: maxWorkers,
	}

	return p
}
func (p *Pool) Start(ctx context.Context) {
	// Dispatch workers.
	p.wg.Add(p.maxWorkers)
	for i := 0; i < p.maxWorkers; i++ {
		go p.worker(ctx)
	}

	go func() {
		for {
			select {
			// Received a job.
			// Dispatch it to workers.
			case job := <-p.jobChan:
				fmt.Println("pool send a job to workerChan")
				p.workerChan <- job
			case <-ctx.Done():
				fmt.Println("Pool received cancel")
				return
			default:
			}
		}

	}()
	return
}

// Wait waits for all workers finish its job and retire.
func (p *Pool) Wait() {
	p.wg.Wait()
}

// Schedule sends the job the p.jobChan.
func (p *Pool) Schedule(job jobFunc) {
	fmt.Println("I send a job to pool")

	p.jobChan <- job
}

// worker is the worker that execute the job received from p.workerChan.
func (p *Pool) worker(ctx context.Context) {
	defer p.wg.Done()
	for {
		select {
		case job := <-p.workerChan:
			fmt.Println("worker received a job")
			job()
		case <-ctx.Done():
			fmt.Println("worker retired!")
			return
		}
	}
}

func main() {
	p := New(5, 5)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	p.Start(ctx)

	complete := make(chan struct{})
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT)

		<-sig
		// shutdown workerpool
		fmt.Println("Received SIGINT")
		cancel()
		p.Wait()
		complete <- struct{}{}
	}()

	p.Schedule(func() {
		defer fmt.Println("Sleeper exited")
		timeout := time.After(3 * time.Second)
		for {
			select {
			case <-timeout:
				fmt.Println("wake up!")
				return
			default:
				time.Sleep(1 * time.Second)
				fmt.Println("slept for 1 second.")
			}
		}
	})

	<-complete
	fmt.Println("Shutdown complete")
}
