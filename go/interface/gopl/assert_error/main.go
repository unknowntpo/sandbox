package main

import (
	"fmt"
	"time"
)

type customError struct {
	When time.Time
	What string
}

func (c *customError) Error() string {
	return fmt.Sprintf("custom error happend at %s, %s", c.When.String(), c.What)
}

func func1() error {
	return &customError{When: time.Now(), What: "something wrong in func1"}
}

func main() {
	err := func1()
	if err != nil {
		if e, ok := err.(*customError); ok {
			fmt.Println("Error as customError type")
			fmt.Println("error:", e)
		}
	}
}
