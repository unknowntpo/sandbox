package main

import (
	"errors"
	"fmt"
	//"os"
)

var Err_Something_Wrong = errors.New("In func2, something is wrong!")

func main() {
	err := func1()

	fmt.Println("In main")
	as_is(err)

	if err != nil {
		fmt.Println("Whole error: ", err)
		// FIXME: Why we panic if we don't use &Err_Something_Wrong as target of errors.As() ?
		if errors.As(err, &Err_Something_Wrong) {
			fmt.Printf("err as Err_Something_Wrong: %#v\n", err)
		} else {
			fmt.Printf("%#v\n", err)
		}

		// unwrap error
		err = errors.Unwrap(err)
		fmt.Println("Error after unwrap for 1 time: ", err)
	}

}

func func1() error {
	err := func2()
	fmt.Println("In func1")
	as_is(err)

	if err != nil {
		return fmt.Errorf("calling func2 failed: %w", err)
	}
	return nil
}

func func2() error {
	return Err_Something_Wrong
}

func as_is(err error) {
	if err == Err_Something_Wrong {
		fmt.Println("err == Err_Something_Wrong")
	}
	if errors.Is(err, Err_Something_Wrong) {
		fmt.Println("error is Err_Something_Wrong")
	}

	if e, ok := err.(*Err_Something_Wrong); ok {
		fmt.Println("error is Err_Something_Wrong type")
	}

	if errors.As(err, &Err_Something_Wrong) {
		fmt.Println("error as Err_Something_Wrong")
	}

}
