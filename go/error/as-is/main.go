package main

import (
	"errors"
	"fmt"
	//"os"
)

var Err_Something_Wrong = errors.New("In func2, something is wrong!")

func main() {
	err := func1()
	if err != nil {
		// FIXME: Why we panic ?
		if errors.As(err, Err_Something_Wrong) {
			fmt.Printf("err as Err_Something_Wrong: %#v\n", err)
		} else {
			fmt.Printf("%#v\n", err)
		}
	}
}

func func1() error {
	err := func2()
	if err != nil {
		return fmt.Errorf("calling func2 failed: %w", err)
	}
	return nil
}

func func2() error {
	return Err_Something_Wrong
}
