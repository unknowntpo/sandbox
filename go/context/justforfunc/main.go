package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	sleepAndTalkContext(ctx, 1*time.Second, "hello from sleepAndTalk")
}

// sleepAndTalkContext wraps the actual sleepAndTalk function, and do some context handling.
func sleepAndTalkContext(ctx context.Context, d time.Duration, msg string) {
	msgChan := make(chan string)
	go sleepAndTalk(msgChan, msg)

	select {
	case msg := <-msgChan:
		fmt.Println(msg)
		fmt.Println("sleepAndTalk done its job!")
	case <-ctx.Done():
		err := fmt.Errorf("sleepAndTalk does not finish its job: %v", ctx.Err())
		fmt.Println(err)
	}
}

// sleepAndTalk is the actual function to sleep for 5 seconds and send the message back to caller.
func sleepAndTalk(msgChan chan<- string, msg string) {
	time.Sleep(5 * time.Second)
	msgChan <- msg
}
