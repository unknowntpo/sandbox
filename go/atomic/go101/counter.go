package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var counter uint32

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddUint32(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter) // Output: 1000
}
