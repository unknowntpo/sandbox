package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // print "context deadline exceeded"
	case <-time.After(3 * time.Second):
		fmt.Println("overlaped")
	}
}
