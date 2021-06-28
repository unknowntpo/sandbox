package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func gen(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int, 3)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-ctx.Done():
				fmt.Println("gen is canceled:", ctx.Err())
				return
			}
		}
	}()
	return out
}

func sq(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-ctx.Done():
				fmt.Println("sq is canceled:", ctx.Err())
				return
			}
		}
	}()
	return out
}

func merge(ctx context.Context, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed or it receives a value
	// from done, then output calls wg.Done.
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-ctx.Done():
				fmt.Println("output listener in merge is canceled: ", ctx.Err())
				return
			}
		}
	}

	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		fmt.Println("merge is canceled: ", ctx.Err())
		close(out)
	}()
	return out
}

func main() {
	// generate slice of int
	genNums := func() []int {
		nums := []int{0}
		for i := 0; i < 10; i++ {
			nums = append(nums, i)
		}

		return nums
	}

	nums := genNums()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	in := gen(ctx, nums...)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(ctx, in)
	c2 := sq(ctx, in)

	// Consume the first value from output.
	out := merge(ctx, c1, c2)
	for i := 0; i < len(nums); i++ {
		select {
		case num := <-out:
			fmt.Println(num)
		case <-ctx.Done():
			fmt.Println("output printer is canceled: ", ctx.Err())
			goto end
		}
		time.Sleep(200 * time.Millisecond)
	}
end:
	time.Sleep(2 * time.Second)
}
