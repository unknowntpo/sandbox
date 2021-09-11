package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := outer()
	fmt.Printf("%+v", err)
}

func inner() error {
	return errors.New("seems we have an error here")
}

func middle() error {
	err := inner()
	if err != nil {
		//return errors.Errorf("middle error: %v", err)
		//return errors.WithMessage(err, "middle error")
		return errors.Wrap(err, "middle")
	}
	return nil
}

func outer() error {
	err := middle()
	if err != nil {
		//return errors.Errorf("outer error: %v", err)
		//return errors.WithMessage(err, "outer error")
		return errors.Wrap(err, "outer")

	}
	return nil
}
