package main

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrNoSuchFileOrDir = errors.New("no such file or directory")

func func1() error {
	err := openFile()
	if err != nil {
		// We only want to add context to error message, not to add additional call stack.
		return errors.WithMessage(err, "failed to open file")
	}
	return nil
}

func openFile() error {
	return ErrNoSuchFileOrDir
}

func main() {
	err := func1()
	if err != nil {
		switch {
		case errors.Is(err, ErrNoSuchFileOrDir):
			// Handle error gracefully.
			// log error
			fmt.Printf("%v\n", err)
		default:
			// Unexpected error
			// log the call stack
			fmt.Printf("%+v\n", err)
		}
	}
}
